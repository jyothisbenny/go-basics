package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dbURI = "mongodb://localhost:27017"

var collection *mongo.Collection

func MongoInit() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	fmt.Println("Connected to MongoDB!")

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to MongoDB: %w", err)
	}

	fmt.Println("MongoDB Tested!")

	return client, nil
}

// Init is automatically called when the package is imported
func Init() {
	var err error
	client, err := MongoInit()
	if err != nil {
		log.Fatalf("Error initializing MongoDB: %v", err)
	}
	collection = client.Database("test").Collection("test")
	fmt.Println("MongoDB initialization complete in model package!")
}

func main() {
	fmt.Println("inside model main!")
}

// Helpers

// insert one record
func InsetOneMovie(movie Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted movie with id: %v \n", inserted.InsertedID)
}

func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated movie with id:", result.ModifiedCount)

}

func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted movie with id:", movieId)

}

func DeleteAllMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted:", result.DeletedCount)
	return result.DeletedCount
}

func GetAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			fmt.Println(err)
		}
		movies = append(movies, movie)
	}
	return movies
}
