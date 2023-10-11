package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// file server
func FileServerHandler() http.Handler {
	return http.FileServer(http.Dir("./static"))
}

// main
func main() {
	port := getEnv("PORT", "3000")
	fs := FileServerHandler()

	log.Printf("Server is running at %s\n", port)

	http.Handle("/", fs)
	http.HandleFunc("/health", health)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handlers
func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
	log.Printf(r.Method + " " + r.Host + r.URL.Path)
}

// utils
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
