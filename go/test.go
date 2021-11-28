package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// The same json tags will be used to encode data into JSON
type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func saveStringIntoFile(filename string, data string) {
	// open the file for writing
	file, _ := os.Create(filename)
	// close the file
	defer file.Close()

	// write the data to the file
	file.WriteString(data)
}

func main() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))

	// save the data to a file
	saveStringIntoFile("./static/testgo.json", string(data))

}
