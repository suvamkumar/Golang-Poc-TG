package users

import (
	"context"
	social_db "crud_with_TG/Golang-Poc-TG/internal/datasources/mongodb/socialdb"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//Insert user ingto the database
func (user *User) Insert() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("person")
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
	collection := social_db.GetMongoCollection("person")
	if err := collection.FindOne(ctx, User{ID: user.ID}).Decode(&user); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

//GetAllUser ...
func (user *User) GetAllUser() ([]User, *errors.RestErr) {
	users := make([]User, 0)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("person")
	cursor, err := collection.Find(ctx, User{})
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		fmt.Println("-----")
		fmt.Println(cursor)
		fmt.Println("-----")
		fmt.Println(user)
		fmt.Println("-----")
		users = append(users, user)
	}
	return users, nil
}

//UpdateUser ...
func (user *User) UpdateUser() *errors.RestErr {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("person")
	updateBson := bson.M{}
	if user.Name != "" {
		updateBson["name"] = user.Name
	}
	if user.Gender != "" {
		updateBson["gender"] = user.Gender
	}
	if user.Age != 0 {
		updateBson["age"] = user.Age
	}
	if user.State != "" {
		updateBson["state"] = user.State
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
	collection := social_db.GetMongoCollection("person")
	_, err := collection.DeleteOne(ctx, User{ID: user.ID})
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
