package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 80, "Port to listen on")
	url  = flag.String("url", "https://saucelabs.com/blog", "Url to redirect to")
)

func main() {
	flag.Parse()

	httpErr := http.ListenAndServe(fmt.Sprintf(":%d", *port), http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, *url, http.StatusPermanentRedirect)
		}))
	if httpErr != nil {
		log.Fatal(httpErr)
	}

	fmt.Printf("Listening on port %d\n", *port)
}
