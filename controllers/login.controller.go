package controllers

import (
	"fmt"
	"net/http"
	"new-project2/helpers"
	"new-project2/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)


func CheckLogin (c echo.Context)  error {
	username := c.FormValue("username")
	password :=	c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Iyan Siagian"
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("mySecretKey"))
	if err != nil {
		fmt.Printf("ERROR : %v", err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func GeneratedHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashingPassword(password)

	return c.JSON(http.StatusOK, hash)
}

