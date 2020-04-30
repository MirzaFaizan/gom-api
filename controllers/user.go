package controllers

import (
	Bcontext "context"
	"log"
	"time"

	Config "github.com/mirzafaizan/gom-api/config"
	Models "github.com/mirzafaizan/gom-api/models"

	"github.com/kataras/iris/context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DB connection
var DB = Config.DB().Database("usergo").Collection("profiles")

// GetAllUsers ...
// Method:   GET
// Resource: this to get all all users
func GetAllUsers(ctx context.Context) {
	results := []*Models.User{}
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := DB.Find(Bcontext.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(Bcontext.TODO()) {
		// create a value into which the single document can be decoded
		var elem Models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// Close the cursor once finished
	cur.Close(Bcontext.TODO())
	ctx.JSON(context.Map{"response": results})
}

//GetUser : function to get user
//Method : GET
// Resource: this to get all all users
func GetUser(ctx context.Context) {
	msisdn := ctx.Params().Get("msisdn")
	if msisdn == "" {
		ctx.JSON(context.Map{"response": "please pass a valid msisdn"})
	}
	result := Models.User{}
	err := DB.FindOne(Bcontext.TODO(), bson.D{primitive.E{Key: "msisdn", Value: msisdn}}).Decode(&result)
	if err != nil {
		ctx.JSON(context.Map{"response": err.Error()})
	}
	ctx.JSON(context.Map{"response": result})
}

//CreateUser : To create new user in DB
// Method:   POST
// Resource: This is to create a new user
func CreateUser(ctx context.Context) {
	params := &Models.User{}
	err := ctx.ReadJSON(params)
	if err != nil {
		ctx.JSON(context.Map{"response": err.Error()})
	} else {
		params.LastUpdate = time.Now()
		result, err := DB.InsertOne(Bcontext.TODO(), params)
		if err != nil {
			ctx.JSON(context.Map{"response": "User succesfully created", "message": err})
		}
		ctx.JSON(context.Map{"response": "User succesfully created", "message": result})
	}
}

//UpdateUser : function to update user
// Method:  PATCH
// Resource: This is to update a user record
func UpdateUser(ctx context.Context) {
	msisdn := ctx.Params().Get("msisdn")
	if msisdn == "" {
		ctx.JSON(context.Map{"response": "please pass a valid msisdn"})
	}
	params := &Models.User{}
	err := ctx.ReadJSON(params)
	if err != nil {
		ctx.JSON(context.Map{"response": err.Error()})
	} else {
		params.InsertedAt = time.Now()
		query := bson.D{primitive.E{Key: "msisdn", Value: msisdn}}
		updatedResult, err := DB.UpdateOne(Bcontext.TODO(), query, params)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		} else {
			ctx.JSON(context.Map{"response": "user record successfully updated", "data": updatedResult})
		}
	}

}

//DeleteUser : delete user
// Method:   DELETE
// Resource: This is to delete a user record
func DeleteUser(ctx context.Context) {
	msisdn := ctx.Params().Get("msisdn")
	if msisdn == "" {
		ctx.JSON(context.Map{"response": "please pass a valid msisdn"})
	}
	params := &Models.User{}
	err := ctx.ReadJSON(params)
	if err != nil {
		ctx.JSON(context.Map{"response": err.Error()})
	} else {
		params.InsertedAt = time.Now()
		query := bson.D{primitive.E{Key: "msisdn", Value: msisdn}}
		deleteResult, err := DB.DeleteOne(Bcontext.TODO(), query)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		} else {
			ctx.JSON(context.Map{"response": "user record successfully deleted", "result": deleteResult})
		}
	}

}
