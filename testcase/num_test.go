package testcase_test

import (
	"go-learn/testcase"
	"testing"
)

func TestSum(t *testing.T) {
	sum := testcase.Sum(1, 2, 3)
	if sum != 6 {
		t.Error("Wrong result")
	}
}

func TestSub(t *testing.T) {
	sub := testcase.Sub(5, 3)
	if sub != 2 {
		t.Error("Wrong result")
	}
}
