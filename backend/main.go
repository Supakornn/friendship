package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Friendship struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Message   string             `json:"message" bson:"message"`
	From      string             `json:"from" bson:"from"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

var (
	client     *mongo.Client
	collection *mongo.Collection
)

// Connect to MongoDB
func connectDB() {
	connectionString := ""
	clientOptions := options.Client().ApplyURI(connectionString)

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("friendshipdb").Collection("friendships")
}

// Get all
func getAllFriendships(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch friendships"})
	}
	defer cursor.Close(ctx)

	var friendships []Friendship
	if err := cursor.All(ctx, &friendships); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to decode friendships"})
	}

	return c.JSON(friendships)
}

// Get by ID
func getFriendshipByID(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid friendship ID"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var friendship Friendship
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&friendship)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Friendship not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch friendship"})
	}

	return c.JSON(friendship)
}

// Create
func createFriendship(c *fiber.Ctx) error {
	friendship := new(Friendship)

	if err := c.BodyParser(friendship); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	friendship.ID = primitive.NewObjectID()
	friendship.CreatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, friendship)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create friendship"})
	}

	return c.Status(fiber.StatusCreated).JSON(friendship)
}

// Delete
func deleteFriendship(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Friendship not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete friendship"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Delete all friendships
func deleteAllFriendships(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete all friendships"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	connectDB()
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/friendships", getAllFriendships)
	app.Get("/friendships/:id", getFriendshipByID)
	app.Post("/friendships", createFriendship)
	app.Delete("/friendships/:id", deleteFriendship)
	app.Delete("/friendships", deleteAllFriendships)

	log.Fatal(app.Listen("0.0.0.0" + getPort()))
}
