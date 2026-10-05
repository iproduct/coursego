package main

import (
	"embed"
	"fmt"
	"net/http"
)

// 1. Embed a single file as a string
//
//go:embed version.txt
var version string

// 2. Embed an entire folder as an isolated file system
//
//go:embed static/*
var staticFiles embed.FS

func main() {
	// Print the string read automatically from version.txt
	fmt.Printf("Starting application version: %s\n", version)

	// Serve the embedded static directory over HTTP
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	fmt.Println("Server running on :8001...")
	http.ListenAndServe(":8001", nil)
}
