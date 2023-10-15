package main

import (
	"log"
	"net/http"
	"path/filepath"
)

// TLS certificate and key file paths
const (
	tlsDir      = `/etc/webhook/certs`
	tlsCertFile = `tls.crt`
	tlsKeyFile  = `tls.key`
)

func main() {
	// Define the paths for the TLS certificate and key files
	certPath := filepath.Join(tlsDir, tlsCertFile)
	keyPath := filepath.Join(tlsDir, tlsKeyFile)

	// Register the webhook handler function
	http.HandleFunc("/add-marvel-label", handleAddMarvelLabel)

	// Configure the HTTPS server
	server := &http.Server{
		Addr: ":8443",
	}
	// Start the HTTPS server
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
}
