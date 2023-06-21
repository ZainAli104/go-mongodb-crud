package controllers

import (
	"context"
	models "go-project/Models"
	"go-project/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser - POST /users
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := config.DB.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting a new user"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetUsers - GET /users
func GetUsers(c *gin.Context) {
	cursor, err := config.DB.Collection("users").Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting users"})
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// GetUser - GET /users/:id
func GetUser(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var user models.User

	err := config.DB.Collection("users").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting a user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser - PUT /users/:id
func UpdateUser(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Collection("users").UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating a user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser - DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))

	_, err := config.DB.Collection("users").DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting a user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
