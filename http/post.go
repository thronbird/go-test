package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var client = &http.Client{}

func main() {
	token := getToken("http://jumpnew.tsign.cn/api", "devops", "Devops#123")
	print(token)
	token2 := getToken("http://192.168.22.166:8080/api", "admin", "Wq1a/R08qOubUE4K")
	print(token2)
}

func getToken(tokenUrl, username, password string) map[string]interface{} {
	URL := fmt.Sprintf("%s/v1/authentication/auth/", tokenUrl)
	payload := url.Values{}
	payload.Set("username", username)
	payload.Set("password", password)
	req, err := http.NewRequest("POST", URL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Close = true
	res, err := client.Do(req)
	if err != nil {
		print(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var result = make(map[string]interface{})
	json.Unmarshal(body, &result)
	return result
}
