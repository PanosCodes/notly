package main

import (
	"log"
	"net/http"

	"github.com/gen2brain/beeep"
)

func notify(rw http.ResponseWriter, r *http.Request) {
	message := r.URL.Query().Get("message")

	err := beeep.Alert("Notly", message, "")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", notify)

	log.Fatal(http.ListenAndServe(":1488", nil))
}
