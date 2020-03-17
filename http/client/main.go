package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	params := map[string]string{
		"username": "calios",
		"password": "123456",
	}
	bs, _ := json.Marshal(params)
	resp, err := http.Post("http://localhost:8080/hello", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("resp: %+v", string(body))
}
