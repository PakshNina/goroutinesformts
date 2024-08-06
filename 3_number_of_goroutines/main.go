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
	result = `Общее:
  Время:        %.2f c
  Память:       %.2f Гб
  Память стека: %.2f Гб
На горутину:
  Время:        %.2f мкс
  Память:       %.2f Кб
`
)

var ch = make(chan byte)

func goroutine() {
	<-ch // Блокируем горутину.
}

func main() {
	runtime.GOMAXPROCS(1) // Используем один процессор.

	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	debug.SetGCPercent(-1) // Выключаем GC.

	for i := 0; i < n; i++ {
		go goroutine()
	}
	runtime.GC() // Запускаем GC.

	totalTime, totalMemory, allStackMemory := getTimeAndMemory()
	goNum := runtime.NumGoroutine()
	fmt.Printf("Количество горутин: %d\n", goNum)
	fmt.Printf(result, totalTime, totalMemory/(1<<30), allStackMemory/(1<<30), totalTime/n*1e6, allStackMemory/float64(goNum))
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
