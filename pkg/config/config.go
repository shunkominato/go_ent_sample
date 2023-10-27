package config

import "os"

func SetConfig() {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		setEnvProd()
	} else if env == "staging" {
		setEnvStage()
	} else {
		setEnvLocal()
	}
}

func setEnvProd() {
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "db")
	os.Setenv("DB_PORT", "2")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_DB_NAME", "engin")
}

func setEnvStage() {
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "db")
	os.Setenv("DB_PORT", "1111")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_DB_NAME", "engin")
}

func setEnvLocal() {
	os.Setenv("PORT", "8080")
	os.Setenv("DB_HOST", "db")
	os.Setenv("DB_PORT", "1111")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_DB_NAME", "engin")
}