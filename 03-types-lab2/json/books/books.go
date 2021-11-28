package books

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
}
