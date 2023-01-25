package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"src/emails"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBConfig struct {
	URI        string
	DB_Name    string
	Collection string
}

type Student struct {
	mnr        string
	hash       string
	name       string
	exam_count int32
}

type ExamStats struct {
	exams   int32
	success int32
	failure int32
	wpcount int32
	modules int32
	booked  int32
	mbooked int32
}

func ConnectDB() mongo.Collection {
	// Read DB Config
	config, err := ioutil.ReadFile("database/db_config.json")
	if err != nil {
		log.Fatal(err)
	}
	dbconfig := DBConfig{}
	err = json.Unmarshal(config, &dbconfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("URI: ", dbconfig.URI, "DB Name: ", dbconfig.DB_Name)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbconfig.URI))
	coll := client.Database(dbconfig.DB_Name).Collection(dbconfig.Collection)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return *coll
}

func GetAllUsers() {

	// Download all students
	student_data := ConnectDB()
	filter := bson.D{{"examcount", bson.D{{"$gte", 0}}}}
	cursor, err := student_data.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Student

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Save results in students.json
	for _, result := range results {
		student, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("students.json", student, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CompaireCampusDualStats(student Student) {
	// Download JSON
	var url = "https://selfservice.campus-dual.de/dash/getexamstats?user=" + student.mnr + "&hash=" + student.hash
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// Read Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Save Body
	err = ioutil.WriteFile("examstats.json", body, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	// Compare student.examcount and EXAMS in examstats4003641.json
	byteStats, _ := ioutil.ReadFile("database/examstats.json")
	var cd_stats ExamStats
	json.Unmarshal(byteStats, &cd_stats)
	if student.exam_count == cd_stats.exams {
		emails.SendEmail(student.mnr, student.name)
	}

}
