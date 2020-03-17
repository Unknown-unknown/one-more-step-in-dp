package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("handler err: %+v", err)
		}
		defer req.Body.Close()
		io.WriteString(w, "Hello, world!"+string(body))
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
