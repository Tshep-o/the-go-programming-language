//dup3 prints duplicate lines in files and their count

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) > 0 {
		for _, filename := range os.Args[1:] {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				counts[line]++
			}
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
