package main

func main() {

}

func ifPrime(num int) bool {
	prime := true
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			prime = false
		}
	}
	return prime
}

func nextPrime(num int) int {
	i := num + 1
	for {
		if ifPrime(i) {
			break
		}
		i++
	}
	return i
}
