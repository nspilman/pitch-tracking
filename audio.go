package main

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

func showDevices() {
	err := portaudio.Initialize()
	if err != nil {
		fmt.Println(err)
	}
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
	var inputDevice *portaudio.DeviceInfo
	if deviceName != "" {
		inputDevice, err = findDeviceByName(string(deviceName))
	} else {
		inputDevice, err = portaudio.DefaultInputDevice()
	}
	if err != nil {
		return nil, err
	}
	streamParameters := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device:   inputDevice,
			Channels: 1,
			Latency:  inputDevice.DefaultLowInputLatency,
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


func initOutput(buffer []float32, deviceName Device) (*portaudio.Stream, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, err
	}
	var outputDevice *portaudio.DeviceInfo
	if deviceName != "" {
		outputDevice, err = findDeviceByName(string(deviceName))
	} else {
		outputDevice, err = portaudio.DefaultInputDevice()
	}

	if err != nil {
		return nil, err
	}

	streamParameters := portaudio.StreamParameters{
		Output: portaudio.StreamDeviceParameters{
			Device:   outputDevice,
			Channels: 1,
			Latency:  outputDevice.DefaultLowOutputLatency,
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
