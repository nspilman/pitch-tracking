package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"gonum.org/v1/gonum/dsp/fourier"
)

func processAudio(in []float32) float64 {
	// Convert float32 to float64 for FFT
	data := make([]float64, len(in))
	for i, v := range in {
		data[i] = float64(v)
	}

	// Create an FFT plan
	fft := fourier.NewFFT(len(data))
	// This performs the FFT and returns complex coefficients
	coeff := fft.Coefficients(nil, data)

	// Find dominant frequency
	return findDominantFrequency(coeff)
}

func findDominantFrequency(coeff []complex128) float64 {
	maxVal := 0.0
	var maxIdx int
	for i, v := range coeff {
		if abs := cmplx.Abs(v); abs > maxVal {
			maxVal = abs
			maxIdx = i
		}
	}
	sampleRate := 44100 // Define as per your setup
	// Calculate frequency
	return float64(maxIdx) * float64(sampleRate) / float64(len(coeff))
}

func main() {
	// Example usage with dummy data
	input := make([]float32, 1024) // Dummy input data
	for i := range input {
		input[i] = float32(math.Sin(2 * math.Pi * float64(i) * 440 / 44100)) // Example: 440 Hz tone
	}

	pitch := processAudio(input)
	fmt.Printf("Detected pitch: %f Hz\n", pitch)
}
