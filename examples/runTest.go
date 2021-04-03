package main

import (
	"fmt"
	"log"
	"time"

	ct "github.com/branogarbo/codeTimer"
)

func main() {
	funcMap := map[interface{}]func() interface{}{
		"fast func": func() interface{} {
			time.Sleep(1 * time.Second)

			return "first test output"
		},
		"slow func": func() interface{} {
			time.Sleep(2 * time.Second)

			return "second test output"
		},
		"empty func": func() interface{} {
			return nil
		},
		"inc": func() interface{} {
			num := 0

			for i := 0; i < 99999999; i++ {
				num++
			}

			return num
		},
	}

	firstOutput, outputs, err := ct.RunTests(funcMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("first output:", firstOutput)
	for _, output := range outputs {
		fmt.Println(output)
	}

}
