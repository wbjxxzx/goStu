package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		fmt.Print(string(data))
		for _, line := range strings.Split(string(data), "\r\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		fmt.Printf("%s\t%d\n", line, n)
	}
}
