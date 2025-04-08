package main

import (
	"log"
	"os"
	"time"
)

func main() {
	deviceName := os.Getenv("DEVICE_NAME")
	if deviceName == "" {
		log.Fatal("DEVICE_NAME not set")
	}

	for {
		log.Printf("Device %s is running. Sending mock data...\n", deviceName)
		time.Sleep(10 * time.Second)
	}
}
