package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://poseidon-master.esign.cn/api/app/args?offset=10&limit=2", nil)
	//req, err := http.NewRequest("GET", "http://localhost:9000/api/app/args?offset=10&limit=2", nil)
	if err != nil {
		return
	}

	client := &http.Client{}

	//for k, v := range r.Header {
	//	for _, vv := range v {
	//		req.Header.Add(k, vv)
	//	}
	//}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	//for k, v := range resp.Header {
	//	for _, vv := range v {
	//		c.Output.Header(k, vv)
	//	}
	//}
	data, _ := ioutil.ReadAll(resp.Body)
	//_ = c.Output.Body(data)
	print(string(data))
}
