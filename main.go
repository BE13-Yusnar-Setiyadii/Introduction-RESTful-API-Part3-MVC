package main

import (
	"fmt"
	"yusnar/rest/api/mvc/config"
	"yusnar/rest/api/mvc/routes"
)

func main() {
	fmt.Println("run init")
	config.InitDB()
	config.InitialMigration()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
