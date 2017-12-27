package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567"))
}
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	mod := n % 3
	buf.WriteString(s[:mod])
	for i := mod; i < n; i++ {
		if i > 0 && i%3 == mod {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
