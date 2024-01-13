package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	factory_user_management "github.com/gozzafadillah/app/UsersManagement"
	data_user_management "github.com/gozzafadillah/app/UsersManagement/data"
	route "github.com/gozzafadillah/routes"
)

var ctx = context.Background()

func main() {
	app := fiber.New()
	db, err := data_user_management.Connect()
	if err != nil {
		panic(err)
	}
	// userManagement
	userManagementHandler := factory_user_management.UserManagementFactory(db, ctx)

	RoutesInit := route.ControllerList{
		UserManagement: userManagementHandler,
	}

	RoutesInit.RouteRegister(app)

	log.Fatal(app.Listen(":8082"))
}
