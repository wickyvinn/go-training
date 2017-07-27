package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("please provide a filename.")
		return
	}
	filename := os.Args[1]
	if err := writeToStdOut(filename); err != nil {
		log.Fatalf("%s", err)
		return
	}

}

func writeToStdOut(filename string) error {
	if b, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		s := string(b)
		log.Println(s)
	}

	return nil

}
