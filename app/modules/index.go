package modules

import (
	"github.com/labstack/echo"

	"work-codes/gohome/app/modules/category"
)

func InitRoute(api *echo.Group) {
	category.InitRoute(api)
}
