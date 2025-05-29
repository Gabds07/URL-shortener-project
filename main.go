package main

import (
	"fmt"
	"net/http"
	"url-shortener/utils"
)

func main() {
	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(wr, "Welcome to URL Shortener!")
	})

	http.HandleFunc("/shorten", func(wr http.ResponseWriter, req *http.Request) {
		url := req.FormValue("url")

		fmt.Println("Payload", url)

		shortURL := utils.GetShortCode()
		fullShortURL := fmt.Sprintf("http://localhost:3000/r/%s", shortURL)

		fmt.Printf("Generated short URL: %s\n", fullShortURL)
	})

	fmt.Println("Server running on: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}