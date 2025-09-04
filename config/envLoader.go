package config

import (
	"os"
)

var (
	PORT       string
	SECRET_KEY string
	USER_EMAIL string
	USER_PASS  string
	DSN        string
)

// init() automatically loads .env and sets package-level variables
func init() {

	//err := godotenv.Load() // load env into memory or process's environment, useful for local dev if env file is used, no need if env is used from environment variable provided by os
	// if err != nil {
	// 	log.Println("❌ Error loading .env file. ⚠️  No .env file found. Using system environment variables.")
	// }

	PORT = getEnv("GO_PORT", "8080")

	SECRET_KEY = os.Getenv("JWT_SECRET")
	USER_EMAIL = os.Getenv("EMAIL_USER")
	USER_PASS = os.Getenv("EMAIL_PASS")
	DSN = os.Getenv("DSN")

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
