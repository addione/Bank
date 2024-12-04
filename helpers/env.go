package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	LOCALE       = "locale"
	MONGO_URI    = "mongouri"
	MYSQL_URI    = "mysqluri"
	TOKEN_SECRET = "token_secret"
)

type EnvHelper struct {
}

func newEnvHelper() *EnvHelper {
	eh := &EnvHelper{}
	return eh
}

func GetEnvVariable(name string) (string, error) {
	godotenv.Load()
	variable := os.Getenv(name)
	return variable, nil
}
