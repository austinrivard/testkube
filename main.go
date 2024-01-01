package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "This is the root of a PR!\n")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		reqData, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Hello %s\n", reqData)
	})

	fmt.Println("Beginning to listen...")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	http.ListenAndServe(":8081", nil)
}
