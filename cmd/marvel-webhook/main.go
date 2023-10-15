package main

import (
	"log"
	"net/http"
	"path/filepath"
)

const (
	tlsDir      = `/etc/webhook/certs`
	tlsCertFile = `tls.crt`
	tlsKeyFile  = `tls.key`
)

func main() {
	certPath := filepath.Join(tlsDir, tlsCertFile)
	keyPath := filepath.Join(tlsDir, tlsKeyFile)

	http.HandleFunc("/add-marvel-label", handleAddMarvelLabel)

	server := &http.Server{
		Addr: ":8443",
	}
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
}
