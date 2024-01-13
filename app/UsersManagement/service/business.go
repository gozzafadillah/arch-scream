package service_user_management

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	user_management "github.com/gozzafadillah/app/UsersManagement/domain"
	encryption "github.com/gozzafadillah/app/UsersManagement/helper"
)

type UserManagementService struct {
	UserManagementRepository user_management.UserRepository
}

// DeleteUser implements user_management.UserService.
func (*UserManagementService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUser implements user_management.UserService.
func (*UserManagementService) GetUser(id int) (user_management.User, error) {
	panic("unimplemented")
}

// GetUsers implements user_management.UserService.
func (um *UserManagementService) GetUsers() ([]user_management.User, error) {
	res, err := um.UserManagementRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Login implements user_management.UserService.
func (um *UserManagementService) Login(email string, password string) (string, error) {
	getUser, err := um.UserManagementRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	cekPassword := encryption.CheckPasswordHash(password, getUser.Password)
	if !cekPassword {
		return "", err
	}
	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// Register implements user_management.UserService.
func (um *UserManagementService) Register(user user_management.User) (user_management.User, error) {
	newPassword, err := encryption.HashPassword(user.Password)
	if err != nil {
		return user_management.User{}, err
	}
	user.Password = newPassword
	res, err := um.UserManagementRepository.Register(user)
	if err != nil {
		return user_management.User{}, err
	}
	return res, nil
}

// UpdateUser implements user_management.UserService.
func (*UserManagementService) UpdateUser(user user_management.User, id int) (user_management.User, error) {
	panic("unimplemented")
}

func NewUserManagementService(userManagementRepository user_management.UserRepository) user_management.UserService {
	return &UserManagementService{
		UserManagementRepository: userManagementRepository,
	}
}
