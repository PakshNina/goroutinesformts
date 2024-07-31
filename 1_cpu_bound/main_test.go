package main

import "testing"

const (
	defaultMaxPrimeNumber = 10000000
)

// Тестируем для 4 и 20 горутин
func BenchmarkRun4Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(defaultMaxPrimeNumber, 4, 4)
	}
}

func BenchmarkRun20Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(defaultMaxPrimeNumber, 20, 4)
	}
}
