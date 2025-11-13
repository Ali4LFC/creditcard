package main

import (
	"fmt"
	"os"
)

func luhnCheck(number string) bool {
	sum := 0
	pos := 0
	for i := len(number) - 1; i >= 0; i-- {
		ch := number[i]
		if ch == ' ' || ch == '\t' {
			continue
		}
		if ch < '0' || ch > '9' {
			fmt.Fprintln(os.Stderr, "Invalid character in number")
			return false
		}
		digit := int(ch - '0')
		if pos%2 == 1 {
			d := digit * 2
			if d > 9 {
				d -= 9
			}
			sum += d
		} else {
			sum += digit
		}
		pos++
	}
	if pos == 0 {
		fmt.Fprintln(os.Stderr, "Empty number")
		return false
	}
	return sum%10 == 0
}
