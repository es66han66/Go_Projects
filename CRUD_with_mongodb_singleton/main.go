package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"example.com/eshan/DB"
	"example.com/eshan/models"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = DB.ConnectDB()

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []models.Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		DB.GetError(err, response)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person models.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person models.Person
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// err := collection.FindOne(ctx, models.Person{ID: id}).Decode(&person) // it also works
	// err := collection.FindOne(ctx, models.Person{Firstname: "Eshan"}).Decode(&person) // it also works
	// err := collection.FindOne(ctx, bson.M{"firstname": "Eshan"}).Decode(&person) // it also works
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&person) // it also works

	if err != nil {
		DB.GetError(err, response)

		return
	}
	json.NewEncoder(response).Encode(person)
}

func DeletePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		DB.GetError(err, response)

		return
	}
	json.NewEncoder(response).Encode(result)
}

func UpdatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person models.Person
	json.NewDecoder(request.Body).Decode(&person)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$set": person,
	})
	if err != nil {
		DB.GetError(err, response)

		return
	}
	json.NewEncoder(response).Encode(result)
}

func main() {
	//Init Router
	router := mux.NewRouter()

	// arrange our route
	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/person/{id}", UpdatePersonEndpoint).Methods("PUT")
	fmt.Println("Starting the application at port 8000...")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", router))

}
