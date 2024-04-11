package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/post", handler)

	fmt.Println("starting mirror server...")
	err := http.ListenAndServe("localhost:8001", nil)

	if err != nil {
		fmt.Println("error launching server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			_, _ = w.Write([]byte("bad request body"))
			return
		}

		fmt.Println(string(bytes))
		_, _ = w.Write(bytes)
		return
	}

	_, _ = w.Write([]byte("only POST method allowed"))
}
