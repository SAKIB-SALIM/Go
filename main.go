package main

import (
    "main.go/modules/system"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Define a struct to map the JSON
type Config struct {
	Webhook string `json:"webhook"`
}

func main() {
	// Open the JSON file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file contents
	bytes, err := ioutil.ReadAll(file)

	var config Config
	err = json.Unmarshal(bytes, &config)

	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred: %v\n", err)
		}
	}()

	system.Run(config.Webhook)
}
