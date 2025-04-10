package environment

import (
	"os"

	"github.com/partnerhub24/konticket-serverless-libs/logger"
)

type AppEnvironment struct {
	AppEnv string
}

type DatabaseEnvironment struct {
	Driver string
	User   string
	Pass   string
	Host   string
	Port   string
	Name   string
}

type MongoDBEnvironment struct {
	MongoDriver string
	MongoUser   string
	MongoPass   string
	MongoHost   string
	MongoPort   string
	MongoName   string
}

type Evironment struct {
	App      AppEnvironment
	Database DatabaseEnvironment
	MongoDB  MongoDBEnvironment
}

var env Evironment

func GetEnv() Evironment {
	if env.App.AppEnv != "" {
		logger.Info("Use ENV from Global cache.")
		return env
	}

	logger.Info("ENV Global cache is missing, Get env in process...")

	app := AppEnvironment{
		AppEnv: os.Getenv("APP_ENV"),
	}

	database := DatabaseEnvironment{
		Driver: os.Getenv("DB_DRIVER"),
		User:   os.Getenv("DB_USER"),
		Pass:   os.Getenv("DB_PASS"),
		Host:   os.Getenv("DB_HOST"),
		Port:   os.Getenv("DB_PORT"),
		Name:   os.Getenv("DB_NAME"),
	}

	mongoDB := MongoDBEnvironment{
		MongoDriver: os.Getenv("MONGO_DRIVER"),
		MongoUser:   os.Getenv("MONGO_USER"),
		MongoPass:   os.Getenv("MONGO_PASS"),
		MongoHost:   os.Getenv("MONGO_HOST"),
		MongoPort:   os.Getenv("MONGO_PORT"),
		MongoName:   os.Getenv("MONGO_NAME"),
	}

	env = Evironment{
		App:      app,
		Database: database,
		MongoDB:  mongoDB,
	}
	return env
}
