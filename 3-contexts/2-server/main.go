package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request initialized")
	defer log.Println("Request finalized")
	select {
	case <-time.After(5 * time.Second):
		// This will be printed in the terminal
		log.Println("Request handled with success")
		// This will returned for the request
		w.Write([]byte("Request handled with success"))
	case <-ctx.Done():
		// This will be printed in the terminal
		log.Println("Request canceled by client")
	}
}
