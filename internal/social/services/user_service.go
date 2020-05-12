package services

import (
	users "crud_with_TG/Golang-Poc-TG/internal/social/domain"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	//UserService available to others to call the method
	UserService userServiceInterface = &userService{}
)

type userService struct{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User, string) (*users.User, *errors.RestErr)
	GetUser(id string) (*users.User, *errors.RestErr)
	DeleteUser(id string) *errors.RestErr
	GetAllUser() ([]users.User, *errors.RestErr)
}

//CreateUser ...
func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.ID = primitive.NewObjectID()
	if err := user.Insert(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser ...
func (s *userService) GetUser(id string) (*users.User, *errors.RestErr) {
	user := users.User{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	user.ID = objectID
	if err := user.GetUser(id); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetAllUser ...
func (s *userService) GetAllUser() ([]users.User, *errors.RestErr) {
	user := users.User{}
	return user.GetAllUser()
}

//UpdateUser ...
func (s *userService) UpdateUser(user users.User, id string) (*users.User, *errors.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	user.ID = objectID
	if err := user.UpdateUser(); err != nil {
		return nil, err
	}
	return &user, nil
}

//DeleteUser ...
func (s *userService) DeleteUser(id string) *errors.RestErr {
	var user users.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	user.ID = objectID
	if err := user.DeleteUser(); err != nil {
		return err
	}
	return nil
}
