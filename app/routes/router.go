package routes

import (
	"github.com/devfeel/dotweb"
)

func InitRoute(server *dotweb.HttpServer) {

	server.GET("/index", func(ctx dotweb.Context) error {
		_, err := ctx.WriteString("welcome to my first web!")
		return err
	})

	api := server.Group("/gohome/api")

	api.GET("/", func(ctx dotweb.Context) error {
		_, err := ctx.WriteString("Hello, GoHome! ")
		return err
	})
}
