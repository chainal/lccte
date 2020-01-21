package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "This is my http server")
		if err != nil {
			fmt.Println("write response error")
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start listen on 8080 error")
	}
}
