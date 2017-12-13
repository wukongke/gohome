package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"work-codes/gohome/app/config"
	"work-codes/gohome/app/modules"
)

func init() {
	fmt.Println("ApiPre: ", config.APIConfig.Prefix)
}

func InitRoute(app *echo.Echo) {

	// 静态文件
	app.Static("/gohome/public", "public")
	app.File("/*", "public/index.html")

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	api := app.Group(config.APIConfig.Prefix)

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, GoHome!")
	})

	modules.InitRoute(api)
}
