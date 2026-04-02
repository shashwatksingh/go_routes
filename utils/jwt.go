package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "dummy"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "userId": userId, "exp": time.Now().Add(time.Hour * 2).Unix()})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Signing Method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("Could not parse token")
	}

	tokenValid := parsedToken.Valid
	if !tokenValid {
		return errors.New("Invalid Token!")
	}
	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Invalid Claim!")
	// }
	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
