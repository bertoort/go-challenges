package main

// ***
// http://golang.org/pkg/strings/
// ***

import (
	"strings"
)

func main() {

}

func lowerCase(word string) string {
	return strings.ToLower(word)
}

func contains(word, w string) bool {
	return strings.Contains(word, w)
}

func repeat(word string, num int) string {
	return strings.Repeat(word, num) + " batman!"
}
