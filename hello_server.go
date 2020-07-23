package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Todo struct
type Message struct {
	Appname string `json:"appName"`
	Message string `json:"incomingMessage"`
}

func main() {

	//handleRequests()
	post()
}

func post() {
	fmt.Println("2. Performing Http Post...")

	message := Message{"drgpayer", "140|Metropolitan Jewish Health System Elderplan|ZZzzzzzzzzzzzzzzzzzzzzzzzzzzzzz12343333333333000000000000000"}
	fmt.Println(message)

	jsonReq, err := json.Marshal(message)
	fmt.Println(bytes.NewBuffer(jsonReq))

	resp, err := http.Post("http://mckessonproducer:8081/mckesson/produce", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Message
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}
