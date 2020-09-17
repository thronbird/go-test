package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// e.g. /rule/period/groupid/metric
	http.HandleFunc("/callback/", func(w http.ResponseWriter, req *http.Request) {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
		b, err := ioutil.ReadAll(req.Body)
		fmt.Print(string(b))
	})

	s := &http.Server{
		Addr:           "0.0.0.0:7777",
		MaxHeaderBytes: 1 << 30,
	}
	log.Fatalln(s.ListenAndServe())
}
