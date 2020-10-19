package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", testWriter)

	log.Println("Listening on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func testWriter(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 2; i++ {
		w.Write([]byte("aaaaaaaaaaaaaaaaaa"))
		time.Sleep(1 * time.Second)
	}
}
