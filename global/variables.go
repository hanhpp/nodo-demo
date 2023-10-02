package global

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
)

var Config *VecConfig

const version = "v1"
const serviceName = "demo-gin"

type VecConfig struct {

	// Server
	ServerPort string
	Prefix     string

	// DB
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
	DbSSLMode  string
	DbTimeZone string

	// Keys
	EncodeIdKey string

	// JWTAuthorize
	SecretKey          string
	DefaultJWTDuration time.Duration
}

func (cf *VecConfig) initDB() {

	/* Postgres DB. Change all DB identifiers below. */
	cf.DbHost = getEnv("DB_HOST", "127.0.0.1")
	cf.DbPort = getEnv("DB_PORT", "5432")
	cf.DbUsername = getEnv("DB_USERNAME", "")
	cf.DbPassword = getEnv("DB_PASSWORD", "")
	cf.DbName = getEnv("DB_NAME", "")
	cf.DbSSLMode = getEnv("DB_SSL_MODE", "disable")
	cf.DbTimeZone = getEnv("DB_TIME_ZONE", "GMT")

}

func (cf *VecConfig) initWeb() {
	// prefix = service + version for api usage
	cf.Prefix = fmt.Sprintf("%s/%s", serviceName, version)
	cf.ServerPort = getEnv("SERVER_PORT", "8080")
}

func (cf *VecConfig) initAuth() {
	cf.SecretKey = getEnv("JWT_KEY", "")

	/* This key was used to generate the session token information. */
	cf.EncodeIdKey = "encode-id-key"

	/* Time duration variables */
	cf.DefaultJWTDuration = time.Hour * 1
}

func newVecConfig() *VecConfig {
	cf := VecConfig{}
	/*Load .env */
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cf.initDB()
	cf.initWeb()
	cf.initAuth()
	// fmt.Printf("%+v\n", cf)

	return &cf
}

func FetchEnvs() {
	Config = newVecConfig()
}
