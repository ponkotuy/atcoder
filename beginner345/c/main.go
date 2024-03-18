package main

import "fmt"

func circular2(x int) int {
	return x * (x - 1) / 2
}

func core(input string) int {
	cMap := make(map[rune]int)
	for _, c := range input {
		cMap[c] += 1
	}
	dup := 0
	for _, value := range cMap {
		if 2 <= value {
			dup += circular2(value)
		}
	}
	if 0 < dup {
		return circular2(len(input)) - (dup - 1)
	}
	return circular2(len(input))
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println(core(input))
}
