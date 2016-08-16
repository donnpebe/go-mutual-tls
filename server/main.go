package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"

	"github.com/peefourtee/go-mutual-tls"
)

var (
	listen  = flag.String("listen", "127.0.0.1:8443", "http server listen address")
	crtFile = flag.String("crt", "valid/server.crt", "server crt file")
	keyFile = flag.String("key", "valid/server.key", "server key file")

	caCrtFile = flag.String("ca-crt", "valid/YoloSwagginsCA.crt", "ca crt file")
)

func main() {
	flag.Parse()

	log.Print("https server crt: ", *crtFile)
	log.Print("https server key: ", *keyFile)
	log.Print("ca crt: ", *caCrtFile)

	server := http.Server{
		Addr: *listen,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, crt := range r.TLS.PeerCertificates {
				log.Printf("client %q %q: signed by %s", crt.Subject.CommonName, crt.IPAddresses, crt.Issuer.CommonName)
			}
			w.Write([]byte("hello world"))
		}),
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  mutual_tls.MustLoadCAPool(*caCrtFile),
		},
	}

	log.Printf("starting server at %s", server.Addr)
	if err := server.ListenAndServeTLS(*crtFile, *keyFile); err != nil {
		log.Panic(err)
	}
}
