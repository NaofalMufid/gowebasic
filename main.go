package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// handler dengan variable
	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About Page"))
	}

	// routing
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mas", handler.MasHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	// handler routing langsung
	mux.HandleFunc("/tos", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Tos Page"))
	})

	log.Println("Starting web on port 8080")

	// run server
	err := http.ListenAndServe(":8080", mux)
	// log error
	log.Fatal(err)
}
