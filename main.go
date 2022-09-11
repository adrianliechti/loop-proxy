package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var portFlag int
	var targetFlag string

	flag.IntVar(&portFlag, "port", 8080, "lcoal proxy port")
	flag.StringVar(&targetFlag, "target", "", "target address")

	flag.Parse()

	if targetFlag == "" {
		log.Fatal("target url is required")
	}

	target, err := url.Parse(targetFlag)

	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host

		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", portFlag), nil))
}
