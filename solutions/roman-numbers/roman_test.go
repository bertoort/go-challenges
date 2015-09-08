package main

import (
	"fmt"
	"testing"
)

func Test_roman(t *testing.T) {
	if roman("IV") != 4 {
		fmt.Println(roman("IV"), "does not equal 4")
		t.Error()
	}
	if roman("XI") != 11 {
		fmt.Println(roman("XI"), "does not equal 11")
		t.Error()
	}
	if roman("C") != 100 {
		fmt.Println(roman("C"), "does not equal 100")
		t.Error()
	}
	if roman("MCMVII") != 1907 {
		fmt.Println(roman("MCMVII"), "does not equal 1907")
		t.Error()
	}
}
