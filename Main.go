/*
	Created on: 30 Jan 2017
	Author: Nilav

	Starts server on default port and registers different handlers
	The mosquitto client is created before starting the server
*/
package main

import (
	"../FutsalApi/programs/controller"
	Mb "../FutsalApi/programs/messageBroker"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	//create mosquitto client initially
	Mb.CreateClient()

	//mux stands for http multiplexer
	//returns a new router for sending the incoming request
	router := mux.NewRouter()

	//HandleFunc registers the handler function for the given pattern in the serverMux
	//add a route for different request
	router.HandleFunc("/user", controller.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/deleteuser/{id}", controller.DeleteUser).Methods("GET")
	router.HandleFunc("/user", controller.AddUser).Methods("POST")
	router.HandleFunc("/user/{id}", controller.UpdateUser).Methods("PUT")

	router.HandleFunc("/login", controller.Login).Methods("POST")

	router.HandleFunc("/payroll/{id}", controller.GetPayrollById).Methods("GET")
	router.HandleFunc("/payrollbyuserid/{id}", controller.GetPayrollById).Methods("GET")

	router.HandleFunc("/payroll", controller.GetPayroll).Methods("GET")
	router.HandleFunc("/payroll", controller.AddPayroll).Methods("POST")
	router.HandleFunc("/payroll/{id}", controller.UpdatePayroll).Methods("PUT")
	router.HandleFunc("/deletepayroll/{id}", controller.DeletePayroll).Methods("GET")
	router.HandleFunc("/deletepayrollbyuserid/{id}", controller.DeletePayroll).Methods("GET")

	router.HandleFunc("/gamedetail", controller.GetGameDetail).Methods("GET")
	router.HandleFunc("/deletegamedetail/{id}", controller.DeleteGameDetail).Methods("GET")
	router.HandleFunc("/gamedetail", controller.AddGameDetail).Methods("POST")
	router.HandleFunc("/gamedetail/{id}", controller.UpdateGameDetail).Methods("PUT")

	//ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.
	http.ListenAndServe(":8080", router)
}
