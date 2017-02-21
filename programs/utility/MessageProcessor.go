package utility

import (
	"encoding/json"
	"log"
	"fmt"
	"github.com/structure/futsalStruct"
)

var res *futsalStruct.Response
var ns map[string]interface{}

func ProcessMessage(message string) {
	//ResNil()			///////////yahan change garya 6

	fmt.Println("Response Message", message)
	//var res Response
	if (message != "") {

		var responseObject map[string]interface{}
		err := json.Unmarshal([]byte(message), &responseObject)
		if err != nil {
			log.Println(err)
		}

		//getting nested object
		jsonRequest := responseObject["Response"].(map[string]interface{})
		jsonHeaderRequest := jsonRequest["Header"].(map[string]interface{})
		res = &futsalStruct.Response{jsonHeaderRequest["RequestId"].(string), jsonHeaderRequest["Module"].(string),
			jsonHeaderRequest["Type"].(string), jsonRequest["Body"].(map[string]interface{})}

		fmt.Println(res.ResponseModule)

	}
}

func ResNil() {
	ns = map[string]interface{}{"Message":"Error"}
	res.ResponseMessage = ns
	res.ResponseRequestId=""
}

func SendMessage() *futsalStruct.Response {

	if (res == nil) {

		res = &futsalStruct.Response{"", "", "", ns}
		return res
	}
	return res
}

