package main

import (
	"fmt"
	"testing"
)

func BenchmarkSquare(b *testing.B) {
	for _, n := range []int{9, 999, 999999999, 9999999999} {
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				square(n)
			}
		})
	}
}

// var fibTests = []struct {
//   n        int // input
//   expected int // expected result
// }{
//   {1, 1},
//   {2, 1},
//   {3, 2},
//   {4, 3},
//   {5, 5},
//   {6, 8},
//   {7, 13},
// }

func BenchmarkRace(b *testing.B) {
	var races = []struct {
		distance     int
		tortoisePace int
		hareNap      int
	}{
		{20, 2, 15},
		{30, 3, 13},
		{5, 2, 1},
	}
	for raceNumber, r := range races {
		t := tortoise{pace: r.tortoisePace}
		h := hare{nap: r.hareNap}
		b.Run(fmt.Sprint(raceNumber), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				race(r.distance, []runner{t, h}...)
			}
		})
	}
}
