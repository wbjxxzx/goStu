package main

import "fmt"

func main() {
	s := []string{"a", "b", "b", "c", "", "c"}
	fmt.Println(unique(s))
}
func unique(s []string) []string {
	if len(s) == 0 {
		return s
	}
	e := s[0]
	ret := s[:0]
	ret = append(ret, e)
	for _, v := range s[1:] {
		if v == e {
			continue
		} else {
			ret = append(ret, v)
			e = v
		}
	}
	return ret
}
