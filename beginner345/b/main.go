package main

import (
	"fmt"
)

func core(input int64) int64 {
	if 0 < input {
		return (input + 9) / 10
	} else {
		return input / 10
	}
}

func main() {
	var input int64
	fmt.Scan(&input)
	fmt.Println(core(input))
}
