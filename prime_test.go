package ifPrime

import ("testing"; "fmt"; "github.com/bertort/go-challenges")

func Test_ifPrime(t *testing.T) {
  if ifPrime(5) != true {
    fmt.Println(ifPrime(5))
    t.Error()
  }
}
