package main

import (
	"fmt"
	"log"
)

var buffer = make([]float32, 64) // Buffer size must be appropriate for your use case

func init() {
	stream, err := initAudio()
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
