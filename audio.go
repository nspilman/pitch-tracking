package main

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

func showDevices() {
	err := portaudio.Initialize()
	devices, err := portaudio.Devices()
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	for _, device := range devices {
		fmt.Println(device, device.Name)
	}
}

func findDeviceByName(name string) (*portaudio.DeviceInfo, error) {
	devices, err := portaudio.Devices()
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if device.Name == name {
			return device, nil
		}
	}
	return nil, fmt.Errorf("device %s not found", name)
}

// initAudio initializes an audio stream to capture audio from the microphone.
func initAudio(buffer []float32, deviceName Device) (*portaudio.Stream, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}
	var device *portaudio.DeviceInfo
	if deviceName != "" {
		device, err = findDeviceByName(string(deviceName))
	} else {
		device, err = portaudio.DefaultInputDevice()
	}
	if err != nil {
		return nil, err
	}
	streamParameters := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device:   device,
			Channels: 1,
			Latency:  device.DefaultLowInputLatency,
		},
		SampleRate:      44100,
		FramesPerBuffer: len(buffer),
	}

	stream, err := portaudio.OpenStream(streamParameters, buffer)
	if err != nil {
		return nil, err
	}
	return stream, nil
}
