package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
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

	targetAddr := fmt.Sprintf("localhost:%s", targetPort)

	// Avvia il server HTTP per la readiness probe
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			msg := fmt.Sprintf("Sono il device %s", deviceName)
			fmt.Fprintln(w, msg)
			log.Println(msg)
		})
		log.Println("HTTP server listening on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Loop per inviare dati mock via TCP
	for {
		conn, err := net.Dial("tcp", targetAddr)
		if err != nil {
			log.Printf("Connection error to %s: %v", targetAddr, err)
			time.Sleep(5 * time.Second)
			continue
		}

		msg := fmt.Sprintf("Device %s says hello at %s\n", deviceName, time.Now().Format(time.RFC3339))
		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Printf("Send error: %v", err)
		} else {
			log.Printf("Sent to %s: %s", targetAddr, msg)
		}
		conn.Close()
		time.Sleep(10 * time.Second)
	}
}
