package main

import (
	"math"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

const (
	maxPrimeNumber = 10000000
	maxProc        = 1
)

func main() {
	runtime.GOMAXPROCS(maxProc)

	// Запускаем трассировку.
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()
	// Пример вызова функции с параметрами.
	run(4, maxProc)
}

func run(goroutineNumber, maxProcNum int) {
	// Устанавливаем количество используемых процессоров.
	runtime.GOMAXPROCS(maxProcNum)

	// Используем WaitGroup для того, чтобы дождаться выполнения всех горутин.
	wg := &sync.WaitGroup{}
	wg.Add(goroutineNumber)

	findInRange(wg, goroutineNumber) // Запускам CPU-bound задачу.
	wg.Wait()
}

// Ищем простые числа в заданном количестве горутин.
func findInRange(wg *sync.WaitGroup, gNum int) {
	step := maxPrimeNumber / gNum
	for i := 0; i < gNum; i++ {
		start := i * step
		end := (i + 1) * step
		if i == gNum-1 {
			end = maxPrimeNumber
		}
		go func(start, end int) {
			runtime.LockOSThread() // вариант с привязкой горутины к конкретному потоку.
			defer runtime.UnlockOSThread()
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
