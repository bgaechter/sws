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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	sws_port := ":" + getEnv("SWS_PORT", "8000")
	responseHandler := func(w http.ResponseWriter, _ *http.Request) {
		res, err := json.Marshal(ResponseObject{Message: os.Getenv("SWS_MESSAGE")})
		if err != nil {
			log.Fatal(err)
		}
		w.Write(res)
	}

	http.HandleFunc("/", responseHandler)
	log.Fatal(http.ListenAndServe(sws_port, nil))
}
