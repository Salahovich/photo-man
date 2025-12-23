package main

import (
	"os"
	"photo-man/ui"
	"runtime"
	"runtime/pprof"
)

func main() {
	ui.StartApp()
}

func startProfiling() {
	f1, _ := os.Create("start.prof")
	err1 := pprof.WriteHeapProfile(f1)
	if err1 != nil {
		return
	}
	err2 := f1.Close()
	if err2 != nil {
		return
	}
}

func endProfiling() {
	runtime.GC()

	// 3. Write the heap profile to the file
	f2, _ := os.Create("end.prof")
	err3 := pprof.WriteHeapProfile(f2)
	if err3 != nil {
		return
	}
	err4 := f2.Close()
	if err4 != nil {
		return
	}
}
