package main

import (
	"log"
	"net/url"
)

func parseURL(rawurl string) (url url.URL) {
	urlPtr, err := url.Parse(rawurl)
	if err != nil {
		log.Printf("Error parsing URL: '%s': %v\n", rawurl, err)
	}
	return *urlPtr
}

// Author models a book author
type Author string

// Category models a book category
type Category string

// Book struct models a book entity
type Book struct {
	ID                  string
	SelfLink            url.URL
	Title               string
	Subtitle            string
	Authors             []Author
	Publisher           string
	PublishedDate       string
	Description         string
	ISBN13              string
	ISBN10              string
	PageCount           int
	PrintType           string
	Categories          []Category
	SmallThumbnail      url.URL
	Thumbnail           url.URL
	Language            string
	PreviewLink         url.URL
	InfoLink            url.URL
	CanonicalVolumeLink url.URL
	RetailPrice         float64
	CurrencyCode        string
	BuyLink             url.URL
	TextSnippet         string
}

var GoBooks = []Book{
	{
		ID:                  "fmd-DwAAQBAJ",
		SelfLink:            parseURL("https://www.googleapis.com/books/v1/volumes/fmd-DwAAQBAJ"),
		Title:               "Hands-On Software Architecture with Golang",
		Subtitle:            "Design and architect highly scalable and robust applications using Go",
		Authors:             []Author{"Jyotiswarup Raiturkar"},
		Publisher:           "Packt Publishing Ltd",
		PublishedDate:       "2018-12-07",
		Description:         "Understand the principles of software architecture with coverage on SOA, distributed and messaging systems, and database modeling Key Features Gain knowledge of architectural approaches on SOA and microservices for architectural decisions Explore different architectural patterns for building distributed applications Migrate applications written in Java or Python to the Go language Book Description Building software requires careful planning and architectural considerations; Golang was developed with a fresh perspective on building next-generation applications on the cloud with distributed and concurrent computing concerns. Hands-On Software Architecture with Golang starts with a brief introduction to architectural elements, Go, and a case study to demonstrate architectural principles. You'll then move on to look at code-level aspects such as modularity, class design, and constructs specific to Golang and implementation of design patterns. As you make your way through the chapters, you'll explore the core objectives of architecture such as effectively managing complexity, scalability, and reliability of software systems. You'll also work through creating distributed systems and their communication before moving on to modeling and scaling of data. In the concluding chapters, you'll learn to deploy architectures and plan the migration of applications from other languages. By the end of this book, you will have gained insight into various design and architectural patterns, which will enable you to create robust, scalable architecture using Golang. What you will learn Understand architectural paradigms and deep dive into Microservices Design parallelism/concurrency patterns and learn object-oriented design patterns in Go Explore API-driven systems architecture with introduction to REST and GraphQL standards Build event-driven architectures and make your architectures anti-fragile Engineer scalability and learn how to migrate to Go from other languages Get to grips with deployment considerations with CICD pipeline, cloud deployments, and so on Build an end-to-end e-commerce (travel) application backend in Go Who this book is for Hands-On Software Architecture with Golang is for software developers, architects, and CTOs looking to use Go in their software architecture to build enterprise-grade applications. Programming knowledge of Golang is assumed.",
		ISBN13:              "9781788625104",
		ISBN10:              "1788625102",
		PageCount:           500,
		PrintType:           "BOOK",
		Categories:          []Category{"Computers"},
		SmallThumbnail:      parseURL("http://books.google.com/books/content?id=fmd-DwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api"),
		Thumbnail:           parseURL("http://books.google.com/books/content?id=fmd-DwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"),
		Language:            "en",
		PreviewLink:         parseURL("http://books.google.com/books?id=fmd-DwAAQBAJ&pg=PT545&dq=golang&hl=&cd=1&source=gbs_api"),
		InfoLink:            parseURL("https://play.google.com/store/books/details?id=fmd-DwAAQBAJ&source=gbs_api"),
		CanonicalVolumeLink: parseURL("https://play.google.com/store/books/details?id=fmd-DwAAQBAJ"),
		RetailPrice:         84.17,
		CurrencyCode:        "BGN",
		BuyLink:             parseURL("https://play.google.com/store/books/details?id=fmd-DwAAQBAJ&rdid=book-fmd-DwAAQBAJ&rdot=1&source=gbs_api"),
		TextSnippet:         "The scope of unit testing is to test individual modules (classes/functions) of the <br> service. These are generally supported by a variety of frameworks (most of which <br> are language-specific). <b>Golang</b> has a very powerful, in-built testing framework,&nbsp;...",
	},
	{
		ID:                  "o86PDwAAQBAJ",
		SelfLink:            parseURL("https://www.googleapis.com/books/v1/volumes/o86PDwAAQBAJ"),
		Title:               "Learn Data Structures and Algorithms with Golang",
		Subtitle:            "Level up your Go programming skills to develop faster and more efficient code",
		Authors:             []Author{"Bhagvan Kommadi"},
		Publisher:           "Packt Publishing Ltd",
		PublishedDate:       "2019-03-30",
		Description:         "Explore Golang's data structures and algorithms to design, implement, and analyze code in the professional setting Key Features Learn the basics of data structures and algorithms and implement them efficiently Use data structures such as arrays, stacks, trees, lists and graphs in real-world scenarios Compare the complexity of different algorithms and data structures for improved code performance Book Description Golang is one of the fastest growing programming languages in the software industry. Its speed, simplicity, and reliability make it the perfect choice for building robust applications. This brings the need to have a solid foundation in data structures and algorithms with Go so as to build scalable applications. Complete with hands-on tutorials, this book will guide you in using the best data structures and algorithms for problem solving. The book begins with an introduction to Go data structures and algorithms. You'll learn how to store data using linked lists, arrays, stacks, and queues. Moving ahead, you'll discover how to implement sorting and searching algorithms, followed by binary search trees. This book will also help you improve the performance of your applications by stringing data types and implementing hash structures in algorithm design. Finally, you'll be able to apply traditional data structures to solve real-world problems. By the end of the book, you'll have become adept at implementing classic data structures and algorithms in Go, propelling you to become a confident Go programmer. What you will learn Improve application performance using the most suitable data structure and algorithm Explore the wide range of classic algorithms such as recursion and hashing algorithms Work with algorithms such as garbage collection for efficient memory management Analyze the cost and benefit trade-off to identify algorithms and data structures for problem solving Explore techniques for writing pseudocode algorithm and ace whiteboard coding in interviews Discover the pitfalls in selecting data structures and algorithms by predicting their speed and efficiency Who this book is for This book is for developers who want to understand how to select the best data structures and algorithms that will help solve coding problems. Basic Go programming experience will be an added advantage.",
		ISBN13:              "9781789618419",
		ISBN10:              "178961841X",
		PageCount:           336,
		PrintType:           "BOOK",
		Categories:          []Category{"Computers"},
		SmallThumbnail:      parseURL("http://books.google.com/books/content?id=o86PDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api"),
		Thumbnail:           parseURL("http://books.google.com/books/content?id=o86PDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"),
		Language:            "en",
		PreviewLink:         parseURL("http://books.google.com/books?id=o86PDwAAQBAJ&pg=PA288&dq=golang&hl=&cd=2&source=gbs_api"),
		InfoLink:            parseURL("https://play.google.com/store/books/details?id=o86PDwAAQBAJ&source=gbs_api"),
		CanonicalVolumeLink: parseURL("https://play.google.com/store/books/details?id=o86PDwAAQBAJ"),
		RetailPrice:         67.33,
		CurrencyCode:        "BGN",
		BuyLink:             parseURL("https://play.google.com/store/books/details?id=o86PDwAAQBAJ&rdid=book-o86PDwAAQBAJ&rdot=1&source=gbs_api"),
		TextSnippet:         "... cursive-in-<b>golang</b>-c196ca5fd489) The following papers are related to classic <br> algorithms: A Mathematical Modeling of Pure, Recursive Algorithms (https://www.<br> researchgate.net/publication/220810107_A_Mathematical_Mod&nbsp;...",
	},
	{
		ID:                  "xfPEDwAAQBAJ",
		SelfLink:            parseURL("https://www.googleapis.com/books/v1/volumes/xfPEDwAAQBAJ"),
		Title:               "From Ruby to Golang",
		Subtitle:            "A Ruby Programmer’s Guide to Learning Go",
		Authors:             []Author{"Joel Bryan Juliano"},
		Publisher:           "Joel Bryan Juliano",
		PublishedDate:       "2019-07-15",
		Description:         "Imagine that you like to learn a new programming language, and you start by leveraging what you already know and bridge the gap in learning specific parts of the new language. This book was created on that idea, it starts with using my existing language knowledge and experience to breakdown Go into familiar Ruby concepts and implementations. The first thing I did to learn Go professionally is to relate to what I know in Ruby. I’ve been a professional Ruby programmer since 2009 and in over a decade of professional experience working as a software engineer, I have worked on multiple programming languages. And proven personally that it’s easier to learn a programming concept from something familiar to me. This helps me to learn the new language faster, which also means being productive much faster as well. This book was created on my first-hand experience of learning Go from my existing knowledge and experience in Ruby. The book was carefully thought from ground-up, collecting familiar patterns, abstracts, and analogs in Ruby, and relate it with a proper implementation in Go. By teaching familiar implementations found in Ruby, you will see the correlation between the two languages, establishing familiar concepts to give you enough knowledge to be comfortable with Go and to start programming with it. Go is an easy language to work with, it’s modern, flexible, powerful and fast. It compiles to binary which gives it an ability for a binary distribution that runs on different platforms, and Go has almost in par performance with C, with package support, memory safety, automatic garbage collection and concurrency built-in. And you get all the nice features from a statically typed language, which IDEs can make use of, and so also improving your development workflow. Notable open-source projects are built using Go (i.e. Docker, Kubernetes, Etherium and Terraform to name a few), this gives you an advantage because those platforms have APIs and SDKs readily available in Go natively for you to use. And many global companies have been using Go in production (i.e. Google, Netflix, Dropbox, Heroku and Uber to name a few), proving that it has been battle-tested and powerful mature language to based your work into. Go is created by an interesting mixed of people. Google is the company that funded Go’s development, and the authors of Go who designed the language are mainly Robert Griese- mer (worked on V8 Javascript Engine, Java HotSpot VM, and the Strongtalk system), Rob Pike (known for Plan 9 and UTF-8), and Ken Thompson (known for Unix, C programming language, Plan 9, UTF-8 and Inferno to name a few). This book will definitely help you get started with Go from your existing Ruby knowledge, and start to hit the ground fast, running!",
		ISBN13:              "9781080944002",
		ISBN10:              "1080944001",
		PageCount:           135,
		PrintType:           "BOOK",
		Categories:          []Category{"Computers"},
		SmallThumbnail:      parseURL("http://books.google.com/books/content?id=xfPEDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api"),
		Thumbnail:           parseURL("http://books.google.com/books/content?id=xfPEDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"),
		Language:            "en",
		PreviewLink:         parseURL("http://books.google.com/books?id=xfPEDwAAQBAJ&pg=PA73&dq=golang&hl=&cd=3&source=gbs_api"),
		InfoLink:            parseURL("https://play.google.com/store/books/details?id=xfPEDwAAQBAJ&source=gbs_api"),
		CanonicalVolumeLink: parseURL("https://play.google.com/store/books/details?id=xfPEDwAAQBAJ"),
		RetailPrice:         46.94,
		CurrencyCode:        "BGN",
		BuyLink:             parseURL("https://play.google.com/store/books/details?id=xfPEDwAAQBAJ&rdid=book-xfPEDwAAQBAJ&rdot=1&source=gbs_api"),
		TextSnippet:         "1 a = [1, 2, 3, 4, 5, 0] 2 a.drop(3) #=&gt; [4, 5, 0] And here&#39;s our <b>Golang</b> <br> implementation drop. 1 a := [...]int{1,2,3,4,5,0} 2 drop := a[3:] 3 4 fmt.Println(drop) <br> 8.7 drop_while Similar to drop, drop_while returns the values that satisfy an <br> expression.",
	},
	// {
	// kind: "books#volume",
	// id: "NvNFDwAAQBAJ",
	// etag: "AXgOlt0cnI4",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/NvNFDwAAQBAJ",
	// volumeInfo: {
	// title: "Cloud Native Programming with Golang",
	// subtitle: "Develop microservice-based high performance web apps for the cloud with Go",
	// authors: [
	// "Mina Andrawos",
	// "Martin Helmich"
	// ],
	// publisher: "Packt Publishing Ltd",
	// publishedDate: "2017-12-28",
	// description: "Discover practical techniques to build cloud-native apps that are scalable, reliable, and always available. Key Features Build well-designed and secure microservices. Enrich your microservices with continous integration and monitoring. Containerize your application with Docker Deploy your application to AWS. Learn how to utilize the powerful AWS services from within your application Book Description Awarded as one of the best books of all time by BookAuthority, Cloud Native Programming with Golang will take you on a journey into the world of microservices and cloud computing with the help of Go. Cloud computing and microservices are two very important concepts in modern software architecture. They represent key skills that ambitious software engineers need to acquire in order to design and build software applications capable of performing and scaling. Go is a modern cross-platform programming language that is very powerful yet simple; it is an excellent choice for microservices and cloud applications. Go is gaining more and more popularity, and becoming a very attractive skill. This book starts by covering the software architectural patterns of cloud applications, as well as practical concepts regarding how to scale, distribute, and deploy those applications. You will also learn how to build a JavaScript-based front-end for your application, using TypeScript and React. From there, we dive into commercial cloud offerings by covering AWS. Finally, we conclude our book by providing some overviews of other concepts and technologies that you can explore, to move from where the book leaves off. What you will learn Understand modern software applications architectures Build secure microservices that can effectively communicate with other services Get to know about event-driven architectures by diving into message queues such as Kafka, Rabbitmq, and AWS SQS. Understand key modern database technologies such as MongoDB, and Amazon’s DynamoDB Leverage the power of containers Explore Amazon cloud services fundamentals Know how to utilize the power of the Go language to access key services in the Amazon cloud such as S3, SQS, DynamoDB and more. Build front-end applications using ReactJS with Go Implement CD for modern applications Who this book is for This book is for developers who want to begin building secure, resilient, robust, and scalable Go applications that are cloud native. Some knowledge of the Go programming language should be sufficient.To build the front-end application, you will also need some knowledge of JavaScript programming.",
	// industryIDentifiers: [
	// {
	// type: "ISBN_13",
	// identifier: "9781787127968"
	// },
	// {
	// type: "ISBN_10",
	// identifier: "1787127966"
	// }
	// ],
	// readingModes: {
	// text: true,
	// image: true
	// },
	// pageCount: 404,
	// printType: "BOOK",
	// categories: [
	// "Computers"
	// ],
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: true,
	// contentVersion: "1.4.4.0.preview.3",
	// panelizationSummary: {
	// containsEpubBubbles: false,
	// containsImageBubbles: false
	// },
	// imageLinks: {
	// smallThumbnail: "http://books.google.com/books/content?id=NvNFDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
	// thumbnail: "http://books.google.com/books/content?id=NvNFDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
	// },
	// language: "en",
	// previewLink: "http://books.google.com/books?id=NvNFDwAAQBAJ&pg=PP3&dq=golang&hl=&cd=4&source=gbs_api",
	// infoLink: "https://play.google.com/store/books/details?id=NvNFDwAAQBAJ&source=gbs_api",
	// canonicalVolumeLink: "https://play.google.com/store/books/details?id=NvNFDwAAQBAJ"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "FOR_SALE",
	// isEbook: true,
	// listPrice: {
	// amount: 75.75,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amount: 75.75,
	// currencyCode: "BGN"
	// },
	// buyLink: "https://play.google.com/store/books/details?id=NvNFDwAAQBAJ&rdid=book-NvNFDwAAQBAJ&rdot=1&source=gbs_api",
	// offers: [
	// {
	// finskyOfferType: 1,
	// listPrice: {
	// amountInMicros: 75750000,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amountInMicros: 75750000,
	// currencyCode: "BGN"
	// }
	// }
	// ]
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "PARTIAL",
	// embeddable: true,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: true
	// },
	// pdf: {
	// isAvailable: true
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=NvNFDwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "SAMPLE",
	// quoteSharingAllowed: false
	// },
	// searchInfo: {
	// textSnippet: "<b>Golang</b>. Copyright © 2017 Packt Publishing All rights reserved. No part of this <br> book may be reproduced, stored in a retrieval system, or transmitted in any form <br> or by any means, without the prior written permission of the publisher, except in <br> the&nbsp;..."
	// }
	// },
	// {
	// kind: "books#volume",
	// id: "TxJ9DwAAQBAJ",
	// etag: "M/DlWUHIiD8",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/TxJ9DwAAQBAJ",
	// volumeInfo: {
	// title: "Go Lang Cryptography",
	// authors: [
	// "Anish Nath"
	// ],
	// publisher: "Anish Nath",
	// description: "Cryptography is for everyone, no matter which role, function you are in, a basic level of security is needed. The style and approach is used in this book is to full-fill all of the cryptography needs for the go lang programmer from beginner to advanced level. What you will learn Encoding/Decoding,Random Number, Hashing, blake, HKDF, PBKDF, Argon2, Scrypt, Bcrypt, RSA, DSA, ECDSA, Curve25519, Nacl, AES, chacha20poly1305 ,RC4, BlowFish, TwoFish, 3DES, HMAC, OpenPGP, SSH-Client, HTTPS, X.509 Certificate Handing, Encrypted PEM files, OCSP",
	// industryIDentifiers: [
	// {
	// type: "ISBN_13",
	// identifier: "9781790681075"
	// },
	// {
	// type: "ISBN_10",
	// identifier: "1790681073"
	// }
	// ],
	// readingModes: {
	// text: false,
	// image: true
	// },
	// pageCount: 200,
	// printType: "BOOK",
	// averageRating: 5,
	// ratingsCount: 1,
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: true,
	// contentVersion: "preview-1.0.0",
	// panelizationSummary: {
	// containsEpubBubbles: false,
	// containsImageBubbles: false
	// },
	// imageLinks: {
	// smallThumbnail: "http://books.google.com/books/content?id=TxJ9DwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
	// thumbnail: "http://books.google.com/books/content?id=TxJ9DwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
	// },
	// language: "en",
	// previewLink: "http://books.google.com/books?id=TxJ9DwAAQBAJ&pg=PA11&dq=golang&hl=&cd=5&source=gbs_api",
	// infoLink: "https://play.google.com/store/books/details?id=TxJ9DwAAQBAJ&source=gbs_api",
	// canonicalVolumeLink: "https://play.google.com/store/books/details?id=TxJ9DwAAQBAJ"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "FOR_SALE",
	// isEbook: true,
	// listPrice: {
	// amount: 12.56,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amount: 12.56,
	// currencyCode: "BGN"
	// },
	// buyLink: "https://play.google.com/store/books/details?id=TxJ9DwAAQBAJ&rdid=book-TxJ9DwAAQBAJ&rdot=1&source=gbs_api",
	// offers: [
	// {
	// finskyOfferType: 1,
	// listPrice: {
	// amountInMicros: 12560000,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amountInMicros: 12560000,
	// currencyCode: "BGN"
	// }
	// }
	// ]
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "PARTIAL",
	// embeddable: true,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: false
	// },
	// pdf: {
	// isAvailable: true,
	// acsTokenLink: "http://books.google.com/books/download/Go_Lang_Cryptography-sample-pdf.acsm?id=TxJ9DwAAQBAJ&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=TxJ9DwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "SAMPLE",
	// quoteSharingAllowed: false
	// },
	// searchInfo: {
	// textSnippet: "... <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. ... org/x/crypto/mdé <b>golang</b> <br> <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. ... org/x/crypto/openpgp/elgamal <br> <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. <b>golang</b>. ... org/x/crypto/ripemd160 <br> <b>golang</b>&nbsp;..."
	// }
	// },
	// {
	// kind: "books#volume",
	// id: "gvBZDwAAQBAJ",
	// etag: "sOAog/fvVvU",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/gvBZDwAAQBAJ",
	// volumeInfo: {
	// title: "Mastering Go",
	// subtitle: "Create Golang production applications using network libraries, concurrency, and advanced Go data structures",
	// authors: [
	// "Mihalis Tsoukalos"
	// ],
	// publisher: "Packt Publishing Ltd",
	// publishedDate: "2018-04-30",
	// description: "Exploring the major features and packages of Go, along with its types and data-structures, enabling the reader to write threadsafe, concurrent cloud, and network applications Key Features Not your typical introduction to the Golang programming language Exploring Golang cradle to grave, completes the developer’s Golang education A thorough exploration into the core libraries and Golang features, that usually are taken for granted In depth explanation, detailing the rationale behind composite data types, Golang concurrency, and the Golang networking library Book Description Often referred to as Golang (albeit wrongly), the Go programming language is really making strides thanks to some masterclass developments, architected by the greatest programming minds. Shopify CEO Tobias Lutke has been recently quoted as saying “Go will be the server language of the future.” Go programmers are in high demand, but - more controversially - Go takes the stage where C and Unix programmers previously led the way. The growth of the Go language has seen it become the means by which systems, networking, web, and cloud applications are implemented. If you’re a Go programmer, you’ll already know some Go syntax and will have written some small projects. However, most Go programmers face the difficulty of having to integrate their Golang skills with production code. With Mastering Go, the author shows you just how to tackle this problem. You'll benefit by mastering the use of the libraries and utilize its features, speed, and efficiency for which the Go ecology is justly famous. Offering a compendium of Go, the book begins with an account of how Go has been implemented. You'll also benefit from an in-depth account of concurrency and systems and network programming imperative for modern-day native cloud development through the course of the book. What you will learn Understand the design choices of Golang syntax Know enough Go internals to be able to optimize Golang code Appreciate concurrency models available in Golang Understand the interplay of systems and networking code Write server-level code that plays well in all environments Understand the context and appropriate use of Go data types and data structures Who this book is for This book is for Golang programmers. You should have previously read an introductory book on Go, or to have worked through the Tour of Go or an equivalent online course. This book will definitely help to remember the basic concepts of concurrency, but network programming will be explained. A certain amount of previous coding and production experience would be helpful.",
	// industryIDentifiers: [
	// {
	// type: "ISBN_13",
	// identifier: "9781788622530"
	// },
	// {
	// type: "ISBN_10",
	// identifier: "1788622537"
	// }
	// ],
	// readingModes: {
	// text: true,
	// image: true
	// },
	// pageCount: 606,
	// printType: "BOOK",
	// categories: [
	// "Computers"
	// ],
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: true,
	// contentVersion: "1.2.2.0.preview.3",
	// panelizationSummary: {
	// containsEpubBubbles: false,
	// containsImageBubbles: false
	// },
	// imageLinks: {
	// smallThumbnail: "http://books.google.com/books/content?id=gvBZDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
	// thumbnail: "http://books.google.com/books/content?id=gvBZDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
	// },
	// language: "en",
	// previewLink: "http://books.google.com/books?id=gvBZDwAAQBAJ&printsec=frontcover&dq=golang&hl=&cd=6&source=gbs_api",
	// infoLink: "https://play.google.com/store/books/details?id=gvBZDwAAQBAJ&source=gbs_api",
	// canonicalVolumeLink: "https://play.google.com/store/books/details?id=gvBZDwAAQBAJ"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "FOR_SALE",
	// isEbook: true,
	// listPrice: {
	// amount: 84.17,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amount: 84.17,
	// currencyCode: "BGN"
	// },
	// buyLink: "https://play.google.com/store/books/details?id=gvBZDwAAQBAJ&rdid=book-gvBZDwAAQBAJ&rdot=1&source=gbs_api",
	// offers: [
	// {
	// finskyOfferType: 1,
	// listPrice: {
	// amountInMicros: 84170000,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amountInMicros: 84170000,
	// currencyCode: "BGN"
	// }
	// }
	// ]
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "PARTIAL",
	// embeddable: true,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: true
	// },
	// pdf: {
	// isAvailable: true
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=gvBZDwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "SAMPLE",
	// quoteSharingAllowed: false
	// },
	// searchInfo: {
	// textSnippet: "This book will definitely help to remember the basic concepts of concurrency, but network programming will be explained. A certain amount of previous coding and production experience would be helpful."
	// }
	// },
	// {
	// kind: "books#volume",
	// id: "l6XBwQEACAAJ",
	// etag: "7hOhFfXJbjA",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/l6XBwQEACAAJ",
	// volumeInfo: {
	// title: "Go Programming for Network Operations: A Golang Network Automation Handbook",
	// authors: [
	// "Tom McAllen"
	// ],
	// publisher: "Independently Published",
	// publishedDate: "2019-02-08",
	// description: "This book illustrates how to apply Go programming to network operations. The topics cover common use cases through examples that are designed to act as a guide and serve as a reference. The reader is assumed to have already gained a fundamental understanding of Go programming; however, the examples are explained for additional clarification. The focus is on using Go for network operations, not on the language itself.",
	// industryIDentifiers: [
	// {
	// type: "ISBN_10",
	// identifier: "1793121230"
	// },
	// {
	// type: "ISBN_13",
	// identifier: "9781793121233"
	// }
	// ],
	// readingModes: {
	// text: false,
	// image: false
	// },
	// pageCount: 112,
	// printType: "BOOK",
	// categories: [
	// "Computers"
	// ],
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: false,
	// contentVersion: "preview-1.0.0",
	// panelizationSummary: {
	// containsEpubBubbles: false,
	// containsImageBubbles: false
	// },
	// imageLinks: {
	// smallThumbnail: "http://books.google.com/books/content?id=l6XBwQEACAAJ&printsec=frontcover&img=1&zoom=5&source=gbs_api",
	// thumbnail: "http://books.google.com/books/content?id=l6XBwQEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api"
	// },
	// language: "en",
	// previewLink: "http://books.google.com/books?id=l6XBwQEACAAJ&dq=golang&hl=&cd=7&source=gbs_api",
	// infoLink: "http://books.google.com/books?id=l6XBwQEACAAJ&dq=golang&hl=&source=gbs_api",
	// canonicalVolumeLink: "https://books.google.com/books/about/Go_Programming_for_Network_Operations_A.html?hl=&id=l6XBwQEACAAJ"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "NOT_FOR_SALE",
	// isEbook: false
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "NO_PAGES",
	// embeddable: false,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: false
	// },
	// pdf: {
	// isAvailable: false
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=l6XBwQEACAAJ&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "NONE",
	// quoteSharingAllowed: false
	// },
	// searchInfo: {
	// textSnippet: "This book illustrates how to apply Go programming to network operations."
	// }
	// },
	// {
	// kind: "books#volume",
	// id: "pSZKDwAAQBAJ",
	// etag: "qGgEv9Nqzto",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/pSZKDwAAQBAJ",
	// volumeInfo: {
	// title: "Security with Go",
	// subtitle: "Explore the power of Golang to secure host, web, and cloud services",
	// authors: [
	// "John Daniel Leon"
	// ],
	// publisher: "Packt Publishing Ltd",
	// publishedDate: "2018-01-31",
	// description: "The first stop for your security needs when using Go, covering host, network, and cloud security for ethical hackers and defense against intrusion Key Features First introduction to Security with Golang Adopting a Blue Team/Red Team approach Take advantage of speed and inherent safety of Golang Works as an introduction to security for Golang developers Works as a guide to Golang security packages for recent Golang beginners Book Description Go is becoming more and more popular as a language for security experts. Its wide use in server and cloud environments, its speed and ease of use, and its evident capabilities for data analysis, have made it a prime choice for developers who need to think about security. Security with Go is the first Golang security book, and it is useful for both blue team and red team applications. With this book, you will learn how to write secure software, monitor your systems, secure your data, attack systems, and extract information. Defensive topics include cryptography, forensics, packet capturing, and building secure web applications. Offensive topics include brute force, port scanning, packet injection, web scraping, social engineering, and post exploitation techniques. What you will learn Learn the basic concepts and principles of secure programming Write secure Golang programs and applications Understand classic patterns of attack Write Golang scripts to defend against network-level attacks Learn how to use Golang security packages Apply and explore cryptographic methods and packages Learn the art of defending against brute force attacks Secure web and cloud applications Who this book is for Security with Go is aimed at developers with basics in Go to the level that they can write their own scripts and small programs without difficulty. Readers should be familiar with security concepts, and familiarity with Python security applications and libraries is an advantage, but not a necessity.",
	// industryIDentifiers: [
	// {
	// type: "ISBN_13",
	// identifier: "9781788622257"
	// },
	// {
	// type: "ISBN_10",
	// identifier: "1788622251"
	// }
	// ],
	// readingModes: {
	// text: true,
	// image: true
	// },
	// pageCount: 340,
	// printType: "BOOK",
	// categories: [
	// "Computers"
	// ],
	// averageRating: 5,
	// ratingsCount: 1,
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: true,
	// contentVersion: "1.1.2.0.preview.3",
	// panelizationSummary: {
	// containsEpubBubbles: false,
	// containsImageBubbles: false
	// },
	// imageLinks: {
	// smallThumbnail: "http://books.google.com/books/content?id=pSZKDwAAQBAJ&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
	// thumbnail: "http://books.google.com/books/content?id=pSZKDwAAQBAJ&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
	// },
	// language: "en",
	// previewLink: "http://books.google.com/books?id=pSZKDwAAQBAJ&printsec=frontcover&dq=golang&hl=&cd=8&source=gbs_api",
	// infoLink: "https://play.google.com/store/books/details?id=pSZKDwAAQBAJ&source=gbs_api",
	// canonicalVolumeLink: "https://play.google.com/store/books/details?id=pSZKDwAAQBAJ"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "FOR_SALE",
	// isEbook: true,
	// listPrice: {
	// amount: 67.33,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amount: 67.33,
	// currencyCode: "BGN"
	// },
	// buyLink: "https://play.google.com/store/books/details?id=pSZKDwAAQBAJ&rdid=book-pSZKDwAAQBAJ&rdot=1&source=gbs_api",
	// offers: [
	// {
	// finskyOfferType: 1,
	// listPrice: {
	// amountInMicros: 67330000,
	// currencyCode: "BGN"
	// },
	// retailPrice: {
	// amountInMicros: 67330000,
	// currencyCode: "BGN"
	// }
	// }
	// ]
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "PARTIAL",
	// embeddable: true,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: true
	// },
	// pdf: {
	// isAvailable: true
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=pSZKDwAAQBAJ&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "SAMPLE",
	// quoteSharingAllowed: false
	// },
	// searchInfo: {
	// textSnippet: "What you will learn Learn the basic concepts and principles of secure programming Write secure Golang programs and applications Understand classic patterns of attack Write Golang scripts to defend against network-level attacks Learn how to ..."
	// }
	// },
	// {
	// kind: "books#volume",
	// id: "f47xsgEACAAJ",
	// etag: "bw/I/iO+wGk",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/f47xsgEACAAJ",
	// volumeInfo: {
	// title: "A Framework for Actors in GoLang",
	// publishedDate: "2014",
	// industryIDentifiers: [
	// {
	// type: "OTHER",
	// identifier: "OCLC:908403397"
	// }
	// ],
	// readingModes: {
	// text: false,
	// image: false
	// },
	// pageCount: 113,
	// printType: "BOOK",
	// categories: [
	// "Parallel programming (Computer science)"
	// ],
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: false,
	// contentVersion: "preview-1.0.0",
	// language: "en",
	// previewLink: "http://books.google.com/books?id=f47xsgEACAAJ&dq=golang&hl=&cd=9&source=gbs_api",
	// infoLink: "http://books.google.com/books?id=f47xsgEACAAJ&dq=golang&hl=&source=gbs_api",
	// canonicalVolumeLink: "https://books.google.com/books/about/A_Framework_for_Actors_in_GoLang.html?hl=&id=f47xsgEACAAJ"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "NOT_FOR_SALE",
	// isEbook: false
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "NO_PAGES",
	// embeddable: false,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: false
	// },
	// pdf: {
	// isAvailable: false
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=f47xsgEACAAJ&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "NONE",
	// quoteSharingAllowed: false
	// }
	// },
	// {
	// kind: "books#volume",
	// id: "cVLHYuAjB4IC",
	// etag: "MR72epK7T58",
	// selfLink: "https://www.googleapis.com/books/v1/volumes/cVLHYuAjB4IC",
	// volumeInfo: {
	// title: "The Wizard of the Stove Pipe Mountains",
	// subtitle: "First Book of the Nebulæ Stone Series",
	// authors: [
	// "Ewart R N Jowett"
	// ],
	// publisher: "Xlibris Corporation",
	// publishedDate: "2012-06-15",
	// description: "Riccard and Rowina are on a mission to find their missing father, Ivan. Eventually, they arrive in the town of Weirdsdale, near The Stove Pipe Mountains. One of these, Vulcan occasionally emits weird lights, during which men caught outside at night disappear. Suspecting that Ivan has been captured by The Lights they access Vulcan, finding it to be the home of Foxwald, who assures them that Ivan is not in Vulcan, but he allows them to stay and search. During Foxwald's absence, battling the forces of evil, they find their father, though he does not recognize them. Returning Foxwald is shocked to find that he himself is responsible for their missing father.",
	// industryIDentifiers: [
	// {
	// type: "ISBN_13",
	// identifier: "9781469165820"
	// },
	// {
	// type: "ISBN_10",
	// identifier: "1469165821"
	// }
	// ],
	// readingModes: {
	// text: true,
	// image: true
	// },
	// pageCount: 737,
	// printType: "BOOK",
	// categories: [
	// "Fiction"
	// ],
	// maturityRating: "NOT_MATURE",
	// allowAnonLogging: false,
	// contentVersion: "2.4.4.0.preview.3",
	// panelizationSummary: {
	// containsEpubBubbles: false,
	// containsImageBubbles: false
	// },
	// imageLinks: {
	// smallThumbnail: "http://books.google.com/books/content?id=cVLHYuAjB4IC&printsec=frontcover&img=1&zoom=5&edge=curl&source=gbs_api",
	// thumbnail: "http://books.google.com/books/content?id=cVLHYuAjB4IC&printsec=frontcover&img=1&zoom=1&edge=curl&source=gbs_api"
	// },
	// language: "en",
	// previewLink: "http://books.google.com/books?id=cVLHYuAjB4IC&pg=PA193&dq=golang&hl=&cd=10&source=gbs_api",
	// infoLink: "http://books.google.com/books?id=cVLHYuAjB4IC&dq=golang&hl=&source=gbs_api",
	// canonicalVolumeLink: "https://books.google.com/books/about/The_Wizard_of_the_Stove_Pipe_Mountains.html?hl=&id=cVLHYuAjB4IC"
	// },
	// saleInfo: {
	// country: "BG",
	// saleability: "NOT_FOR_SALE",
	// isEbook: false
	// },
	// accessInfo: {
	// country: "BG",
	// viewability: "PARTIAL",
	// embeddable: true,
	// publicDomain: false,
	// textToSpeechPermission: "ALLOWED",
	// epub: {
	// isAvailable: true,
	// acsTokenLink: "http://books.google.com/books/download/The_Wizard_of_the_Stove_Pipe_Mountains-sample-epub.acsm?id=cVLHYuAjB4IC&format=epub&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
	// },
	// pdf: {
	// isAvailable: true,
	// acsTokenLink: "http://books.google.com/books/download/The_Wizard_of_the_Stove_Pipe_Mountains-sample-pdf.acsm?id=cVLHYuAjB4IC&format=pdf&output=acs4_fulfillment_token&dl_type=sample&source=gbs_api"
	// },
	// webReaderLink: "http://play.google.com/books/reader?id=cVLHYuAjB4IC&hl=&printsec=frontcover&source=gbs_api",
	// accessViewStatus: "SAMPLE",
	// quoteSharingAllowed: false
	// },
	// searchInfo: {
	// textSnippet: "“<b>Golang</b> about here! Enemy to Waangann.” “You&#39;re going nowhere!” ordered the <br> leader of the three. “If there are savages around here, you can stay and give us a <br> hand if they come at us.” The Waangan tracker rolled his eyes. “I not stay—hear&nbsp;..."
	// }
	// }
	// ]
	// }
}
