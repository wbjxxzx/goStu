// compare tow strings have the same char, even they do not hav the same order
package main

import (
	"fmt"
)

func main() {
	fmt.Println(cmpString("abc", "bca"))
	fmt.Println(cmpString("aabbc", "baabc"))
	fmt.Println(cmpString("abcde", "abcd"))
	fmt.Println((cmpString("ab中国同同国a", "中国a国同ba同")))
}
func cmpString(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	ma := make(map[rune]int)
	mb := make(map[rune]int)
	for _, v := range a {
		ma[v]++
	}
	for _, v := range b {
		mb[v]++
	}
	for k, v := range ma {
		if bv, ok := mb[k]; !ok || bv != v {
			return false
		}
	}
	return true
}
