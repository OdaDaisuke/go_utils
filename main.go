package main

import (
	"github.com/OdaDaisuke/go_utils/timing"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("num goroutine: ", runtime.NumGoroutine())
	timing.SetInterval(func () {
		time.Sleep(5 * time.Second)
		fmt.Println("num goroutine: ", runtime.NumGoroutine())
	}, 1, true)
	//runtime.GC()
	//debug.FreeOSMemory()
	time.Sleep(time.Hour)
}
