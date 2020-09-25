package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Test !!!"))
}
func main() {
	http.HandleFunc("/hello", hello)

	log.Println("Listening on :9999...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("Error")
	}
}
