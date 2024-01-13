package factory_user_management

import (
	"context"

	data_user_management "github.com/gozzafadillah/app/UsersManagement/data"
	handler_user_management "github.com/gozzafadillah/app/UsersManagement/handler"
	service_user_management "github.com/gozzafadillah/app/UsersManagement/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserManagementFactory(db *mongo.Database, ctx context.Context) handler_user_management.UserManagementHandler {
	// userManagement
	userManagementRepository := data_user_management.NewUserManagementRepository(db, ctx)
	userManagementService := service_user_management.NewUserManagementService(userManagementRepository)
	userManagementHandler := handler_user_management.NewUserManagementHandler(userManagementService)
	return userManagementHandler
}
