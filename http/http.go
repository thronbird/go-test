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
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "hijacking not supported", 500)
		return
	}

	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer conn.Close()

	for i := 1; i <= 2; i++ {
		_, err = bufrw.WriteString("HTTP/1.1 200 OK\n\n")
		_, err = bufrw.WriteString("this")
		_, err = bufrw.WriteString("is")
		_, err = bufrw.WriteString("the")
		_, err = bufrw.WriteString("response")
		_, err = bufrw.WriteString("body")
		time.Sleep(1 * time.Second)
	}
}
