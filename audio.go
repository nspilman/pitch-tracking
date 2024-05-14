package main

import (
	"github.com/gordonklaus/portaudio"
)

// initAudio initializes an audio stream to capture audio from the microphone.
func initAudio(buffer []float32) (*portaudio.Stream, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}

	// Open the default audio device with a buffer of size 2048
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), &buffer)
	if err != nil {
		return nil, err
	}
	return stream, nil
}
