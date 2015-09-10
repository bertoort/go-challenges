package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_matrix(t *testing.T) {
	if !reflect.DeepEqual(twoByTwo(), [][]int{[]int{1, 0}, []int{0, 1}}) {
		fmt.Println(twoByTwo())
		t.Error("should be [[1 0] [0 1]]")
	}
}
