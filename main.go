package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var portNum = flag.String("listen", ":8080", "listen string")
var targetHost = flag.String("host", "https://example.com", "host prefix to forward connections to")

func main() {

	flag.Parse()

	http.HandleFunc("/submit", process)
	http.HandleFunc("/token", process)

	log.Fatal(http.ListenAndServe(*portNum, nil))
}

func process(w http.ResponseWriter, req *http.Request) {
	target := fmt.Sprintf("%v%v", *targetHost, req.URL.Path)

	log.Printf("Forwarding request from %v to %v", req.URL, target)

	resp, err := http.Post(target, "application/json", req.Body)
	if err != nil {
		log.Printf("could not make request to %v, %v", target, err)
		http.Error(w, "could not make request "+err.Error(), 500)
		return
	}
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("could not output result from %v, %v", target, err)
		http.Error(w, "could not output result "+err.Error(), 500)
		return
	}
}
