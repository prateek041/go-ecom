package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUri      string
	Port       string
	Publichost string
}

var ENV = initConfig()

func initConfig() *Config {
	godotenv.Load()

	return &Config{
		DBUri:      getEnv("MONGO_URL", "mongodb://localhost:27017"),
		Port:       getEnv("PORT", ":9090"),
		Publichost: getEnv("PUBLICHOST", "http://localhost:"),
	}
}

func getEnv(key, fallback string) string {
	env, ok := os.LookupEnv("MONGO_URL")
	if !ok {
		log.Printf("Env Var: %s not found using Fallback: %s", key, fallback)
		return fallback
	}
	return env
}
