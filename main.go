package main

import (
	"math"
)

func hasher(pattern string) int {
	strMap := map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8}
	var hash int

	for i, rune := range pattern {
		// fmt.Println(i, strMap[rune], math.Pow(10.00, float64(len(pattern)-i)))
		exp := float64(strMap[rune]) * math.Pow(10.00, float64(len(pattern)-1-i))
		hash += int(exp)
	}
	return hash
}

func main() {
	str := "ccaccaaedba"
	target := "acc"
	pattern := hasher(target)
	for i := 0; i < len(str)-2; i++ {
		if (hasher(str[i : i+3])) == pattern {
			print("found")
		} else {
			print("better luck next time bitches")
		}
	}
}
