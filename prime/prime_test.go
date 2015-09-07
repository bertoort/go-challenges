package ifPrime

import ("testing"; "fmt")

func Test_ifPrime(t *testing.T) {
  if ifPrime(5) != true {
    fmt.Println(ifPrime(5))
    t.Error()
  }
}
