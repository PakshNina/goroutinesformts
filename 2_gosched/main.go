package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // Запускаем на 1 процессоре.

	// Трассировка
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	var wg sync.WaitGroup
	wg.Add(2)

	go funcA(&wg)
	go funcB(&wg)

	wg.Wait()
}

func funcA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("A%d", i)
		runtime.Gosched()
	}
}

func funcB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("B%d", i)
		runtime.Gosched()
	}
}
