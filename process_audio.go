package main

import (
	"math/cmplx"
	"gonum.org/v1/gonum/dsp/window"

	"gonum.org/v1/gonum/dsp/fourier"
)

func processAudio(in []float32) float64 {
	// Convert float32 to float64 for FFT
	data := make([]float64, len(in))
	for i, v := range in {
		data[i] = float64(v)
	}

	// Apply a window function to reduce spectral leakage
	window.Hann(data)

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
	for i := 0; i < len(coeff)/2; i++ { // Only consider the first half of the coefficients
		if abs := cmplx.Abs(coeff[i]); abs > maxVal {
			maxVal = abs
			maxIdx = i
		}
	}
	// Calculate frequency
	return float64(maxIdx) * float64(sampleRate) / float64(len(coeff))
}