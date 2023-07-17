package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/toufiq-austcse/go-api-boilerplate/config"
	"time"
)

type JwtCustomClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id int) (token string, expireAt int64, err error) {
	expireAt = time.Now().AddDate(0, 0, 1).Unix()
	claims := JwtCustomClaims{Id: id, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireAt,
		Issuer:    config.AppConfig.JWT_ISSUER,
		IssuedAt:  time.Now().Unix(),
	}}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := jwtToken.SignedString([]byte(config.AppConfig.JWT_SECRET_KEY))
	if err != nil {
		return "", 0, err
	}
	return signedString, expireAt, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(config.AppConfig.JWT_SECRET_KEY), nil
	})
}

func GetCompanyIdFromToken(token *jwt.Token) (int, error) {
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println("claims ", claims)
	return int(claims["id"].(float64)), nil

}
