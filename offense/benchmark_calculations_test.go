package offense

import "testing"

func BenchmarkReadCSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadCSV()
	}
}

func BenchmarkCreateBatterCollection(b *testing.B) {
	records := ReadCSV()
	for i := 0; i < b.N; i++ {
		CreateBatterCollection(records)
	}

}

func BenchmarkFanOutProcessing(b *testing.B) {
	records := ReadCSV()
	for i := 0; i < b.N; i++ {
		FanOutProcessing(records)
	}
}

func BenchmarkWorkerPoolProcessing(b *testing.B) {
	records := ReadCSV()
	for i := 0; i < b.N; i++ {
		WorkerPoolProcessing(records)
	}
}
