package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	words := strings.Split(strings.TrimSpace(string(b)), " ")

	dict := map[string]int{}
	for _, word := range words {

		count, ok := dict[word]
		if ok {
			dict[word] = count + 1 // dict[word]++
		} else {
			dict[word] = 1
		}
	}

	for word, count := range dict {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// bonus do it without buffering the stdin to memory
// hint: bufio.NewScanner, bufio.ScanWords
