package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigApp struct {
	Server    Server
	KamipaDB  KamipaDatabase
	SimipaDB  SimipaDatabase
	JWT       JWT
	Redis     Redis
	Mediamipa Mediamipa
	Midtrans  Midtrans
}

func GetConfig() *ConfigApp {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &ConfigApp{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		KamipaDB: KamipaDatabase{
			Host:     os.Getenv("KAMIPA_DB_HOST"),
			Port:     os.Getenv("KAMIPA_DB_PORT"),
			Username: os.Getenv("KAMIPA_DB_USERNAME"),
			Password: os.Getenv("KAMIPA_DB_PASSWORD"),
			DBName:   os.Getenv("KAMIPA_DB_NAME"),
		},

		SimipaDB: SimipaDatabase{
			Host:     os.Getenv("SIMIPA_DB_HOST"),
			Port:     os.Getenv("SIMIPA_DB_PORT"),
			Username: os.Getenv("SIMIPA_DB_USERNAME"),
			Password: os.Getenv("SIMIPA_DB_PASSWORD"),
			DBName:   os.Getenv("SIMIPA_DB_NAME"),
		},

		JWT: JWT{
			Key:    os.Getenv("JWT_KEY"),
			Issuer: os.Getenv("JWT_ISSUER"),
		},

		Redis: Redis{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		Mediamipa: Mediamipa{
			BaseUrl: os.Getenv("MEDIAMIPA_BASE_URL"),
		},
		Midtrans: Midtrans{
			Key:    os.Getenv("MIDTRANS_KEY"),
			IsProd: os.Getenv("MIDTRANS_ENV") == "Production",
		},
	}
}
