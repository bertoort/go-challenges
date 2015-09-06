package main

import("fmt"; "math")

func main() {
  mod(10, 2);
  mod(2, 10);
}

func mod(num1, num2 float64) {
  fmt.Println(math.Mod(num1, num2))
}
