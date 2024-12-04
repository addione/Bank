package helpers

import (
	"errors"
	"math/big"
	"time"

	"github.com/addione/New/models"
	"github.com/golang-jwt/jwt/v5"
)

type JwtHelper struct {
	secret string
}

func newJwtHelper() *JwtHelper {
	secret, _ := GetEnvVariable(TOKEN_SECRET)

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
		"exp":    time.Now().Add(time.Hour * 1000).Unix(),
	})

	return token.SignedString([]byte(j.secret))
}

func (j *JwtHelper) VerifyToken(token string) (*models.UserMysql, error) {

	user := &models.UserMysql{}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token.")
		}

		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return user, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return user, errors.New("invalid token claims")
	}

	user.Email = claims["email"].(string)
	user.ID = int64(claims["userId"].(float64))

	return user, nil
}
