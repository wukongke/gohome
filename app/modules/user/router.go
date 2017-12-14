package user

import (
	. "work-codes/gohome/app/modules/user/controllers"

	"github.com/labstack/echo"
)

func InitRoute(api *echo.Group) {

	api.GET("/users", UserController.List)
	api.POST("/users", UserController.Create)
	api.POST("/demo", UserController.Demo)
}
