package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"healthy-pacil/config"
	"healthy-pacil/model"
)

func GetUserClaims(c *fiber.Ctx) (*model.CustomClaim, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &model.CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.CustomClaim)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is expired")
	}

}
