package main

import("fmt")

func main() {
  fmt.Println(ifPrime(5));
  fmt.Println(ifPrime(6));
  nextPrime(12);
}

func ifPrime(num int) bool {
  prime := true
  for i := 2; i < num/2; i++ {
    if num % i == 0 {
      prime = false
    }
  }
  return prime
}

func nextPrime(num int) {
  i := num + 1;
  for {
    if ifPrime(i) {
      break;
    }
    i++
  }
  fmt.Println(i)
}
