package main

import (
	"fmt"
)

func core(input string) bool {
	var l = len(input)
	for i, c := range input {
		if i == 0 {
			if c != '<' {
				return false
			}
		} else if i == l-1 {
			if c != '>' {
				return false
			}
		} else {
			if c != '=' {
				return false
			}
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scan(&str)
	if core(str) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
