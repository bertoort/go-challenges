package main

import (
	"fmt"
	"testing"
)

func Test_roman(t *testing.T) {
	if roman("IV") != 4 {
		fmt.Println(roman("IV"))
		t.Error("does not equal 4")
	}
	if roman("XI") != 11 {
		fmt.Println(roman("XI"))
		t.Error("does not equal 11")
	}
	if roman("C") != 100 {
		fmt.Println(roman("C"))
		t.Error("does not equal 100")
	}
	if roman("MCMVII") != 1907 {
		fmt.Println(roman("MCMVII"))
		t.Error("does not equal 1907")
	}
}
