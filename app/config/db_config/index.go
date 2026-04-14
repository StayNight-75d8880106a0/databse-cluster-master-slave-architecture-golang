package db_config

import "os"

type DBConfig struct {
	MASTER_HOST string
	SLAVE1_HOST string
	SLAVE2_HOST string
	SLAVE3_HOST string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
	DB_TIMEZONE string
}

func DB_Config() *DBConfig {
	return &DBConfig{
		MASTER_HOST: os.Getenv("MASTER_HOST"),
		SLAVE1_HOST: os.Getenv("SLAVE1_HOST"),
		SLAVE2_HOST: os.Getenv("SLAVE2_HOST"),
		SLAVE3_HOST: os.Getenv("SLAVE3_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		DB_TIMEZONE: os.Getenv("DB_TIMEZONE"),
	}
}
