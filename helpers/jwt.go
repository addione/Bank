package helpers

import (
	"math/big"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JwtHelper struct {
	secret string
}

func newJwtHelper() *JwtHelper {
	godotenv.Load()
	secret := os.Getenv("tokenSecret")

	return &JwtHelper{
		secret: secret,
	}
}

type PrivateKey struct {
	secret string
	D      *big.Int
}

func (j *JwtHelper) GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 10).Unix(),
	})

	return token.SignedString([]byte(j.secret))
}
