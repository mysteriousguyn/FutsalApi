/*
	Created on: 3 Feb 2017
	Author: Nilav

	Implementing MQTT client library to create client and for connection to
	MQTT brokers via TCP,TLS or websockets
*/
package messageBroker

import (
	Mqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"log"
	"os"
	"FutsalApi/programs/utility"
)

var client Mqtt.Client

var connHandler Mqtt.OnConnectHandler = func(client Mqtt.Client) {
	fmt.Println("The Connection is successful")
	SubscribeMessage()
}

var lostHandler Mqtt.ConnectionLostHandler = func(client Mqtt.Client, err error) {
	fmt.Printf("The Connection error:=", err)
}

//define a function for the default message handler
// MessageHandler is a callback type which can be set to be
// executed upon the arrival of messages published to topics
// to which the client is subscribed.
var msgHandler Mqtt.MessageHandler = func(client Mqtt.Client, msg Mqtt.Message) {
	//fmt.Println("Received message : topic=%s, message=%s\n", msg.Topic(), msg.Payload()
	utility.ProcessMessage(string(msg.Payload()))

}

/*
 SetDefaultPublishHandler sets the MessageHandler that will be called when a message
 is received that does not match any known subscriptions.
 */
func CreateClient() {

	//create new configuration for client with some default values
	//opts := Mqtt.NewClientOptions().AddBroker("tcp://172.16.14.45:1883")
	opts := Mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("futsalClient")
	opts.SetAutoReconnect(true)
	opts.SetMessageChannelDepth(uint(1024))
	opts.SetOnConnectHandler(connHandler)
	opts.SetConnectionLostHandler(lostHandler)

	//create and start a client using the above clientOptions
	client = Mqtt.NewClient(opts)
	resp := client.Connect()

	if (resp.Wait()&&resp.Error() != nil) {
		fmt.Println("the error while connecting the client is")
		log.Println(resp.Error())
	}
}

/*
Subscribe to the particular topic and request messages to be delivered
at a maximum qos of zero, wait for the receipt to confirm the message
*/
func SubscribeMessage() {

	//subscribe to the particular topic with qos value 2
	//after subscribe the callback function is called whose end value is stored in handler
	token := client.Subscribe("URAPI", 2, msgHandler)

	if (token.Wait()&&token.Error() != nil) {
		fmt.Println("error")
		log.Println(token.Error())
	}
}

//publish messsage to the particular topic
func PublishMessage(message string) {
	fmt.Println(message)
	text := fmt.Sprintf(message)
	token := client.Publish("APIUR", 2, false, text)
	token.Wait()
}

func Unsubscribe() {
	if token := client.Unsubscribe("URAPI"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
}
