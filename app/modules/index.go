package modules

import (
	"github.com/labstack/echo"

	"work-codes/gohome/app/modules/category"
	"work-codes/gohome/app/modules/user"
)

func InitRoute(api *echo.Group) {
	category.InitRoute(api)
	user.InitRoute(api)
}
