package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)


func main() {

	responseHandler := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "%s\n", os.Getenv("MESSAGE"))
	}

	http.HandleFunc("/", responseHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
