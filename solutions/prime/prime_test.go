package main

import (
	"fmt"
	"testing"
)

func Test_ifPrime(t *testing.T) {
	if ifPrime(5) != true {
		fmt.Println(ifPrime(5))
		t.Error()
	}
	if nextPrime(5) != 7 {
		fmt.Println(nextPrime(5), "does not equal 7")
		t.Error()
	}
}
