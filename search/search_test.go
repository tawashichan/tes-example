package search

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Input Input
	Want  int
}

type Input struct {
	nums   []int
	target int
}

func Test_BSearch1(t *testing.T) {
	testCase := TestCase{
		Input: Input{
			nums:   []int{1, 2, 3, 5},
			target: 3,
		},
		Want: 2,
	}
	ExecTestBsearch(t, testCase)
}

func ExecTestBsearch(t *testing.T, testCase TestCase) {
	input := testCase.Input
	result := BSearch(input.nums, input.target)
	if result != testCase.Want {
		t.Error(fmt.Sprintf("want: %v,result: %v", testCase.Want, result))
	}
}
