package users

import (
	"context"
	users_db "crud_with_gin_gonic/internal/datasources/mongodb/usersdb"
	"crud_with_gin_gonic/internal/utils/errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

func init() {
	collection = users_db.GetMongoInstance().Database("usersdb").Collection("users")
}

//Insert user ingto the database
func (user *User) Insert() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// id := fmt.Sprintf("%v", res.InsertedID)
	// user.ID = id[10 : len(id)-2]
	return nil
}

//GetUser get single user from users db
func (user *User) GetUser(id string) *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := collection.FindOne(ctx, User{ID: user.ID}).Decode(&user); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//GetAllUser ...
func (user *User) GetAllUser() ([]User, *errors.RestErr) {
	users := make([]User, 0)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, User{})
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

//UpdateUser ...
func (user *User) UpdateUser() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	updateBson := bson.M{}
	if user.FirstName != "" {
		updateBson["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		updateBson["last_name"] = user.LastName
	}
	if user.Email != "" {
		updateBson["email"] = user.Email
	}
	if user.Status != "" {
		updateBson["status"] = user.Status
	}
	update := bson.M{"$set": updateBson}
	result, err := collection.UpdateOne(ctx, User{ID: user.ID}, update)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	fmt.Println(result.ModifiedCount)
	return nil
}

//DeleteUser ...
func (user *User) DeleteUser() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.DeleteOne(ctx, User{ID: user.ID})
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
