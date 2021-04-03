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

package codeTimer

import (
	"errors"
	"sync"
	"time"
)

type TestOutput struct {
	FuncName   interface{}
	Duration   time.Duration
	FuncOutput interface{}
}

func RunTests(funcs map[interface{}]func() interface{}) (TestOutput, []TestOutput, error) {
	var (
		wg          sync.WaitGroup
		testOutputs []TestOutput
	)

	testOutputChan := make(chan TestOutput)

	wg.Add(len(funcs))

	for funcName, function := range funcs {
		go func(funcName interface{}, function func() interface{}) {
			startTime := time.Now().UTC()

			funcOutput := function()

			duration := time.Now().UTC().Sub(startTime)

			wg.Done()

			testOutput := TestOutput{
				FuncName:   funcName,
				Duration:   duration,
				FuncOutput: funcOutput,
			}

			testOutputs = append(testOutputs, testOutput)
			testOutputChan <- testOutput
		}(funcName, function)
	}

	wg.Wait()

	select {
	case firstOut := <-testOutputChan:
		return firstOut, testOutputs, nil
	default:
		return TestOutput{}, []TestOutput{}, errors.New("no value sent to testOutputChan")
	}
}
