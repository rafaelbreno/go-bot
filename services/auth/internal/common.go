package internal

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Common struct {
	Env    Env
	Logger *zap.Logger
}

type Env struct {
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgDBName   string
	PgName     string

	RedisHost     string
	RedisPort     string
	RedisName     string
	RedisDB       int
	RedisPassword string
}

func NewCommon() *Common {
	l, _ := zap.NewProduction()

	if os.Getenv("IS_PROD") != "true" {
		l.Info("Loading .env file")
		if err := godotenv.Load(); err != nil {
			l.Error(err.Error())
		}
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		l.Error(err.Error())
		return &Common{}
	}

	c := Common{
		Logger: l,
		Env: Env{
			PgHost:        os.Getenv("PGSQL_HOST"),
			PgPassword:    os.Getenv("PGSQL_PASSWORD"),
			PgUser:        os.Getenv("PGSQL_USER"),
			PgPort:        os.Getenv("PGSQL_PORT"),
			PgDBName:      os.Getenv("PGSQL_DBNAME"),
			PgName:        os.Getenv("PGSQL_NAME"),
			RedisHost:     os.Getenv("REDIS_HOST"),
			RedisPort:     os.Getenv("REDIS_PORT"),
			RedisName:     os.Getenv("REDIS_NAME"),
			RedisDB:       redisDB,
			RedisPassword: os.Getenv("REDIS_PASSWORD"),
		},
	}

	return &c
}
