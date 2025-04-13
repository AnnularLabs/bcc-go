package main

import (
	"io"
	"os"
)

func main() {
	for {
		f, err := os.Open("/usr/share/dict/words")
		if err != nil {
			panic(err)
		}
		io.ReadAll(f)
		f.Close()
	}
}
