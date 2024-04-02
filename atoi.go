package main

import (
	"strconv"
)

func Atoi(s string) int {
	result := 0
	sign := 1
	for i, char := range s {
		if i == 0 && (char == '+' || char == '-') {
			if char == '-' {
				sign = -1
			}
			continue
		}
		if char >= '0' && char <= '9' {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				return 0
			}
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result * sign
}
