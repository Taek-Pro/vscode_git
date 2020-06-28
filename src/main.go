package main

import (
	"fmt"
	"runtime"
	"sync"
)

type GROBAL_DATA struct {
	nResult int
	mutex   sync.RWMutex
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // All CPU Use

	var mainData = new(GROBAL_DATA)

	for i := 0; i < 100; i++ {
		go AddResult(mainData, 1)
		go PrintResult(mainData)
	}

	fmt.Scanln()
	fmt.Printf("MainResult : %d\n", mainData.nResult)

	return
}

func AddResult(mainData *GROBAL_DATA, nValue int) {
	mainData.mutex.Lock()
	mainData.nResult += nValue

	mainData.mutex.Unlock()
}

func PrintResult(mainData *GROBAL_DATA) {
	mainData.mutex.RLock()
	fmt.Printf("Result : %d\n", mainData.nResult)
	mainData.mutex.RUnlock()
}
