package users

import (
	"context"
	"crud_with_TG/Golang-Poc-TG/external/services/tigergraph"
	social_db "crud_with_TG/Golang-Poc-TG/internal/datasources/mongodb/socialdb"
	"crud_with_TG/Golang-Poc-TG/internal/utils/errors"
	"time"
)

//Insert user ingto the database
// func (friendship *Friendship) Insert() *errors.RestErr {
// 	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	// _, err := collection.InsertOne(ctx, user)
// 	// if err != nil {
// 	// 	return errors.NewInternalServerError(err.Error())
// 	// }
// 	// id := fmt.Sprintf("%v", res.InsertedID)
// 	// user.ID = id[10 : len(id)-2]
// 	return nil
// }

//Insert ...
func (friendship *Friendship) Insert() (map[string]interface{}, *errors.RestErr) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("friendship")
	_, err := collection.InsertOne(ctx, friendship)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	tg := tigergraph.TG{ConnectionString: "http://localhost:9000/graph"}
	response := tg.UpsertSingleEdge("social", "friendship", friendship.From, friendship.To)
	return response, nil
}

//InsertMany user into the database
func (friendship *Friendship) InsertMany(bData []interface{}) (map[string]interface{}, *errors.RestErr) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("friendship")
	_, err := collection.InsertMany(ctx, bData)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	tg := tigergraph.TG{ConnectionString: "http://localhost:9000/graph"}
	response := tg.UpsertManyEdges("social", "friendship", bData)
	return response, nil
}

//GetUser get single user from users db
func (friendship *Friendship) GetUser(id string) *errors.RestErr {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// if err := collection.FindOne(ctx, User{ID: user.ID}).Decode(&user); err != nil {
	// 	return errors.NewInternalServerError(err.Error())
	// }
	return nil
}

//GetAllFriends ...
func (friendship *Friendship) GetAllFriends() ([]Friendship, *errors.RestErr) {
	friends := make([]Friendship, 0)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := social_db.GetMongoCollection("friendship")
	cursor, err := collection.Find(ctx, Friendship{})
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var friend Friendship
		cursor.Decode(&friend)
		friends = append(friends, friend)
	}
	return friends, nil
}

//UpdateUser ...
func (friendship *Friendship) UpdateUser() *errors.RestErr {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// updateBson := bson.M{}
	// if user.Name != "" {
	// 	updateBson["name"] = user.Name
	// }
	// if user.Gender != "" {
	// 	updateBson["gender"] = user.Gender
	// }
	// if user.Age != "" {
	// 	updateBson["age"] = user.Age
	// }
	// if user.State != "" {
	// 	updateBson["state"] = user.State
	// }
	// update := bson.M{"$set": updateBson}
	// result, err := collection.UpdateOne(ctx, User{ID: user.ID}, update)
	// if err != nil {
	// 	return errors.NewInternalServerError(err.Error())
	// }
	// fmt.Println(result.ModifiedCount)
	return nil
}

//DeleteUser ...
func (friendship *Friendship) DeleteUser() *errors.RestErr {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// _, err := collection.DeleteOne(ctx, User{ID: user.ID})
	// if err != nil {
	// 	return errors.NewInternalServerError(err.Error())
	// }
	return nil
}
