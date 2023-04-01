package initializers

import "github.com/joho/godotenv"

func LoadEnv() {
	// Ignore error not loaded .env because environment variables will be injected when we build image
	godotenv.Load()
}
