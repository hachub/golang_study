package main

import (
	"fmt"
	"testing"
)

type testpair struct {
	x1, y1, x2, y2 float64
	result         float64
}

var tests = []testpair{
	{2, 2, 4, 2, 2},
	{2, 2, 10, 2, 8},
}

func TestDistance(t *testing.T) {
	for _, pair := range tests {
		r := distance(pair.x1, pair.y1, pair.x2, pair.y2)

		if r != pair.result {
			t.Error(fmt.Sprintf("For [%v,%v,%v,%v] expected %v, got %v",
				pair.x1, pair.y1, pair.x2, pair.y2,
				pair.result, r,
			))
		}
	}
}
