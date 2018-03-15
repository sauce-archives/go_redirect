package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 80, "Port to listen on")
	url  = flag.String("url", "https://saucelabs.com/blog", "Url to redirect to")
)

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, *url, http.StatusPermanentRedirect)
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func main() {
	flag.Parse()

	fmt.Printf("Listening on port %d\n", *port)
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/", redirect)
	httpErr := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if httpErr != nil {
		log.Fatal(httpErr)
	}

}
