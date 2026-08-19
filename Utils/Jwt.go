package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "gune1997"

func JwtToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretkey))
}

func VerifyToken(token string) (int64, error) {
	vToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse the token")
	}
	verifiedToken := vToken.Valid
	if !verifiedToken {
		return 0, errors.New("Invalid Token")
	}

	claims, ok := vToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}
	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil

}
