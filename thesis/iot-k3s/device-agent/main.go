package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	deviceName := os.Getenv("DEVICE_NAME")
	if deviceName == "" {
		log.Fatal("DEVICE_NAME not set")
	}

	targetPort := os.Getenv("PORT")
	if targetPort == "" {
		log.Fatal("PORT not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Il device \"%s\" è pronto", deviceName)
		fmt.Fprintln(w, msg)
		log.Println(msg)
	})

	addr := fmt.Sprintf(":%s", targetPort)
	log.Printf("HTTP server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
