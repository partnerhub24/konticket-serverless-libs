package environment

import (
	"fmt"
	"os"
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

type Evironment struct {
	App      AppEnvironment
	Database DatabaseEnvironment
}

var env Evironment

func GetEnv() Evironment {
	if env.App.AppEnv != "" {
		fmt.Println("Use ENV from Global cache.")
		return env
	}

	fmt.Println("ENV Global cache is missing, Get env in process...")

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

	env = Evironment{
		App:      app,
		Database: database,
	}

	return env
}
