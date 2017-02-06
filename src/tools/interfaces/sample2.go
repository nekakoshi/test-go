package gg

import (
	"fmt"
)

type trimmedString string

func (t trimmedString) trim() trimmedString {
	return t[:3]
}

func main() {
	var t trimmedString = "abcdefg"
	fmt.Println(t.trim())

	var s string = string(t)

	fmt.Println(s)
}
