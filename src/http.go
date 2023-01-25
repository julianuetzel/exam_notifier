package main

import (
	"context"
	"log"
	"net/http"
	"src/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func addStudent(c *gin.Context) {
	var newStudent database.Student

	// Bind recieved JSON to newStudent
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	// Save in Database
	coll := database.ConnectDB()
	_, err := coll.InsertOne(context.TODO(), newStudent)
	if err != nil {
		log.Fatalln(err)
	}

	// Return HTTP-Status
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func deleteStudent(c *gin.Context) {
	mnr := c.Param("matrikelnr")

	coll := database.ConnectDB()
	filter := bson.D{{"mnr", mnr}}

	// Look for mnr in Database and delete
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatalln(err)
	}
	if result != nil {
		c.IndentedJSON(http.StatusOK, result)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student not found"})
}
