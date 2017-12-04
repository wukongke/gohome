package routes

import (
	"github.com/devfeel/dotweb"
)

func InitRoute(server *dotweb.HttpServer) {

	api := server.Group("/gohome/api")

	api.GET("/", func(ctx dotweb.Context) error {
		_, err := ctx.WriteString("Hello, GoHome! ")
		return err
	})
}
