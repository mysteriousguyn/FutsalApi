/*
	Created on: 2 Feb 2017
	Author: Nilav

	PayrollService creates a new JSON for payroll info with request body and extra parameters
	Publish message to the message Broker
	Subscribe response from the message Broker
*/
package service

import (
	"encoding/json"
	"fmt"
	"github.com/structure/futsalStruct"
	Mb"../messageBroker"
	"../utility"
	"reflect"
	"github.com/RandomGeneration"
	"log"
)

//function to publish the new payroll information to mosquitto broker
//response is received subscribing the particular topic
func AddPayrollInfo(payroll futsalStruct.Payroll) string {

	requestId := RandomGeneration.GenerateRandom()

	header := futsalStruct.Header{requestId, "PayrollManagement", "ADD"}

	jsonFormat := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
			{"Body", payroll},
		},
	}
	//creating the json of the request bpdy and header
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

	fmt.Println("what the = ", string(jsonMessage))

	val := res.ResponseMessage["Message"]
	fmt.Println("what the = ", string(jsonMessage))
	fmt.Println(reflect.ValueOf(val).Kind() != reflect.Map)

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "PayrollManagement"&&
		res.ResponseType == "ADD") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		jsonMessage, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(jsonMessage)
	}
	return string(jsonMessage)
}

//function to retrieve the information of the payroll info in json format
func GetPayrollInfo() string {

	requestId := RandomGeneration.GenerateRandom()
	header := futsalStruct.Header{requestId, "PayrollManagement", "SELECT"}

	jsonFormat := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Header", header},
		},
	}

	//creating the json of the request bpdy and header
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

	fmt.Println("what the = ", string(jsonMessage))

	message := res.ResponseMessage["Message"]
	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (message == "Error") {
		res = utility.SendMessage()
		message = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "PayrollManagement"&&
		res.ResponseType == "SELECT") {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		jsonMessage, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(jsonMessage)
	}
	return string(jsonMessage)
}

//func to retrieve the information of payroll based on user Id
func GetPayrollInfoById(id string, api string) string {

	requestId := RandomGeneration.GenerateRandom()
	var header futsalStruct.Header

	if (api == "payroll") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "SELECT BY ID"}
	} else if (api == "payrollbyuserid") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "SELECT BY USERID"}
	}

	jsonFormat := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
		},
	}

	//creating the json of the request bpdy and header
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

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "PayrollManagement"&&
		(res.ResponseType == "SELECT BY ID" || res.ResponseType == "SELECT BY USERID")) {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		jsonMessage, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(jsonMessage)
	}
	return string(jsonMessage)
}

//func to update the user information
func UpdatePayroll(id string, payroll futsalStruct.Payroll) string {

	requestId := RandomGeneration.GenerateRandom()

	header := futsalStruct.Header{requestId, "PayrollManagement", "UPDATE"}

	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
			{"Body", payroll},
		},
	}

	headerJson, _ := json.Marshal(format)

	Mb.PublishMessage(string(headerJson))

	res := utility.SendMessage()
	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)

	fmt.Println("what the = ", string(rs))

	val := res.ResponseMessage["Message"]
	fmt.Println("what the = ", string(rs))
	fmt.Println(reflect.ValueOf(val).Kind() != reflect.Map)

	//for (reflect.ValueOf(val).Kind()!=reflect.Map&&val!="Email already exists"){
	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "PayrollManagement"&&
		res.ResponseType == "UPDATE" ) {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}

//func to delete the user information based on id
func DeletePayroll(id string, api string) string {

	requestId := RandomGeneration.GenerateRandom()
	var header futsalStruct.Header

	if (api == "deletepayroll") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "DELETE"}
	} else if (api == "deletepayrollbyuserid") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "DELETE BY USERID"}
	}
	//creating the json of the request bpdy and header
	format := map[string]interface{}{
		"Request":utility.JsonOrder{
			{"Id", id},
			{"Header", header},
		},
	}
	//creating the json of the request bpdy and header
	jsonValue, _ := json.Marshal(format)

	//publish message on the particular topic in mosquitto
	Mb.PublishMessage(string(jsonValue))

	res := utility.SendMessage()
	utility.ResNil()
	rs, _ := json.Marshal(res.ResponseMessage)

	fmt.Println("what the = ", string(rs))

	val := res.ResponseMessage["Message"]
	fmt.Println("what the = ", string(rs))
	fmt.Println(reflect.ValueOf(val).Kind() != reflect.Map)

	for (val == "Error") {
		res = utility.SendMessage()
		val = res.ResponseMessage["Message"]

	}

	if (res.ResponseRequestId == requestId&&res.ResponseModule == "PayrollManagement"&&
		(res.ResponseType == "DELETE" || res.ResponseType == "DELETE BY USERID")) {

		fmt.Println("Inside here is request Id=", res.ResponseRequestId)
		rs, _ = json.Marshal(res.ResponseMessage)
		//utility.ResNil()
		return string(rs)
	}
	return string(rs)
}