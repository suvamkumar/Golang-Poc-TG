package users

import (
	"context"

	"crud_with_TG/Golang-Poc-TG/external/services/tigergraph"
	social_db "crud_with_TG/Golang-Poc-TG/internal/datasources/mongodb/socialdb"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"

	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//Insert user into the database
func (user *User) Insert() (map[string]interface{}, *errors.RestErr) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("person")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	tg := tigergraph.TG{ConnectionString: "http://localhost:9000/graph"}
	response := tg.UpsertSingleVertex("social", "person", user.Name, user)
	// id := fmt.Sprintf("%v", res.InsertedID)
	// user.ID = id[10 : len(id)-2]
	return response, nil
}

//InsertMany user into the database
func (user *User) InsertMany(bData []interface{}) (map[string]interface{}, *errors.RestErr) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("person")
	_, err := collection.InsertMany(ctx, bData)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	tg := tigergraph.TG{ConnectionString: "http://localhost:9000/graph"}
	response := tg.UpsertMultipleVertex("social", "person", bData)
	return response, nil
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
