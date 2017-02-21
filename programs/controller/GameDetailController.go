/**
 * Created by nilav on 2/17/17.
 */
package controller

import (
	"github.com/structure/futsalStruct"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/RandomGeneration"
)

//method to handle url for adding game detail information
func AddGameDetail(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var gameDetail futsalStruct.GameDetail

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &gameDetail)        //decoding the request body to the structure

	header := futsalStruct.Header{requestId, "GameDetailManagement", "ADD"}
	ParseToJson(header, "", gameDetail)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for getting all game detail information
func GetGameDetail(w http.ResponseWriter, r *http.Request) {

	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier

	header := futsalStruct.Header{requestId, "GameDetailManagement", "SELECT"}
	ParseToJson(header, "", nil)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for updating game detail information based on particular id
func UpdateGameDetail(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()        //generating the globally unique identifier
	var gameDetail futsalStruct.GameDetail

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &gameDetail)        //decoding the request body to the structure

	header := futsalStruct.Header{requestId, "GameDetailManagement", "UPDATE"}
	ParseToJson(header, id, gameDetail)

	fmt.Fprint(w, EndResponse(requestId))
}

//method to handle url for deleting game detail information of particular id
func DeleteGameDetail(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]        //getting the path variable from url
	requestId := RandomGeneration.GenerateRandom()      //generating the globally unique identifier

	header := futsalStruct.Header{requestId, "GameDetailManagement", "DELETE"}
	ParseToJson(header, id, nil)

	fmt.Fprint(w, EndResponse(requestId))
}