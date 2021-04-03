// Copyright 2021 Brian Longmore

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 		http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

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
		fmt.Println("no value sent to testOutputChan")
	}

	fmt.Println("tests:", testOutputs)
}
