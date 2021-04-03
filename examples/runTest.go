package main

import (
	"log"
	"time"

	util "github.com/branogarbo/funcTimer"
	img "github.com/branogarbo/imgcli/util"
)

func main() {

	util.RunTests(map[interface{}]func() interface{}{
		"fast func": func() interface{} {
			time.Sleep(1 * time.Second)

			return "first test output"
		},
		"slow func": func() interface{} {
			time.Sleep(2 * time.Second)

			return "second test output"
		},
		"imgcli conversion": func() interface{} {
			_, err := img.OutputImage(img.OutputConfig{
				Src:          "../imgcli/examples/images/portrait.jpg",
				OutputMode:   "ascii",
				AsciiPattern: " .:-=+*#%@",
				OutputWidth:  500,
			})
			if err != nil {
				log.Fatal(err)
			}

			return nil
		},
		"empty func": func() interface{} {
			return nil
		},
	})

}
