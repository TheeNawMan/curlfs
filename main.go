package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	directoryPath = "."        // Directory to serve and store uploads
	certFile      = "cert.pem" // TLS certificate
	keyFile       = "key.pem"  // TLS private key
	httpsPort     = 8443       // HTTPS port
	httpPort      = 8080       // HTTP fallback port
)

func main() {
	// Check if directory exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}

	// Unified handler for GET and POST
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.FileServer(http.Dir(directoryPath)).ServeHTTP(w, r)
		case http.MethodPost:
			uploadHandler(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Check if certs exist
	if fileExists(certFile) && fileExists(keyFile) {
		fmt.Printf("HTTPS server started at https://localhost:%d\n", httpsPort)
		err := http.ListenAndServeTLS(
			fmt.Sprintf(":%d", httpsPort),
			certFile,
			keyFile,
			nil,
		)
		if err != nil {
			fmt.Printf("Failed to start HTTPS: %v\n", err)
		}
	} else {
		fmt.Printf("TLS certs not found. Falling back to HTTP at http://localhost:%d\n", httpPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil)
		if err != nil {
			fmt.Printf("Failed to start HTTP server: %v\n", err)
		}
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Max 10MB
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dstPath := filepath.Join(directoryPath, handler.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Error creating file on server", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
