/**
 * Created by nilav on 2/20/17.
 */
package controller

import (
	"encoding/json"
	"github.com/structure/futsalStruct"
	Mb"../messageBroker"
	"log"
	"../utility"
)

/*
it get the response message and encode the message in json format
takes request id as a parameter
*/
func EndResponse(requestId string) string {

	res := utility.SendMessage()
	utility.ResNil()        //default struct value

	rs, err := json.Marshal(res.ResponseMessage)
	if (err != nil) {
		log.Fatal(err)
	}

	val := res.ResponseRequestId

	//the application will wait until the topic subscribed contains the message
	for (val == "") {
		res = utility.SendMessage()
		val = res.ResponseRequestId
	}

	if (res.ResponseRequestId == requestId) {
		rs, _ = json.Marshal(res.ResponseMessage)
		return string(rs)
	}
	return string(rs)
}

/*
encode the header, structure and/or id to json format
it takes header struct, default id and the structure as a parameter
*/
func ParseToJson(header futsalStruct.Header, id string, domain interface{}) {

	var jsonFormat interface{}

	if (id == "") {
		if (domain == nil) {
			jsonFormat = map[string]interface{}{
				"Request":utility.JsonOrder{
					{"Header", header},
				},
			}
		}

		if (domain != nil) {

			jsonFormat = map[string]interface{}{
				"Request":utility.JsonOrder{
					{"Header", header},
					{"Body", domain},
				},
			}
		}
	} else {
		if (domain == nil) {
			jsonFormat = map[string]interface{}{
				"Request":utility.JsonOrder{
					{"Id", id},
					{"Header", header},
				},
			}
		}

		if (domain != nil) {

			jsonFormat = map[string]interface{}{

				"Request":utility.JsonOrder{
					{"Id", id},
					{"Header", header},
					{"Body", domain},
				},
			}
		}
	}

	jsonValue, err := json.Marshal(jsonFormat)	//creating the json

	if (err != nil) {
		log.Fatal(err)
	}

	Mb.PublishMessage(string(jsonValue))
}