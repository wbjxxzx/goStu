// intsToString is like fmt.Sprint(values) but adds commas.
package main

import "fmt"
import "bytes"

func main() {
	fmt.Println(intsToString([]int{1, 2, 3, 4}))
}
func intsToString(val []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range val {
		if i > 0 {
			buf.WriteString(",")

		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
