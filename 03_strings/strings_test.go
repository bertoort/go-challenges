package main

import (
	"fmt"
	"testing"
)

func Test_lowerCase(t *testing.T) {
	if lowerCase("HeLLo WORld") != "hello world" {
		fmt.Println(lowerCase("HeLLo WORld"))
		t.Error("should be hello world")
	}
}
//
// func Test_contains(t *testing.T) {
// 	if contains("team", "i") != false {
// 		fmt.Println(contains("team", "i"))
// 		t.Error("should be false")
// 	}
// }
//
// func Test_repeat(t *testing.T) {
// 	if repeat("na", 10) != "nananananananananana batman!" {
// 		fmt.Println(repeat("na", 10))
// 		t.Error("should be nananananananananana batman!")
// 	}
// }
