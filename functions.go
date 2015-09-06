package main

import ("fmt")

func main() {
  hello();
  passingString("world");
  passingNumber(1);
}

func hello() {
  fmt.Println("hello world")
}

func passingString(word string) {
  fmt.Println("hello " + word)
}

func passingNumber(num int) {
  fmt.Println("you are my number ", num);
}
