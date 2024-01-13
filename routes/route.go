package route

import (
	"github.com/gofiber/fiber/v2"
	handler_user_management "github.com/gozzafadillah/app/UsersManagement/handler"

	jwtware "github.com/gofiber/contrib/jwt"
)

type ControllerList struct {
	UserManagement handler_user_management.UserManagementHandler
}

func (cl *ControllerList) RouteRegister(f *fiber.App) {
	// without permission
	f.Post("/register", cl.UserManagement.RegisterUser)
	f.Post("/login", cl.UserManagement.Login)
	// JWT Middleware
	f.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
	// with permission
	f.Get("/users", cl.UserManagement.GetUsers)
}
