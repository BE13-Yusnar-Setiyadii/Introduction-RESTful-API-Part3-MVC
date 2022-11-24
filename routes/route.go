package routes

import (
	"fmt"
	"yusnar/rest/api/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	fmt.Println("run main")
	e := echo.New()
	e.GET("/users", controllers.GetUsersController)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserByIdController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)
	e.Logger.Fatal(e.Start(":8000"))

	return e
}
