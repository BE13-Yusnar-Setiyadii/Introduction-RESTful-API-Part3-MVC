package main

import (
	"fmt"
	"yusnar/rest/api/mvc/config"
	"yusnar/rest/api/mvc/middlewares"
	"yusnar/rest/api/mvc/routes"
)

func main() {
	fmt.Println("run init")
	config.InitDB()
	config.InitialMigration()

	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
