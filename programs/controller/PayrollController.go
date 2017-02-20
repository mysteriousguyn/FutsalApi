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
	"../service"
	"github.com/gorilla/mux"
	"github.com/structure/futsalStruct"
	"io/ioutil"
	"encoding/json"
	"strings"
)

func AddPayroll(w http.ResponseWriter, r *http.Request) {

	var payroll futsalStruct.Payroll
	requestBody, _ := ioutil.ReadAll(r.Body)

	//decoding the request body
	json.Unmarshal(requestBody, &payroll)

	response := service.AddPayrollInfo(payroll)

	fmt.Fprintf(w, response)

}

func GetPayroll(w http.ResponseWriter, r *http.Request) {

	response := service.GetPayrollInfo()

	fmt.Fprintf(w, response)

}

func GetPayrollById(w http.ResponseWriter, r *http.Request) {

	//getting the path variable from url
	id := mux.Vars(r)["id"]
	fmt.Println("The id of payroll is:=", id)

	uriSegment := strings.Split(r.URL.Path, "/")
	api := uriSegment[1]
	fmt.Println(api)
	response := service.GetPayrollInfoById(id, api)

	fmt.Fprintf(w, response)

}

func UpdatePayroll(w http.ResponseWriter, r *http.Request) {

	//getting the path variable from url
	id := mux.Vars(r)["id"]

	var payroll futsalStruct.Payroll
	requestBody, _ := ioutil.ReadAll(r.Body)

	//decode the request body
	json.Unmarshal(requestBody, &payroll)
	response := service.UpdatePayroll(id, payroll)

	fmt.Fprintf(w, response)
}

func DeletePayroll(w http.ResponseWriter, r *http.Request) {

	//getting the path variable from url
	id := mux.Vars(r)["id"]
	uriSegment := strings.Split(r.URL.Path, "/")
	api := uriSegment[1]
	response := service.DeletePayroll(id, api)

	fmt.Fprintf(w, response)
}
