package main

import (
	"fmt"
	"math"

	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate = 44100
	frequency  = 440.0 // A4 note
)

func sineWave(samples []float32, phase *float64) {
	for i := range samples {
		samples[i] = float32(math.Sin(2 * math.Pi * frequency * *phase / sampleRate))
		*phase++
	}
}

func initSinewave() {
	defer portaudio.Terminate()

	var phase float64
	buffer := make([]float32, 2048)
	stream, err := initOutput(buffer, BlackHole2ch)
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	err = stream.Start()
	if err != nil {
		panic(err)
	}

	for {
		sineWave(buffer, &phase)
		pitch := processAudio(buffer) // Pass the buffer directly
		fmt.Printf("Detected pitch: %f Hz\n", pitch)
		if err := stream.Write(); err != nil {
			panic(err)
		}
	}
}
