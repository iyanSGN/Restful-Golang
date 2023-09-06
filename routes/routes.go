package routes

import (
	"net/http"
	"new-project2/controllers"
	"new-project2/middleware"

	"github.com/labstack/echo"
) 


func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

	e.GET("/pegawai", controllers.FetchPegawai, middleware.IsAuthenticated)
	e.POST("/pegawai", controllers.StorePegawai)
	e.PUT("/pegawai/:id", controllers.UpdatePegawai)
	e.DELETE("pegawai/:id", controllers.DeletePegawai)

	e.GET("login/:password", controllers.GeneratedHashPassword)
	e.POST("/register", controllers.CheckLogin)

	return e
}