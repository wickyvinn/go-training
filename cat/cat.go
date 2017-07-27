package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	count := flag.Bool("count", false, "print # of bytes")
	stderr := flag.Bool("stderr", false, "print to stderr")

	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalf("please provide a filename.")
		return
	}
	s, num, err := readfile(flag.Args()[0])
	if err != nil {
		log.Fatalf("error with reading file.")
		return
	}

	if *stderr == true {
		os.Stderr.WriteString(s)
	} else {
		fmt.Println(s)
	}

	if *count == true {
		fmt.Printf("# of bytes: %d\n", num)
	}
}

func readfile(filename string) (string, int, error) {
	if b, err := ioutil.ReadFile(filename); err != nil {
		return "", 0, err
	} else {
		return string(b), len(b), nil
	}
}
