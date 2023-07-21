package libs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DbHost string
var DbUser string
var DbPassword string
var DbName string
var DbPort string

var HashKey string
var TokenTtlHour int

var RedisHost string
var RedisPort string
var RedisPassword string
var RedisDB int

var SystemCompanyId uint
var RootCompanyId uint
var DemoCompanyId uint

var CacheKeyPrefix string

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error load env file")
	}

	sid, _ := strconv.Atoi(os.Getenv("SYSTEM_COMPANY_ID"))
	SystemCompanyId = uint(sid)

	rid, _ := strconv.Atoi(os.Getenv("ROOT_COMPANY_ID"))
	RootCompanyId = uint(rid)

	did, _ := strconv.Atoi(os.Getenv("DEMO_COMPANY_ID"))
	DemoCompanyId = uint(did)

	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbPort = os.Getenv("DB_PORT")

	HashKey = os.Getenv("HASH_KEY")

	TokenTtlHour, _ = strconv.Atoi(os.Getenv("TOKEN_TTL_SECOND"))

	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	if rdb, err := strconv.Atoi(os.Getenv("REDIS_DB")); err == nil {
		RedisDB = rdb
	}

	CacheKeyPrefix = os.Getenv("CACHE_KEY_PREFIX")
}

func IsSystem(id uint) bool {
	return id == SystemCompanyId
}

func IsRoot(id uint) bool {
	return id == RootCompanyId
}

func IsDemo(id uint) bool {
	return id == DemoCompanyId
}
