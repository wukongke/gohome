package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/logger"

	"work-codes/gohome/app/routes"
)

func StartServer() error {
	app := dotweb.New()
	logger.SetEnabledConsole(true)
	routes.InitRoute(app.HttpServer)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	fmt.Println("port: ", port)
	err = app.StartServer(port)
	return err
}

func main() {
	err := StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
