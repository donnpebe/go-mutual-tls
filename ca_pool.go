package mutual_tls

import (
	"crypto/x509"
	"io/ioutil"
)

func MustLoadCAPool(crtFile string) *x509.CertPool {
	b, err := ioutil.ReadFile(crtFile)
	if err != nil {
		panic("failed to load CA crt: " + err.Error())
	}
	pool := x509.NewCertPool()

	if !pool.AppendCertsFromPEM(b) {
		panic("cannot add cert bytes to cert pool")
	}
	return pool
}
