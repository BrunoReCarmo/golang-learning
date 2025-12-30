package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// Register a handler for the /health endpoint.
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		// Example: /product/1900 -> id = "1900"
		id := r.URL.Path[len("/product/"):]
		//%s === string expected
		fmt.Fprintf(w, "The product id you requested is: %s", id)
	})

	// Notify that the server is starting.
	log.Println("Init server...")

	// Start a TCP listener on localhost:8080
	// Nill == NULL
	// !error console log FATAL ERROR
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err) // Log and exit if there's an error starting the listener
	}

	// Start the HTTP server in a separate goroutine, so main() can continue.
	go func() {
		log.Fatal(http.Serve(l, nil)) // Attach default serve mux (registered handlers)
	}()

	// Notify that we're about to send a client request to our own server
	log.Println("Sending request...")

	// Make a GET request to the running HTTP server's /health endpoint
	res, err := http.Get("http://localhost:8080/product/6969")
	if err != nil {
		log.Fatal(err) // Log and exit if request fails
	}
	defer res.Body.Close() // Always close the response body when done
	// Notify that the response will be read
	log.Println("Reading response...")

	// Stream the response body to Stdout (the console)
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err) // Log and exit if copying response fails
	}
}
