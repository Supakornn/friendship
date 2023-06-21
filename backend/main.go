package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FriendShip struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Message   string             `json:"message"`
	From      string             `json:"from"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
}

var friendshipCollection *mongo.Collection

func main() {
	// ENV
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatalf("Some error occured. Err: %s", enverr)
	}
	val := os.Getenv("DATABASE")

	// MongoDB
	clientOptions := options.Client().ApplyURI(val)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	friendshipCollection = client.Database("my-friendship").Collection("friendship")

	// router
	router := mux.NewRouter()
	router.HandleFunc("/create", createTodo).Methods("POST")

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://example.com", "http://localhost:5173"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(router)

	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", c))

}

// Create
func createTodo(w http.ResponseWriter, r *http.Request) {
	var friendship FriendShip
	json.NewDecoder(r.Body).Decode(&friendship)
	friendship.CreatedAt = time.Now()

	result, err := friendshipCollection.InsertOne(context.TODO(), friendship)
	if err != nil {
		log.Fatal(err)
	}

	friendship.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(friendship)
}
