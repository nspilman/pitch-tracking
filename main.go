package main

import (
	"fmt"
	"log"
)

var buffer = make([]float32, 2048) // Buffer size must be appropriate for your use case

type Device string

const (
	BlackHole2ch  Device = "BlackHole 2ch"
	DefaultDevice Device = ""
)

func main() {

	stream, err := initAudio(buffer, BlackHole2ch)
	if err != nil {
		log.Fatalf("Error initializing audio: %v", err)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		log.Fatalf("Error starting audio stream: %v", err)
	}
	defer stream.Stop()

	for {
		err = stream.Read()
		if err != nil {
			log.Printf("Error reading audio: %v", err)
			continue
		}

		pitch := processAudio(buffer) // Pass the buffer directly
		fmt.Printf("Detected pitch: %f Hz\n", pitch)
	}
}
