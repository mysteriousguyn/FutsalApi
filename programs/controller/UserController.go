/*
	Created on: 31 Jan 2017
	Author: Nilav

	UserController receives a request related to Users
	and delegates tasks to be performed based on the parameters

*/
package controller

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"

	"io/ioutil"
	"github.com/structure/futsalStruct"
	"encoding/json"
	"github.com/RandomGeneration"
)

//method to handle url for adding user information
func AddUser(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var user futsalStruct.User

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &user)        //decoding the request body to the structure

	header := futsalStruct.Header{requestId, "UserManagement", "ADD"}
	ParseToJson(header, "", user)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for retrieving all user information
func GetUser(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier

	header := futsalStruct.Header{requestId, "UserManagement", "SELECT"}
	ParseToJson(header, "", nil)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for retrieving user information of particular id
func GetUserById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier

	header := futsalStruct.Header{requestId, "UserManagement", "SELECT BY ID"}
	ParseToJson(header, id, nil)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for updating user information of particular id
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var user futsalStruct.User

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &user)

	header := futsalStruct.Header{requestId, "UserManagement", "UPDATE"}
	ParseToJson(header, id, user)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for deleting user information of particular id
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]      //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier

	header := futsalStruct.Header{requestId, "UserManagement", "DELETE"}
	ParseToJson(header, id, nil)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for login and checking the corressponding email and password
func Login(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var user futsalStruct.User

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &user)        //decoding the request body to the structure

	header := futsalStruct.Header{requestId, "UserManagement", "LOGIN"}
	ParseToJson(header, "", user)

	fmt.Fprint(w, EndResponse(requestId))
}
