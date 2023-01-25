package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"src/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()
	database.GetAllUsers()

	router := gin.Default()
	router.POST("/", addStudent)
	router.DELETE("/{matrikelnr}", deleteStudent)
	// TODO: noch ändern wenn server adresse
	router.Run("localhost:9000")

	byteStudents, err := ioutil.ReadFile("database/students.json")
	if err != nil {
		log.Fatal(err)
	}
	students := []database.Student{}
	err = json.Unmarshal(byteStudents, &students)
	for _, student := range students {
		database.CompaireCampusDualStats(student)
	}
}
