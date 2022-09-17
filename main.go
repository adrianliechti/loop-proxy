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
	var addressFlag string
	var portFlag int

	var targetFlag string

	var keyFile string
	var certFile string

	flag.StringVar(&addressFlag, "address", "", "local address")
	flag.IntVar(&portFlag, "port", 8080, "lcoal proxy port")

	flag.StringVar(&targetFlag, "target", "", "target address")

	flag.StringVar(&keyFile, "key-file", "", "tls key file")
	flag.StringVar(&certFile, "cert-file", "", "tls certificate file")

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

	addr := fmt.Sprintf("%s:%d", addressFlag, portFlag)

	if certFile != "" {
		log.Fatal(http.ListenAndServeTLS(addr, certFile, keyFile, nil))

	} else {
		log.Fatal(http.ListenAndServe(addr, nil))
	}
}
