package handler_user_management

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	user_management "github.com/gozzafadillah/app/UsersManagement/domain"
)

type UserManagementHandler struct {
	UserManagementService user_management.UserService
}

func NewUserManagementHandler(userManagementService user_management.UserService) UserManagementHandler {
	return UserManagementHandler{
		UserManagementService: userManagementService,
	}
}

func (umh *UserManagementHandler) RegisterUser(c *fiber.Ctx) error {
	req := user_management.User{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	res, err := umh.UserManagementService.Register(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data":    res,
	})
}

func (umh *UserManagementHandler) GetUsers(c *fiber.Ctx) error {
	deadline, _ := c.Context().Deadline()
	err := c.Context().Conn().SetReadDeadline(deadline.Add(15 * time.Second))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	res, err := umh.UserManagementService.GetUsers()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    res,
	})
}

func (umh *UserManagementHandler) Login(c *fiber.Ctx) error {
	req := user_management.Login{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	res, err := umh.UserManagementService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":     "Success",
		"accessToken": res,
	})
}
