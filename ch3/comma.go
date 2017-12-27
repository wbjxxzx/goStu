// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(comma(strconv.Itoa(123456789)))
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
