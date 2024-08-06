package main

import "testing"

// Тестируем для 4 и 4 горутин
func BenchmarkRun4Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(4, 4)
	}
}

// Тестируем для 4 и 20 горутин
func BenchmarkRun20Goroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run(20, 4)
	}
}
