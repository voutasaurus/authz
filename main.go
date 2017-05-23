package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		log.Fatal(r.Status)
	}
	fmt.Println(extractLink(r.Header.Get("Link")))
}

func extractLink(h string) string {
	b := strings.Index(h, "<")
	e := strings.Index(h, ">")
	if b < -1 || e < -1 {
		return "error, link not found"
	}
	return h[b+1 : e]
}
