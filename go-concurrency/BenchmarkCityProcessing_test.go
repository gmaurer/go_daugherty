package main

import "testing"

func BenchmarkChannelWorkerPoolExample(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ChannelWorkerPoolExample()
	}

}

func BenchmarkFanOutExample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FanOutExample()
	}

}

func BenchmarkSerialProcessCities(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SerialProcessCities()
	}

}
