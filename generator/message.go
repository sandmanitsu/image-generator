package main

import (
	"encoding/json"
	"log"
	"os"
)

type Data struct {
	Data []Message
}

type Message struct {
	Template string          `json:"template"`
	Payload  json.RawMessage `json:"payload"`
}

// type Payload struct {
// 	Size  string `json:"size"`
// 	Price string `json:"price"`
// 	Type  string `json:"type"`
// 	File  string `json:"file"`
// }

func CreateMessage() Data {
	raw, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal("error read data.json file", err)
	}

	var messages Data
	if err := json.Unmarshal(raw, &messages); err != nil {
		log.Fatal("error unmarshal json", err)
	}

	return messages
}
