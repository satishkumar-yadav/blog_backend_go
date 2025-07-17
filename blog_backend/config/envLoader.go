package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
const (
	SECRET_KEY string
)*/

var (
	PORT          string
	ROUTER_TYPE   string
	DB_TYPE       string
	MYSQL_DB      string
	SECRET_KEY    string
	USER_EMAIL    string
	USER_PASS     string
	DSN           string
	LOGIN_QUERY   string
	REG_CHECK     string
	APPROVE_QUERY string
)

// init() automatically loads .env and sets package-level variables
func init() {
	err := godotenv.Load() // load env into memory
	if err != nil {
		log.Println("❌ Error loading .env file. ⚠️  No .env file found. Using system environment variables.")
	}

	PORT = getEnv("GO_PORT", "8080")
	ROUTER_TYPE = getEnv("ROUTER_TYPE", "fiber")
	DB_TYPE = getEnv("DB_TYPE", "gorm")
	MYSQL_DB = getEnv("MYSQL_DATABASE", "mysql db")

	SECRET_KEY = os.Getenv("JWT_SECRET")
	USER_EMAIL = os.Getenv("EMAIL_USER")
	USER_PASS = os.Getenv("EMAIL_PASS")
	LOGIN_QUERY = os.Getenv("LOGIN_QUERY")
	REG_CHECK = os.Getenv("REG_CHECK")
	APPROVE_QUERY = os.Getenv("APPROVE_QUERY")

	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")

	DB_HOST := getEnv("DB_HOST", "localhost")
	MYSQL_PORT := getEnv("MYSQL_PORT", "3306")
	// DB_HOST_URL := getEnv("DB_HOST_URL", "127.0.0.1")
	// DB_REMOTE_HOST := getEnv("DB_REMOTE_HOST", "db remote")
	// MYSQL_REMOTE_HOST := getEnv("MYSQL_REMOTE_HOST", "mysql remote")
	MYSQL_DSN := getEnv("DSN", "user:password@tcp(127.0.0.1:3306)/db")

	DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASS, DB_HOST, MYSQL_PORT, MYSQL_DB)
	DSN = MYSQL_DSN

}

// getEnv returns the value of an environment variable or a fallback
func getEnv(key, fallback string) string { // key-the name of environment variable we want to get, fallback-default value to return if environment variable is not set or empty
	if value := os.Getenv(key); value != "" {
		return value
	}
	// fmt.Printf("Missing required environment variable : %s, so using fallback values\n", key)
	// log.Fatalf("Missing required environment variable : %s, so using fallback values", key)
	return fallback
}
