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
	"../service"
	"github.com/gorilla/mux"

	"io/ioutil"
	"github.com/structure/futsalStruct"
	"encoding/json"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	var user futsalStruct.User
	requestBody, _ := ioutil.ReadAll(r.Body)

	//decoding the request body
	json.Unmarshal(requestBody, &user)

	response := service.AddUser(user)

	fmt.Fprint(w, response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	response := service.GetUser()

	fmt.Fprint(w, response)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {

	//getting the path variable from url
	id := mux.Vars(r)["id"]
	response := service.GetUserById(id)

	fmt.Fprint(w, response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	//getting the path variable from url
	id := mux.Vars(r)["id"]
	var user futsalStruct.User
	requestBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(requestBody, &user)
	response := service.UpdateUserById(id, user)

	fmt.Fprint(w, response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	response := service.DeleteUserById(id)

	fmt.Fprint(w, response)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var user futsalStruct.User
	requestBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(requestBody, &user)
	response := service.LoginUser(user)
	fmt.Println("the response is=",response)
	fmt.Fprint(w, response)


}
