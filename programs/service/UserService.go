/*
	Created on: 31 Jan 2017
	Author: Nilav

	UserService creates a new JSON with request body and extra parameters
	Publish message to the message Broker
	Subscribe response from the message Broker
*/
package service

import (
	"encoding/json"
	"github.com/structure/futsalStruct"
	"github.com/RandomGeneration"
	Mb"../messageBroker"
	"../utility"

	"fmt"
	"log"
)

func AddUser(user futsalStruct.User) string {
	requestId := RandomGeneration.GenerateRandom()

	header := futsalStruct.Header{requestId, "UserManagement", "ADD"}

	//creates the format for Json
	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
			{"Body", user},
		},

	}
	//creates JSON
	jsonValue, err := json.Marshal(format)

	if (err != nil) {
		log.Println(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()
	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)

	val := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]
	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "UserManagement"&&
		res.ResponseType == "ADD") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}

func GetUser() string {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "UserManagement", "SELECT"}

	//creates the format for Json
	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
		},
	}

	//creating the json
	jsonValue, err := json.Marshal(format)
	if (err != nil) {
		log.Println(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()
	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)
	val := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "UserManagement"&&
		res.ResponseType == "SELECT") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}

func GetUserById(id string) string {
	requestId := RandomGeneration.GenerateRandom()

	header := futsalStruct.Header{requestId, "UserManagement", "SELECT BY ID"}

	//creates the format for Json
	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
		},
	}

	//creating the json
	jsonValue, err := json.Marshal(format)
	if (err != nil) {
		log.Println(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()

	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)
	val := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "UserManagement"&&
		res.ResponseType == "SELECT BY ID") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}

func UpdateUserById(id string, user futsalStruct.User) string {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "UserManagement", "UPDATE"}

	//creates the format for Json
	format := map[string]interface{}{

		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
			{"Body", user},
		},
	}

	//creating the json of the request bpdy and header
	jsonValue, err := json.Marshal(format)
	if (err != nil) {
		log.Println(err)
	}

	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()

	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)
	val := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "UserManagement"&&
		res.ResponseType == "UPDATE") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}

func DeleteUserById(id string) interface{} {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "UserManagement", "DELETE"}

	//creates the format for Json
	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
		},
	}

	//creating the json
	jsonValue, err := json.Marshal(format)
	if (err != nil) {
		log.Println(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()

	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)
	val := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "UserManagement"&&
		res.ResponseType == "DELETE") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}

func LoginUser(user futsalStruct.User) interface{} {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "UserManagement", "LOGIN"}

	//creates the format for Json
	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
			{"Body", user},
		},
	}

	//creating the json
	jsonValue, err := json.Marshal(format)
	if (err != nil) {
		log.Println(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()
	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)
	val := res.ResponseMessage["Message"]

	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]
	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "UserManagement"&&
		res.ResponseType == "LOGIN") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}