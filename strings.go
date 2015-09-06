package main

// ***
// http://golang.org/pkg/strings/
// ***

import ("fmt"; "strings")

func main() {
  lowerCase("HeLLo WORld");
  contains("team", "I");
  contains("team", "tea");
  repeat("na", 10);
}

func lowerCase(word string) {
  fmt.Println(strings.ToLower(word))
}

func contains(word, w string) {
  fmt.Println(strings.Contains(word, w))
}

func repeat(word string, num int) {
  fmt.Println(strings.Repeat(word, num) + " batman!")
}
