package main

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/peefourtee/go-mutual-tls"
)

var (
	crtFile = flag.String("crt", "", "client crt file")
	keyFile = flag.String("key", "", "client key file")

	caCrtFile = flag.String("ca-crt", "valid/YoloSwagginsCA.crt", "ca crt file")
)

func main() {
	flag.Parse()

	log.Print("https client crt: ", *crtFile)
	log.Print("https client key: ", *keyFile)
	log.Print("ca crt: ", *caCrtFile)

	crt, err := tls.LoadX509KeyPair(*crtFile, *keyFile)
	if err != nil {
		log.Fatal("failed to load client keypair: ", err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{crt},
				RootCAs:      mutual_tls.MustLoadCAPool(*caCrtFile),
			},
		},
	}

	resp, err := client.Get("https://127.0.0.1:8443")
	if err != nil {
		log.Fatal("failed to get http response: ", err)
	}
	defer resp.Body.Close()

	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Printf("failed to read response: %s", err)
	} else {
		log.Printf("response: %q", data)
	}
}
