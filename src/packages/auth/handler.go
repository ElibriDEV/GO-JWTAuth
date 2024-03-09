package auth

import (
	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
	"jwt-auth/src/response"
	"net/http"
)

type HandlerAuth struct{}

func CreateGroup(rootGroup *gin.RouterGroup) *gin.RouterGroup {
	authHandler := &HandlerAuth{}
	group := rootGroup.Group("/auth")
	{
		group.GET("/sign-in", authHandler.signIn)
		group.GET("/refresh", authHandler.refresh)
	}
	return group
}

// @Summary      Sign-in
// @Description  Get access & refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        guid   query      string  true  "User-GUID"
// @Success      200  {object}  TokenResponse
// @failure 400 {object} response.ErrorResponse
// @Router       /auth/sign-in/ [get]
func (handler HandlerAuth) signIn(context *gin.Context) {
	guidParam := context.Query("guid")
	if guidParam == "" {
		response.ThrowError(context, http.StatusBadRequest, "GUID is required.")
		return
	}
	if guid.IsGuid(guidParam) == false {
		response.ThrowError(context, http.StatusBadRequest, "Invalid GUID.")
		return
	}
	refresh, access := NewAuthService().SignIn(guidParam)
	context.JSON(http.StatusOK, map[string]interface{}{
		"refreshToken": refresh,
		"accessToken":  access,
	})
}

// @Summary      Refresh
// @Description  Update access & refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer-Access
// @Security     Bearer-Refresh
// @Success      200  {object}  TokenResponse
// @failure 400 {object} response.ErrorResponse
// @failure 401 {object} response.ErrorResponse
// @Router       /auth/refresh/ [get]
func (handler HandlerAuth) refresh(context *gin.Context) {
	access := context.GetHeader("Access")
	refresh := context.GetHeader("Refresh")
	newRefresh, newAccess, err := NewAuthService().Refresh(access, refresh)
	if err != nil {
		response.ThrowError(context, http.StatusUnauthorized, err.Error())
		return
	}
	context.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  newAccess,
		"refreshToken": newRefresh,
	})
}
