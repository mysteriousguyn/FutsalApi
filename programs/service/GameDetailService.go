/**
 * Created by nilav on 2/20/17.
 */
package service

import (
	"encoding/json"
	"fmt"
	Mb"../messageBroker"
	"github.com/structure/futsalStruct"
	"github.com/RandomGeneration"
	"../utility"
	"log"
)

func AddGameDetail(gamedetail futsalStruct.GameDetail) string {

	requestId := RandomGeneration.GenerateRandom()

	header := futsalStruct.Header{requestId, "GameDetailManagement", "ADD"}

	jsonFormat := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
			{"Body", gamedetail},
		},
	}
	//creating the json of the request bpdy and header
	jsonValue, _ := json.Marshal(jsonFormat)

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()
	utility.ResNil()
	jsonMessage, err := json.Marshal(res.ResponseMessage)

	if (err != nil) {
		log.Fatal(err)
	}

	fmt.Println("what the = ", string(jsonMessage))

	message := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (message == "Error") {
		res = utility.SendMessage()
		message = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "GameDetailManagement"&&
		res.ResponseType == "ADD") {

		jsonMessage, err = json.Marshal(res.ResponseMessage)

		if (err != nil) {
			log.Fatal(err)
		}

		return string(jsonMessage)
	}
	return string(jsonMessage)
}

func GetGameDetail() string {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "GameDetailManagement", "SELECT"}

	//creates the format for Json
	jsonFormat := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
		},
	}

	//creating the json
	jsonValue, err := json.Marshal(jsonFormat)
	if (err != nil) {
		log.Println(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()
	utility.ResNil()
	jsonMessage, err := json.Marshal(res.ResponseMessage)

	if (err != nil) {
		log.Println(err)
	}

	message := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (message == "Error") {
		res = utility.SendMessage()
		message = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "GameDetailManagement"&&
		res.ResponseType == "SELECT") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		jsonMessage, err = json.Marshal(res.ResponseMessage)

		if (err != nil) {
			log.Println(err)
		}
		return string(jsonMessage)
	}
	return string(jsonMessage)
}

func UpdateGameDetail(id string, gameDetail futsalStruct.GameDetail) string {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "GameDetailManagement", "UPDATE"}

	//creates the format for Json
	jsonFormat := map[string]interface{}{

		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
			{"Body", gameDetail},
		},
	}

	//creating the json of the request bpdy and header
	jsonValue, err := json.Marshal(jsonFormat)
	if (err != nil) {
		log.Println(err)
	}

	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()

	utility.ResNil()

	jsonMessage, err := json.Marshal(res.ResponseMessage)
	if (err != nil) {
		log.Println(err)
	}

	message := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (message == "Error") {
		res = utility.SendMessage()
		message = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "GameDetailManagement"&&
		res.ResponseType == "UPDATE") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		jsonMessage, err = json.Marshal(res.ResponseMessage)
		if (err != nil) {
			log.Println(err)
		}

		return string(jsonMessage)
	}
	return string(jsonMessage)
}

func DeleteGameDetail(id string) interface{} {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "GameDetailManagement", "DELETE"}

	//creates the format for Json
	jsonFormat := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
		},
	}

	//creating the json
	jsonValue, err := json.Marshal(jsonFormat)
	if (err != nil) {
		log.Fatal(err)
	}

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()

	utility.ResNil()

	jsonMessage, err := json.Marshal(res.ResponseMessage)
	if (err != nil) {
		log.Fatal(err)
	}


	message := res.ResponseMessage["Message"]

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (message == "Error") {
		res = utility.SendMessage()
		message = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "GameDetailManagement"&&
		res.ResponseType == "DELETE") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		jsonMessage, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(jsonMessage)
	}
	return string(jsonMessage)
}