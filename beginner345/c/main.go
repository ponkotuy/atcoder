package main

import (
	"bufio"
	"fmt"
	"os"
)

func circular2(x int) int {
	return x * (x - 1) / 2
}

func core(input string) int {
	cMap := make([]int, 26)
	for _, c := range input {
		cMap[c-'a'] += 1
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
	rdr := bufio.NewReaderSize(os.Stdin, 1e6+10)
	fmt.Fscan(rdr, &input)
	fmt.Println(core(input))
}
