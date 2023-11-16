package main

import (
	"encoding/json"
	"fmt"
	"github.com/iproduct/coursego/fmi-2023-04-methods-interfaces-lab/books"
	"log"
)

var jsonText = `[{
      "kind": "books#volume",
      "id": "ba4IEAAAQBAJ",
      "etag": "gXZAfne9N7A",
      "selfLink": "https://www.googleapis.com/books/v1/volumes/ba4IEAAAQBAJ",
      "volumeInfo": {
        "title": "GO Programming in easy steps",
        "subtitle": "Discover Google’s Go language (golang)",
        "authors": [
          "Mike McGrath"
        ],
        "publisher": "In Easy Steps Limited",
        "publishedDate": "2020-11-13",
        "description": "GO Programming in easy steps has an easy-to-follow style that will appeal to anyone who wants to begin coding computer programs with Google’s Go programming language. The code in the listed steps within the book is color-coded making it easier for beginners to grasp. You need have no previous knowledge of any computer programming language so it's ideal for the newcomer. GO Programming in easy steps instructs you how to write code to create your own computer programs. It contains separate chapters demonstrating how to store information in data structures, how to control program flow using control structures, and how to create re-usable blocks of code in program functions. There are complete step-by-step example programs that demonstrate each aspect of coding, together with screenshots that illustrate the actual output when each program is executed. GO Programming in easy steps begins by explaining how to easily create a programming environment on your own computer, so you can quickly begin to create your own working programs by copying the book's examples. After demonstrating the essential building blocks of computer programming it describes how to use data abstraction for object-oriented programming and demonstrates how to code goroutines and channels for concurrency in your programs. Table of Contents 1. Get Started 2. Store Values 3. Perform Operations 4. Control Flow 5. Produce Functions 6. Build Structures 7. Create Arrays 8. Harness Time 9. Manage Data 10. Handle Input 11. Employ Concurrency 12. Request Responses",
        "industryIdentifiers": [
          {
            "type": "ISBN_13",
            "identifier": "9781840789270"
          },
          {
            "type": "ISBN_10",
            "identifier": "1840789271"
          }
        ],
        "readingModes": {
          "text": true,
          "image": true
        },
        "pageCount": 337,
        "printType": "BOOK",
        "categories": [
          "Computers"
        ],
        "maturityRating": "NOT_MATURE",
        "allowAnonLogging": true,
        "contentVersion": "1.1.1.0.preview.3",
        "panelizationSummary": {
          "containsEpubBubbles": false,
          "containsImageBubbles": false
        },
        "imageLinks": {
          "smallThumbnail": "http://books.google.com/books/content?id=ba4IEAAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
          "thumbnail": "http://books.google.com/books/content?id=ba4IEAAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
        },
        "language": "en",
        "previewLink": "http://books.google.com/books?id=ba4IEAAAQBAJ&printsec=frontcover&dq=golang&hl=&cd=1&source=gbs_api",
        "infoLink": "http://books.google.com/books?id=ba4IEAAAQBAJ&dq=golang&hl=&source=gbs_api",
        "canonicalVolumeLink": "https://books.google.com/books/about/GO_Programming_in_easy_steps.html?hl=&id=ba4IEAAAQBAJ"
      },
      "saleInfo": {
        "country": "BG",
        "saleability": "NOT_FOR_SALE",
        "isEbook": false
      },
      "accessInfo": {
        "country": "BG",
        "viewability": "PARTIAL",
        "embeddable": true,
        "publicDomain": false,
        "textToSpeechPermission": "ALLOWED",
        "epub": {
          "isAvailable": true,
          "acsTokenLink": "http://books.google.com/books/download/GO_Programming_in_easy_steps-sample-epub.acsm?id=ba4IEAAAQBAJ&format=epub&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
        },
        "pdf": {
          "isAvailable": true,
          "acsTokenLink": "http://books.google.com/books/download/GO_Programming_in_easy_steps-sample-pdf.acsm?id=ba4IEAAAQBAJ&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
        },
        "webReaderLink": "http://play.google.com/books/reader?id=ba4IEAAAQBAJ&hl=&source=gbs_api",
        "accessViewStatus": "SAMPLE",
        "quoteSharingAllowed": false
      },
      "searchInfo": {
        "textSnippet": "GO Programming in easy steps begins by explaining how to easily create a programming environment on your own computer, so you can quickly begin to create your own working programs by copying the book&#39;s examples."
      }
    },
    {
      "kind": "books#volume",
      "id": "jvGaswEACAAJ",
      "etag": "MMCghvEnMak",
      "selfLink": "https://www.googleapis.com/books/v1/volumes/jvGaswEACAAJ",
      "volumeInfo": {
        "title": "Mastering Go",
        "subtitle": "Create Golang Production Applications Using Network Libraries, Concurrency, and Advanced Go Data Structures",
        "authors": [
          "Mihalis Tsoukalos"
        ],
        "publisher": "Packt Publishing",
        "publishedDate": "2018-04-30",
        "description": "Exploring the major features and packages of Go, along with its types and data-structures, enabling the reader to write threadsafe, concurrent cloud, and network applications Key Features Not your typical introduction to the Golang programming language Exploring Golang cradle to grave, completes the developer's Golang education A thorough exploration into the core libraries and Golang features, that usually are taken for granted In depth explanation, detailing the rationale behind composite data types, Golang concurrency, and the Golang networking library Book Description The Go programming language, often referred to as Golang (albeit wrongly), is really making strides, with some masterclass developments, architected by the greatest programming minds. Tobias Lutke, CEO of Shopify, recently quoted as saying \"Go will be the server language of the future\", powerful words, with much ambition. Go programmers are in high demand, but more controversially, Go takes the stage, where C and Unix programmers previously led the way. The growth of the Go language has seen it become the means by which systems, networking, web, and cloud applications are implemented. Comfortable with syntax, you'll benefit by mastering the use of the libraries and utilise its features, speed, and efficiency, for which the Go ecology is justly famous. You already know a little Go syntax and you've written some small projects, most Go programmers face the difficulty of having to integrate their Golang skills with production code. Typical introductions to Go programming, often stop short of this transition, the author continue on, showing you just how to tackle this. Offering a compendium of Go, the book begins with an account of how Go has been implemented, also, the reader will benefit from a dedicated chapter, an in-depth account of concurrency, systems and network programming, imperative for modern-day native cloud development. What you will learn Understand the design choices of Golang syntax Know enough Go internals to be able to optimize Golang code Appreciate concurrency models available in Golang Understand the interplay of systems and networking code Write server-level code that plays well in all environments Understand the context and appropriate use of Go data types and data structures Who this book is for This book is for Golang programmers. You should have previously read an introductory book on Go, or to have worked through the Tour of Go or an equivalent online course. This book will definitely help to remember the basic concepts of concurrency, but network programming will be explained. A certain amount of previous coding and production experience would be helpful.",
        "industryIdentifiers": [
          {
            "type": "ISBN_10",
            "identifier": "1788626540"
          },
          {
            "type": "ISBN_13",
            "identifier": "9781788626545"
          }
        ],
        "readingModes": {
          "text": false,
          "image": false
        },
        "pageCount": 606,
        "printType": "BOOK",
        "categories": [
          "Computers"
        ],
        "maturityRating": "NOT_MATURE",
        "allowAnonLogging": false,
        "contentVersion": "preview-1.0.0",
        "panelizationSummary": {
          "containsEpubBubbles": false,
          "containsImageBubbles": false
        },
        "imageLinks": {
          "smallThumbnail": "http://books.google.com/books/content?id=jvGaswEACAAJ&printsec=frontcover&img=1&zoom=5&source=gbs_api",
          "thumbnail": "http://books.google.com/books/content?id=jvGaswEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api"
        },
        "language": "en",
        "previewLink": "http://books.google.com/books?id=jvGaswEACAAJ&dq=golang&hl=&cd=2&source=gbs_api",
        "infoLink": "http://books.google.com/books?id=jvGaswEACAAJ&dq=golang&hl=&source=gbs_api",
        "canonicalVolumeLink": "https://books.google.com/books/about/Mastering_Go.html?hl=&id=jvGaswEACAAJ"
      },
      "saleInfo": {
        "country": "BG",
        "saleability": "NOT_FOR_SALE",
        "isEbook": false
      },
      "accessInfo": {
        "country": "BG",
        "viewability": "NO_PAGES",
        "embeddable": false,
        "publicDomain": false,
        "textToSpeechPermission": "ALLOWED",
        "epub": {
          "isAvailable": false
        },
        "pdf": {
          "isAvailable": false
        },
        "webReaderLink": "http://play.google.com/books/reader?id=jvGaswEACAAJ&hl=&source=gbs_api",
        "accessViewStatus": "NONE",
        "quoteSharingAllowed": false
      },
      "searchInfo": {
        "textSnippet": "This book will definitely help to remember the basic concepts of concurrency, but network programming will be explained. A certain amount of previous coding and production experience would be helpful."
      }
    }]`

var GoBooks []books.Item

func main() {
	err := json.Unmarshal([]byte(jsonText), &GoBooks)
	if err != nil {
		log.Fatalf("JSON unmarshalling failed: %s", err)
	}
	for i, book := range GoBooks {
		fmt.Printf("%d: %#v: %s\n", i+1, book.VolumeInfo.Title, book.VolumeInfo.Subtitle)
	}

	// Marshalling data to JSON
	data, err := json.MarshalIndent(GoBooks, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s", err)
	}
	fmt.Printf("AFTER MARSHALLING to JSON:\n%s\n", data)

}
