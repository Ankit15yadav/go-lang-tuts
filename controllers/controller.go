package controllers

import (
	"context"
	"fmt"
	"log"

	"github.com/ankit15yadav/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://yadavankit97620:h9QfMvAqOFY9ACvM@cluster0.l7gow1e.mongodb.net/golang"

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongoDB

func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to mongoDB
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")

}

// MONGODB helpers - file

//insert 1 record

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insertd 1 movie in db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMove(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", result.ModifiedCount)
}

// delete 1 record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie deleted", result.DeletedCount)
}

// delete all record from mongoDB
func deleteAllMovie() {

	filter := bson.D{{}}

	delete, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", delete.DeletedCount)
}

// get all movies from database
func getAllMovies() []bson.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
