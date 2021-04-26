package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := []byte(`{"aaa" :{
    "bbb":["ccc", "doc"],
    "安全": []
    }}`)
	//y := make(map[string]interface{})
	//y := new(map[string]interface{})
	y := &struct {
		A map[string]interface{} `json:"aaa"`
	}{A: make(map[string]interface{})}

	//print(x)
	if err := json.Unmarshal(x, y); err != nil {
		print(err)
	}
	fmt.Printf("%#v", *y)
	//dumpMap("", jsonMap)
}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}
