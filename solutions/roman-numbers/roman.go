package main

import (
	"strings"
)

func main() {

}

func forRoman(letters string) int {
	split := strings.Split(letters, "")
	sum := 0
	skip := false
	for k, v := range split {
		if skip == false {
			curr := convertRoman(strings.ToLower(v))
			next := 0
			if len(split)-1 != k {
				next = convertRoman(strings.ToLower(split[k+1]))
			}
			if curr >= next {
				sum += curr
			} else {
				sum += next - curr
				skip = true
			}
		} else {
			skip = false
		}
	}
	return sum
}

func recursionRoman(letters string, total int) int {
	split := strings.Split(letters, "")
	if len(split) == 0 {
		return total
	} else if len(split) == 1 {
		total += convertRoman(strings.ToLower(split[0]))
		return total
	}
	curr := convertRoman(strings.ToLower(split[0]))
	next := convertRoman(strings.ToLower(split[1]))
	if curr >= next {
		total += curr
		return recursionRoman(strings.Join(split[1:], ""), total)
	}
	total += next - curr
	return recursionRoman(strings.Join(split[2:], ""), total)
}

func convertRoman(letter string) int {
	num := 1
	switch letter {
	case "i":
		num = 1
	case "v":
		num = 5
	case "x":
		num = 10
	case "l":
		num = 50
	case "c":
		num = 100
	case "d":
		num = 500
	case "m":
		num = 1000
	}
	return num
}
