package services

import (
	users "crud_with_TG/Golang-Poc-TG/internal/social/domain"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	//FriendshipService available to others to call the method
	FriendshipService friendshipServiceInterface = &friendshipService{}
)

type friendshipService struct{}

type friendshipServiceInterface interface {
	CreateUser(users.Friendship) (*users.Friendship, *errors.RestErr)
	UpdateUser(users.Friendship, string) (*users.Friendship, *errors.RestErr)
	GetUser(id string) (*users.Friendship, *errors.RestErr)
	DeleteUser(id string) *errors.RestErr
	GetAllUser() ([]users.Friendship, *errors.RestErr)
}

//CreateUser ...
func (s *friendshipService) CreateUser(user users.Friendship) (*users.Friendship, *errors.RestErr) {
	user.ID = primitive.NewObjectID()
	if err := user.Insert(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser ...
func (s *friendshipService) GetUser(id string) (*users.Friendship, *errors.RestErr) {
	friends := users.Friendship{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	friends.ID = objectID
	if err := friends.GetUser(id); err != nil {
		return nil, err
	}
	return &friends, nil
}

//GetAllUser ...
func (s *friendshipService) GetAllUser() ([]users.Friendship, *errors.RestErr) {
	friends := users.Friendship{}
	return friends.GetAllUser()
}

//UpdateUser ...
func (s *friendshipService) UpdateUser(user users.Friendship, id string) (*users.Friendship, *errors.RestErr) {
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
func (s *friendshipService) DeleteUser(id string) *errors.RestErr {
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
