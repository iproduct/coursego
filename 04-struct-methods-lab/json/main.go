package main

import (
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursegopro/04-struct-methods-lab/json/books"
	"log"
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
	fmt.Println("AFTER UNMARSHAL:\n")
	for i, book := range books {
		fmt.Printf("%d: %#v\n", i, book)
	}
}
