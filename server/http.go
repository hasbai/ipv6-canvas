package server

import (
	"log"
	"net/http"
	"strconv"
)

func getImage(w http.ResponseWriter, r *http.Request) {
	data, err := IMG.Encode()
	if err != nil {
		log.Println("unable to encode image.")
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	_, err = w.Write(data)
	if err != nil {
		log.Println("unable to write image.")
	}
}

func serveHTTP(httpAddr string) {
	log.Println("HTTP listening on", httpAddr)
	err := http.ListenAndServe(httpAddr, nil)
	if err != nil {
		log.Fatalf("Error serving http: %s", err)
	}
}

func init() {
	http.HandleFunc("/image", getImage)
}
