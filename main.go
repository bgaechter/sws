package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)


type ResponseObject struct {
	Message string
}

func main() {

	responseHandler := func(w http.ResponseWriter, _ *http.Request) {
		res, err := json.Marshal(ResponseObject{Message: os.Getenv("MESSAGE")})
		if err != nil {
			log.Fatal(err)
		}
		w.Write(res)
	}

	http.HandleFunc("/", responseHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
