package main

import (
)

var buffer = make([]float32, 1024) // Buffer size must be appropriate for your use case

type Device string

const (
	BlackHole2ch  Device = "BlackHole 2ch"
	DefaultDevice Device = ""
)

func main() {
	initSinewave()
	// inputStream, err := initAudio(buffer, BlackHole2ch)
	// if err != nil {
	// 	log.Fatalf("Error initializing audio: %v", err)
	// }
	// defer inputStream.Close()

	// err = inputStream.Start()
	// if err != nil {
	// 	log.Fatalf("Error starting audio inputStream: %v", err)
	// }
	// defer inputStream.Stop()

	// // Initialize output stream
	// outputStream, err := initOutput(buffer)
	// if err != nil {
	// 	log.Fatalf("Error initializing output stream: %v", err)
	// }
	// defer outputStream.Close()

	// // Start the output stream
	// err = outputStream.Start()
	// if err != nil {
	// 	log.Fatalf("Error starting output stream: %v", err)
	// }
	// defer outputStream.Stop()

	// for {
	// 	err = inputStream.Read()
	// 	if err != nil {
	// 		log.Printf("Error reading audio: %v", err)
	// 		continue
	// 	}

	// 	err = outputStream.Write()
	// 	if err != nil {
	// 		log.Fatalf("Error writing to output stream: %v", err)
	// 	}

	// 	pitch := processAudio(buffer) // Pass the buffer directly
	// 	fmt.Printf("Detected pitch: %f Hz\n", pitch)
	// }
}
