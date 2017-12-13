package category

import (
	"github.com/labstack/echo"

	. "work-codes/gohome/app/modules/category/controllers"
)

func InitRoute(api *echo.Group) {

	api.GET("/categorys", CategoryController.List)
	api.GET("/categorys/:id", CategoryController.Detail)
	api.POST("/categorys", CategoryController.Create)
	api.POST("/categorys/:id", CategoryController.Edit)
}
