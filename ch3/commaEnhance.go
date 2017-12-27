package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("12345678"))
	fmt.Println(comma("-12345"))
	fmt.Println(comma("+1345."))
	fmt.Println(comma("+1345.13"))
}
func comma(s string) string {
	var buf bytes.Buffer
	flag := ""
	if strings.HasPrefix(s, "-") {
		flag = "-"
		s = s[1:]
	}
	if strings.HasPrefix(s, "+") {
		flag = "+"
		s = s[1:]
	}
	buf.WriteString(flag)
	dot := strings.Index(s, ".")
	if dot == -1 {
		dot = len(s)
	}
	n := len(s[:dot])
	mod := n % 3
	buf.WriteString(s[:dot][:mod])
	for i := mod; i < n; i++ {
		if i > 0 && i%3 == mod {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[:dot][i])
	}
	if dot != len(s) {
		buf.WriteString(s[dot:])
	}
	return buf.String()
}
