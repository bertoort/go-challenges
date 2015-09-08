package main

func main() {

}

func hello() string {
	return "hello world"
}

func passingNumber(num int) int {
	return num
}

func passingString(word string) string {
	return "hello " + word
}

func multipleReturns(word string, num int) (string, int) {
	return word, num
}
