package main

import "fmt"

const SIZE = 10

func main() {
	a := [SIZE]int{1, 2, 3, 4, 5, 6}
	reverse(&a)
	fmt.Println(a)
}
func reverse(a *[SIZE]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
