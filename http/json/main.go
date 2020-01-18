package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main()  {
	// Structs --> JSON
	data, err := json.Marshal(goBooks)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// Prettier formatting
	data, err = json.MarshalIndent(goBooks, "", "     ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// JSON -> structs
	var books []Book
	if err := json.Unmarshal(data, &books); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println("AFTER UNMARSHAL:\n", books) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}