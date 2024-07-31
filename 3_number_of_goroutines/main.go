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
	n      = 5000000
	result = `Number of goroutines: %d
Total:
  Time:          %.2f s
  Memory:        %.2f Gb
  Scanned stack last GC: %.2f Mb
Per goroutine:
  Time:          %.2f µs
  Memory:        %.2f Kb
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

	totalTime, totalMemory, scannedStack := getTimeAndMemory()
	fmt.Printf(result, n, totalTime, totalMemory/(1<<30), scannedStack/(1<<20), totalTime/n*1e6, totalMemory/n)
}

func getTimeAndMemory() (float64, float64, float64) {
	s := []metrics.Sample{
		{Name: "/cpu/classes/user:cpu-seconds"},
		{Name: "/memory/classes/total:bytes"},
		{Name: "/gc/scan/stack:bytes"},
	}
	metrics.Read(s)
	return s[0].Value.Float64(), float64(s[1].Value.Uint64()), float64(s[2].Value.Uint64())
}
