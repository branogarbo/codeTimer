package main

import (
	"fmt"
	"sync"
	"time"
)

func RunTests(funcs map[interface{}]func() interface{}) {
	var (
		wg          sync.WaitGroup
		testOutputs []TestOutput
	)

	testOutputChan := make(chan TestOutput)

	wg.Add(len(funcs))

	fmt.Println("starting tests...")

	for funcName, function := range funcs {
		go func(funcName interface{}, function func() interface{}) {
			startTime := time.Now().UTC()

			out := function()

			duration := time.Now().UTC().Sub(startTime)

			wg.Done()

			testOutput := TestOutput{
				FuncName: funcName,
				Duration: duration,
				Out:      out,
			}

			testOutputs = append(testOutputs, testOutput)
			testOutputChan <- testOutput
		}(funcName, function)
	}

	wg.Wait()

	select {
	case firstOut := <-testOutputChan:
		fmt.Printf(`test "%v" finished first: %v`+"\n", firstOut.FuncName, firstOut)
	default:
		fmt.Println("no test provided an output")
	}

	fmt.Println("tests:", testOutputs)
}
