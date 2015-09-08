package main

import (
	"fmt"
	"testing"
)

func Test_hello(t *testing.T) {
	if hello() != "hello world" {
		fmt.Println(hello())
		t.Error("should say hello world")
	}
}

func Test_passingNumber(t *testing.T) {
	if passingNumber(1) != 1 {
		fmt.Println(passingNumber(1))
		t.Error("should say 1")
	}
}

func Test_passingString(t *testing.T) {
	if passingString("friend") != "hello friend" {
		fmt.Println(passingString("friend"))
		t.Error("should say hello friend")
	}
}

func Test_multipleReturns(t *testing.T) {
	word, num := multipleReturns("hello", 5)
	if word != "hello" {
		fmt.Println(word)
		t.Error("should be hello")
	}
	if num != 5 {
		fmt.Println(num)
		t.Error("should be 5")
	}
}
