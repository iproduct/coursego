package main

import (
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursego/03-types-lab2/json/books"
	"log"
	"strings"
)

func main() {
	// Structs --> JSON
	data, err := json.Marshal(books.GoBooks)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// Prettier formatting
	data, err = json.MarshalIndent(books.GoBooks, "", "     ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//// JSON -> structs
	var books []books.Book
	if err := json.Unmarshal(data, &books); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println("\nAFTER UNMARSHAL:")
	for i, book := range books {
		fmt.Printf("%d: %#v\n", i, book)
	}

	// Using json.Encoder
	fmt.Println("\nUSING ENCODER:")
	writer := &strings.Builder{}
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(books)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", writer.String())

	//// JSON -> structs
	decoder := json.NewDecoder(strings.NewReader(writer.String()))
	if err := decoder.Decode(&books); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println("AFTER DECODER UNMARSHAL:")
	for i, book := range books {
		fmt.Printf("%d: %#v\n", i, book)
	}
}
