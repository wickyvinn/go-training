package main

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	for input, want := range map[int]int{
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
	} {
		name := fmt.Sprintf("square(%d)", input)
		t.Run(name, func(t *testing.T) {
			if have := square(input); want != have {
				t.Errorf("want %d, have %d", want, have)
			}
		})
	}
}
