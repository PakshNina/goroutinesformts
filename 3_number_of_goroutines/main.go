package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/metrics"
	"runtime/trace"
)

const (
	n      = 1000000
	result = `Total:
  Time:         %.2f s
  Memory:       %.2f Gb
  Stack memory: %.2f Gb
Per goroutine:
  Time:         %.2f µs
  Memory:       %.2f Kb
`
)

var ch = make(chan byte)

func goroutine() {
	<-ch // Блокируем горутину.
}

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	debug.SetGCPercent(-1) // Выключаем GC.
	runtime.GOMAXPROCS(1)  // Используем один процессор.

	for i := 0; i < n; i++ {
		go goroutine()
	}
	runtime.GC()

	totalTime, totalMemory, allStackMemory := getTimeAndMemory()
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf(result, totalTime, totalMemory/(1<<30), allStackMemory/(1<<30), totalTime/n*1e6, allStackMemory/n)
}

func getTimeAndMemory() (float64, float64, float64) {
	s := []metrics.Sample{
		{Name: "/cpu/classes/user:cpu-seconds"},
		{Name: "/memory/classes/total:bytes"},
		{Name: "/memory/classes/heap/stacks:bytes"},
	}
	metrics.Read(s)
	return s[0].Value.Float64(), float64(s[1].Value.Uint64()), float64(s[2].Value.Uint64())
}
