package main

import (
	"math"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	// Запускаем трассировку.
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	// Make a copy of MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	// Пример вызова функции с параметрами.
	run(10000000, 20, 4)
}

func run(maxPrimeNumber, goroutineNumber, maxProcs int) {
	// Устанавливаем количество используемых процессоров.
	runtime.GOMAXPROCS(maxProcs)

	// Используем WaitGroup для того, чтобы дождаться выполнения всех горутин.
	wg := &sync.WaitGroup{}
	wg.Add(goroutineNumber)

	findInRange(wg, maxPrimeNumber, goroutineNumber) // Запускам CPU-bound задачу.
	wg.Wait()
}

// Ищем простые числа в заданном количестве горутин.
func findInRange(wg *sync.WaitGroup, maxPrimeNum, gNum int) {
	step := maxPrimeNum / gNum
	for i := 0; i < gNum; i++ {
		start := i * step
		end := (i + 1) * step
		if i == gNum-1 {
			end = maxPrimeNum
		}
		go func(start, end int) {
			//runtime.LockOSThread() --- вариант с привязкой горутины к конкретному потоку.
			//defer runtime.UnlockOSThread()
			defer wg.Done()
			findPrimeNumbers(start, end)
		}(start, end)
	}
}

// Функция, которая находит простые числа в заданном диапазоне.
func findPrimeNumbers(start, end int) []int {
	var primes []int

	for num := start; num <= end; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}

	return primes
}

// Проверка, что число является простым.
func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	limit := int(math.Sqrt(float64(num)))
	for i := 2; i <= limit; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
