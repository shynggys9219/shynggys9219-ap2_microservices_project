package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var collection *mongo.Collection

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `json:"email"`
	Password string             `json:"password,omitempty"`
	Hashed   string             `bson:"hashed"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("mongo.Connect: %v", err)
	}
	collection = client.Database("userdb").Collection("users")

	r := gin.Default()

	r.POST("/register", registerHandler)
	r.POST("/login", loginHandler)

	r.Run(":8001")
}

func registerHandler(c *gin.Context) {
	var input User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Hashed = string(hashed)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.M{"email": input.Email, "hashed": input.Hashed})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "registration failed"})
		return
	}

	// Notify statistics service
	_, err = http.Post("http://localhost:8002/increment", "application/json", nil)
	if err != nil {
		fmt.Printf("http.Post: statistics: %v, status: %v\n", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

func loginHandler(c *gin.Context) {
	var input User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err := collection.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Hashed), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
