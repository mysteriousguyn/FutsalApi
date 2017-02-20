/**
 * Created by nilav on 2/17/17.
 */
package controller

import (
	"github.com/structure/futsalStruct"
	"../service"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func AddGameDetail(w http.ResponseWriter, r *http.Request) {

	var gameDetail futsalStruct.GameDetail
	requestBody, _ := ioutil.ReadAll(r.Body)

	//decoding the request body
	json.Unmarshal(requestBody, &gameDetail)

	response := service.AddGameDetail(gameDetail)

	fmt.Fprint(w, response)
}


func GetGameDetail(w http.ResponseWriter, r *http.Request) {

	response := service.GetGameDetail()

	fmt.Fprint(w, response)
}


func UpdateGameDetail(w http.ResponseWriter, r *http.Request) {

	//getting the path variable from url
	id := mux.Vars(r)["id"]
	var gameDetail futsalStruct.GameDetail
	requestBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(requestBody, &gameDetail)
	response := service.UpdateGameDetail(id, gameDetail)

	fmt.Fprint(w, response)
}

func DeleteGameDetail(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	response := service.DeleteGameDetail(id)

	fmt.Fprint(w, response)
}