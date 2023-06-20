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

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title"`
	Completed bool               `json:"completed"`
	CreatedAt time.Time          `json:"createdAt,omitempty"`
}

var todosCollection *mongo.Collection

func main() {
	// ENV
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatalf("Some error occured. Err: %s", enverr)
	}
	val := os.Getenv("DATABASE")
	println(val)

	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI(val)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	todosCollection = client.Database("todoapp").Collection("todos")

	// Set up HTTP server with Gorilla Mux
	router := mux.NewRouter()
	router.HandleFunc("/api/todos", createTodo).Methods("POST")

	// Create CORS options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://example.com", "http://localhost:5173"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(router)

	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", c))

}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.CreatedAt = time.Now()

	result, err := todosCollection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatal(err)
	}

	todo.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(w).Encode(todo)
}
