package main

import (
	"fmt"
	"sync"
	"time"
)

type TestOutput struct {
	FuncName interface{}
	Duration time.Duration
	Out      interface{}
}

func main() {
	var (
		wg          sync.WaitGroup
		testOutputs []TestOutput
	)

	testOutputChan := make(chan TestOutput)

	funcs := map[interface{}]func() interface{}{
		"fast func": func() interface{} {
			time.Sleep(1 * time.Second)

			return "first test output"
		},
		"slow func": func() interface{} {
			time.Sleep(2 * time.Second)

			return "second test output"
		},
	}

	/////////////////////////////

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
