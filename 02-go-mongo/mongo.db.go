package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define your struct for the data you want to store
type Critter struct {
	Species string
   Name string
	Age  int
	Level int
	Exp int
}

func main() {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    // Disconnect from MongoDB when finished
    defer func() {
        if err = client.Disconnect(context.Background()); err != nil {
            log.Fatal(err)
        }
    }()

    // Access a MongoDB collection
    critterCollection := client.Database("go-learn").Collection("critters")

    // Insert a document
    critterCollection.InsertOne(context.Background(), Critter{"Snorlax", "Alice", 6, 2, 102})
    critterCollection.InsertOne(context.Background(), Critter{"Pikachu", "Kappa", 2, 3, 201})
    critterCollection.InsertOne(context.Background(), Critter{"Mongoose", "Jamal", 3, 1, 21})
    
	 fmt.Println("Made some critters")
}
