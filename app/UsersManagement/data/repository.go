package data_user_management

import (
	"context"
	"crypto/rand"
	"math/big"

	user_management "github.com/gozzafadillah/app/UsersManagement/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserManagementRepository struct {
	Mongo *mongo.Database
	Ctx   context.Context
}

// GetUserByEmail implements user_management.UserRepository.
func (um *UserManagementRepository) GetUserByEmail(email string) (user_management.User, error) {
	var user user_management.User
	err := um.Mongo.Collection("User").FindOne(um.Ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user_management.User{}, err
	}
	return user, nil
}

// DeleteUser implements user_management.UserRepository.
func (um *UserManagementRepository) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUser implements user_management.UserRepository.
func (um *UserManagementRepository) GetUser(id int) (user_management.User, error) {
	var user user_management.User
	err := um.Mongo.Collection("User").FindOne(um.Ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return user_management.User{}, err
	}
	return user, nil
}

// GetUsers implements user_management.UserRepository.
func (um *UserManagementRepository) GetUsers() ([]user_management.User, error) {
	projection := bson.M{
		"password": 0, // Menentukan field "password" untuk tidak disertakan
		"id":       0, // Menentukan field "id" untuk tidak disertakan
	}
	// Mengambil kursor data dari koleksi "User" dalam MongoDB
	res, err := um.Mongo.Collection("User").Find(um.Ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer res.Close(um.Ctx)

	// Membuat slice untuk menyimpan data pengguna
	var users []user_management.User

	// Iterasi melalui hasil kursor
	for res.Next(um.Ctx) {
		var v user_management.User // Gantilah YourUserType dengan tipe data yang sesuai
		if err := res.Decode(&v); err != nil {
			return nil, err
		}
		users = append(users, ToDomain(v))
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Login implements user_management.UserRepository.
func (um *UserManagementRepository) Login(email string, password string) (string, error) {
	var user user_management.User
	err := um.Mongo.Collection("User").FindOne(um.Ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

// Register implements user_management.UserRepository.
func (um *UserManagementRepository) Register(user user_management.User) (user_management.User, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	user.ID = int(randNumber.Int64())
	_, err := um.Mongo.Collection("User").InsertOne(um.Ctx, user)

	if err != nil {
		return user_management.User{}, err
	}

	return user, nil
}

// UpdateUser implements user_management.UserRepository.
func (um *UserManagementRepository) UpdateUser(user user_management.User, id int) (user_management.User, error) {
	panic("unimplemented")
}

func NewUserManagementRepository(mongo *mongo.Database, ctx context.Context) user_management.UserRepository {
	return &UserManagementRepository{
		Mongo: mongo,
		Ctx:   ctx,
	}
}
