package main

import (
	"fmt"
	"testing"
)

func Test_roman(t *testing.T) {
	if forRoman("IV") != 4 {
		fmt.Println(forRoman("IV"), "does not equal 4")
		t.Error()
	}
	if forRoman("XI") != 11 {
		fmt.Println(forRoman("XI"), "does not equal 11")
		t.Error()
	}
	if forRoman("C") != 100 {
		fmt.Println(forRoman("C"), "does not equal 100")
		t.Error()
	}
	if forRoman("MCMVII") != 1907 {
		fmt.Println(forRoman("MCMVII"), "does not equal 1907")
		t.Error()
	}
	if recursionRoman("IV", 0) != 4 {
		fmt.Println(recursionRoman("IV", 0), "does not equal 4")
		t.Error()
	}
	if recursionRoman("MCMVII", 0) != 1907 {
		fmt.Println(recursionRoman("MCMVII", 0), "does not equal 1907")
		t.Error()
	}
}
