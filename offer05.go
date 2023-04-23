package main

import (
	"fmt"
	"strings"
)

// offer05不能够通过ac但我打印的很奇怪
func replaceSpace(s string) string {
	strings.Replace(s, " ", "%20", -1)
	return s
}
func main() {
	s := "wo w o"
	s = strings.Replace(s, " ", "%20", -1)
	fmt.Printf(s)
}
