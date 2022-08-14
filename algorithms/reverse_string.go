package main

import (
	io "fmt"
)

func ReverseString(str string) string {

	if len(str) == 0 {
		return ""
	}

	return ReverseString(str[1:]) + string(str[0])
}

func main() {

	originalStr := "josias carneiro"
	reversedStr := ReverseString(originalStr)
	io.Println(io.Sprintf("%v reversed = %v", originalStr, reversedStr))
}
