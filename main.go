package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	Name  string
	Phone string
}

func dbConnect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// collection := client.Database("test").Collection("trainers")

}

func main() {
	//	route()
	dbConnect()
	routes()

}

// Function to call when end point request
func homePage(w http.ResponseWriter, r *http.Request) {

	// var article Article
	u, res := json.Marshal(Article{Name: "hello", Phone: "welcome"})
	fmt.Println(res)
	fmt.Println(u)
	w.Write(u)

}

func routes() {
	// adding routes using http pacakge
	http.HandleFunc("/", homePage)
	// Start server listening on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
