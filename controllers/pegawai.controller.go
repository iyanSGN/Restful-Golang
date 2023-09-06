package controllers

import (
	"net/http"
	"new-project2/models"
	"strconv"

	"github.com/labstack/echo"
)

func FetchPegawai(c echo.Context) error {
	result, err := models.FetchPegawai()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string {"Message" : err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama 		:= c.FormValue("nama")
	pekerjaan 	:= c.FormValue("pekerjaan")
	umur 		:= c.FormValue("umur")
	telepon 	:= c.FormValue("telepon")
	email 		:= c.FormValue("email")
	password 	:= c.FormValue("password")

	result, err := models.StorePegawai(nama, pekerjaan, umur, telepon, email, password)
	if err != nil  {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)

	
}

func UpdatePegawai(c echo.Context) error {
	myId := c.Param("id")
	nama 		:= c.FormValue("nama")
	pekerjaan 	:= c.FormValue("pekerjaan")
	umur 		:= c.FormValue("umur")
	telepon 	:= c.FormValue("telepon")
	email 		:= c.FormValue("email")
	password 	:= c.FormValue("password")

	id, err := strconv.Atoi(myId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
		

	result, err := models.UpdatePegawai(id, nama, pekerjaan, umur, telepon, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)	
}

func DeletePegawai(c echo.Context) error {
	myId := c.Param("id")

	id, err := strconv.Atoi(myId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}