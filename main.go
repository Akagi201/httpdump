package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.Write(w)
	io.Copy(w, r.Body)
}

func main() {
	port := flag.Int("port", 8080, "Listen on port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Starting httpdump server on port:", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
