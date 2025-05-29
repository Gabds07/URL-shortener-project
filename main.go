package main

import (
	"context"
	"fmt"
	"net/http"
	"url-shortener/utils"
)

var ctx = context.Background()

func main() {

	dbClient := utils.NewRedisClient()

	if dbClient == nil {
		fmt.Println("Failed to connect to Redis")
		return
	}

	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(wr, "Welcome to URL Shortener!")
	})

	http.HandleFunc("/shorten", func(wr http.ResponseWriter, req *http.Request) {
		url := req.FormValue("url")

		fmt.Println("Payload", url)

		shortURL := utils.GetShortCode()

		fullShortURL := fmt.Sprintf("http://localhost:3000/r/%s", shortURL)

		fmt.Printf("Generated short URL: %s\n", fullShortURL)

		utils.SetKey(&ctx, dbClient, shortURL, url, 0)
	})

	http.HandleFunc("/r/{code}", func(wr http.ResponseWriter, req *http.Request) {
		key := req.PathValue("code")

		if key == "" {
			http.Error(wr, "invalid URL", http.StatusBadRequest)
			return
		}

		longURL, err := utils.GetLongURL(&ctx, dbClient, key)

		if err != nil {
			http.Error(wr, "Shortener URL not found", http.StatusNotFound)
			return
		}

		http.Redirect(wr, req, longURL, http.StatusPermanentRedirect)
	})

	fmt.Println("Server running on: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}


