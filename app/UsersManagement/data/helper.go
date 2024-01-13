package data_user_management

import user_management "github.com/gozzafadillah/app/UsersManagement/domain"

func ToDomain(data interface{}) user_management.User {
	return user_management.User{
		MongoID:   data.(user_management.User).MongoID,
		ID:        data.(user_management.User).ID,
		FirstName: data.(user_management.User).FirstName,
		LastName:  data.(user_management.User).LastName,
		Email:     data.(user_management.User).Email,
		Password:  data.(user_management.User).Password,
		Fee:       data.(user_management.User).Fee,
	}
}
