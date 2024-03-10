package initializators

import (
	"github.com/joho/godotenv"
	"jwt-auth/common/convertor"
	"log"
	"os"
	"reflect"
)

var Config ConfigStruct

type ConfigStruct struct {
	MongoInitDBRootUsername string
	MongoInitDBRootPassword string
	MongoURL                string

	SignKey    string
	RefreshTTL int
	AccessTTL  int

	Hash []byte

	ApplicationPort string
}

func getEnv(path ...string) {
	if err := godotenv.Load(path...); err != nil {
		log.Fatal("Cannot load env: ", err.Error())
	}
}

func (environment *ConfigStruct) setEnvironmentData() {
	environment.MongoInitDBRootUsername = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	environment.MongoInitDBRootPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	environment.MongoURL = os.Getenv("MONGO_URL")

	environment.SignKey = os.Getenv("SIGN_KEY")
	environment.RefreshTTL = convertor.String2int(os.Getenv("REFRESH_TTL"))
	environment.AccessTTL = convertor.String2int(os.Getenv("ACCESS_TTL"))

	environment.Hash = []byte(os.Getenv("HASH_COST"))

	environment.ApplicationPort = os.Getenv("APPLICATION_PORT")
}

func (environment *ConfigStruct) validate() {
	environment.setEnvironmentData()

	field := reflect.ValueOf(*environment)
	typeOfS := field.Type()
	for fieldIndex := 0; fieldIndex < field.NumField(); fieldIndex++ {
		if field.Field(fieldIndex).Interface() == "" {
			log.Fatal(".env undefined key ", typeOfS.Field(fieldIndex).Name)
		}
	}
	return
}

func LoadEnv() {
	getEnv(".env")
	Config = ConfigStruct{}
	Config.setEnvironmentData()
	Config.validate()
	return
}
