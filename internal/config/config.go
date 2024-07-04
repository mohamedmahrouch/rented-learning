package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DatabaseURL string
}

func LoadConfig() Config {
    err := godotenv.Load()
    if err != nil {

        panic("Error loading .env file")
    }

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {

        panic("DATABASE_URL environment variable is not set")
    }

    return Config{
        DatabaseURL: dbURL,
    }
}
