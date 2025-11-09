package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// set app port from env port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a file server to serve files from the "public" directory
	fs := http.FileServer(http.Dir("files/public"))

	// Handle requests to /public/ by stripping the prefix and serving files
	http.Handle("/files/public/", http.StripPrefix("/public/", fs))

	// Optionally, serve an index.html file at the root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "files/public/index.html")
			return
		}
		// If not the root, let the static file server handle it (if applicable)
		fs.ServeHTTP(w, r)
	})

	log.Printf("Server starting on %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
