package hackerrank

import (
	"fmt"
	"testing"
)

type testpair struct {
	x1, v1, x2, v2 int32
	result         string
}

var tests = []testpair{
	{0, 3, 4, 2, "YES"},
	{0, 2, 5, 3, "NO"},
}

func TestKangoroo(t *testing.T) {
	for _, pair := range tests {
		r := kangaroo(pair.x1, pair.v1, pair.x2, pair.v2)
		if r != pair.result {
			t.Error(fmt.Sprintf(
				"For [%v, %v, %v, %v] expected %v, got %v",
				pair.x1, pair.v1, pair.x2, pair.v2, pair.result, r,
			))
		}
	}
}
