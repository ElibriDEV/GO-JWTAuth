package auth

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jwt-auth/initializators"
	"log"
	"strings"
	"time"
)

type UserToken struct {
	ID      primitive.ObjectID `bson:"_id"`
	Guid    string
	Refresh string
}

type tokenClaims struct {
	jwt.StandardClaims
	GUID string `json:"guid"`
}

type ServiceAuth struct {
	signKey string
}

func (service *ServiceAuth) generateSingleToken(guid string, expires time.Duration) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expires).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		guid,
	}).SignedString([]byte(service.signKey))
}

func (service *ServiceAuth) generateRefreshToken(guid string) string {
	refreshToken, err := service.generateSingleToken(guid, time.Duration(initializators.Config.RefreshTTL)*24*time.Hour)
	if err != nil {
		log.Fatal("Refresh signature error: ", err.Error())
	}
	_, err = initializators.MongoManager.Collection.UpdateOne(
		context.TODO(),
		bson.D{{"guid", guid}},
		bson.D{{"$set", bson.D{{"refresh", Encode(refreshToken)}}}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return refreshToken
}

func (service *ServiceAuth) generateAccessToken(guid string) string {
	accessToken, err := service.generateSingleToken(guid, time.Duration(initializators.Config.AccessTTL)*time.Minute)
	if err != nil {
		log.Fatal("Access signature error: ", err.Error())
	}
	return accessToken
}

func (service *ServiceAuth) generateTokens(guid string) (string, string) {
	refreshToken := service.generateRefreshToken(guid)
	accessToken := service.generateAccessToken(guid)
	return refreshToken, accessToken
}

func (service *ServiceAuth) parseToken(token string) (*tokenClaims, string, error) {
	if token == "" {
		return nil, "", errors.New("unauthorized")
	}
	tokenParts := strings.Split(token, " ")
	if len(tokenParts) != 2 {
		return nil, "", errors.New("unauthorized")
	}
	if tokenParts[0] != "Bearer" {
		return nil, "", errors.New("unauthorized")
	}
	parseToken, err := jwt.ParseWithClaims(tokenParts[1], &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unauthorized")
		}
		return []byte(service.signKey), nil
	})
	if err != nil {
		return nil, "", err
	}
	claims, ok := parseToken.Claims.(*tokenClaims)
	if !ok {
		return nil, "", errors.New("unauthorized")
	}
	return claims, tokenParts[1], nil
}

func (service *ServiceAuth) SignIn(guid string) (string, string) {
	return service.generateTokens(guid)
}

func (service *ServiceAuth) Refresh(accessToken, refreshToken string) (string, string, error) {
	accessData, _, err := service.parseToken(accessToken)
	if err != nil {
		return "", "", err
	}
	refreshData, refresh, err := service.parseToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	var result UserToken
	err = initializators.MongoManager.Collection.FindOne(
		context.TODO(),
		bson.D{{"guid", refreshData.GUID}},
	).Decode(&result)
	refreshEncoded := Encode(refresh)
	if err != nil || accessData.IssuedAt != refreshData.IssuedAt || refreshEncoded != result.Refresh {
		return "", "", errors.New("unauthorized")
	}
	newRefresh, newAccess := service.generateTokens(refreshData.GUID)
	return newRefresh, newAccess, nil
}

func NewAuthService() *ServiceAuth {
	return &ServiceAuth{initializators.Config.SignKey}
}
