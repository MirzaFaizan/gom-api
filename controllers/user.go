package controllers

import (
	"context"
	"time"

	Config "github.com/mirzafaizan/gom-api/config"
	Models "github.com/mirzafaizan/gom-api/models"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DB connection
var DB = Config.DB().Database("usergo").Collection("profiles")

func handleErr(ctx iris.Context, err error) {
	ctx.JSON(iris.Map{"response": err.Error()})
}

// GetAllUsers ...
// Method:   GET
// Resource: this to get all all users
func GetAllUsers(ctx iris.Context) {
	results := []*Models.User{}
	c := context.TODO()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := DB.Find(c, bson.D{{}})
	if err != nil {
		handleErr(ctx, err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(c) {
		// create a value into which the single document can be decoded
		var elem Models.User
		err := cur.Decode(&elem)
		if err != nil {
			handleErr(ctx, err)
			return
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		handleErr(ctx, err)
		return
	}
	// Close the cursor once finished
	cur.Close(c)
	ctx.JSON(iris.Map{"response": results})
}

//GetUser : function to get user
//Method : GET
// Resource: this to get all all users
func GetUser(ctx iris.Context) {
	msisdn := ctx.Params().Get("msisdn")

	result := Models.User{}
	err := DB.FindOne(context.TODO(), bson.D{primitive.E{Key: "msisdn", Value: msisdn}}).Decode(&result)
	if err != nil {
		handleErr(ctx, err)
		return
	}
	ctx.JSON(iris.Map{"response": result})
}

//CreateUser : To create new user in DB
// Method:   POST
// Resource: This is to create a new user
func CreateUser(ctx iris.Context) {
	params := &Models.User{}
	err := ctx.ReadJSON(params)
	if err != nil {
		handleErr(ctx, err)
		return
	}

	params.LastUpdate = time.Now()
	result, err := DB.InsertOne(context.TODO(), params)
	if err != nil {
		handleErr(ctx, err)
		return
	}
	ctx.JSON(iris.Map{"response": "User successfully created", "message": result})
}

//UpdateUser : function to update user
// Method:  PATCH
// Resource: This is to update a user record
func UpdateUser(ctx iris.Context) {
	msisdn := ctx.Params().Get("msisdn")
	params := &Models.User{}
	err := ctx.ReadJSON(params)
	if err != nil {
		handleErr(ctx, err)
		return
	}
	params.InsertedAt = time.Now()
	query := bson.D{primitive.E{Key: "msisdn", Value: msisdn}}
	updatedResult, err := DB.UpdateOne(context.TODO(), query, params)
	if err != nil {
		handleErr(ctx, err)
		return
	}
	ctx.JSON(iris.Map{"response": "user record successfully updated", "data": updatedResult})
}

//DeleteUser : delete user
// Method:   DELETE
// Resource: This is to delete a user record
func DeleteUser(ctx iris.Context) {
	msisdn := ctx.Params().Get("msisdn")

	params := &Models.User{}
	err := ctx.ReadJSON(params)
	if err != nil {
		handleErr(ctx, err)
		return
	}

	params.InsertedAt = time.Now()
	query := bson.D{primitive.E{Key: "msisdn", Value: msisdn}}
	deleteResult, err := DB.DeleteOne(context.TODO(), query)
	if err != nil {
		handleErr(ctx, err)
		return
	}
	ctx.JSON(iris.Map{"response": "user record successfully deleted", "result": deleteResult})
}
