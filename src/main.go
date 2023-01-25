package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"src/database"
)

func main() {
	database.ConnectDB()
	database.GetAllUsers()
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
