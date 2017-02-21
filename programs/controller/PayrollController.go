/*
	Created on: 1 Feb 2017
	Author: Nilav

	PayrollController receives a request related to Payroll
	and delegates tasks to be performed based on the parameters

*/
package controller

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/structure/futsalStruct"
	"io/ioutil"
	"encoding/json"
	"strings"
	"github.com/RandomGeneration"
)

//method to handle url for adding payroll detail information
func AddPayroll(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var payroll futsalStruct.Payroll

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &payroll)        //decoding the request body to the structure

	header := futsalStruct.Header{requestId, "PayrollManagement", "ADD"}
	ParseToJson(header, "", payroll)

	fmt.Fprintf(w, EndResponse(requestId))
}

//method to handle url for retrieving all payroll detail information
func GetPayroll(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier

	header := futsalStruct.Header{requestId, "PayrollManagement", "SELECT"}
	ParseToJson(header, "", nil)

	fmt.Fprintf(w, EndResponse(requestId))
}

/*
method to handle url for retrieving payroll detail information of particular id
or user id
*/
func GetPayrollById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var header futsalStruct.Header

	uriSegment := strings.Split(r.URL.Path, "/")        //separates the uri on /
	api := uriSegment[1]        //get the second substring after separation

	if (api == "payroll") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "SELECT BY ID"}
	} else if (api == "payrollbyuserid") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "SELECT BY USERID"}
	}
	ParseToJson(header, id, nil)

	fmt.Fprintf(w, EndResponse(requestId))
}

//method to handle url for updating payroll detail information of particular id
func UpdatePayroll(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var payroll futsalStruct.Payroll

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &payroll)        //decoding the request body to the structure

	header := futsalStruct.Header{requestId, "PayrollManagement", "UPDATE"}
	ParseToJson(header, id, payroll)

	fmt.Fprintf(w, EndResponse(requestId))
}

/*
method to handle url for deleting payroll detail information of particular id
or user id
*/
func DeletePayroll(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()
	var header futsalStruct.Header

	uriSegment := strings.Split(r.URL.Path, "/")        //separates the uri on "/"
	api := uriSegment[1]        //get the second substring after separation

	if (api == "deletepayroll") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "DELETE"}
	} else if (api == "deletepayrollbyuserid") {
		header = futsalStruct.Header{requestId, "PayrollManagement", "DELETE BY USERID"}
	}
	ParseToJson(header, id, nil)

	fmt.Fprintf(w, EndResponse(requestId))
}
