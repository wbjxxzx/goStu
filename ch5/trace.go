package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	log.Printf("in bigSlowOperation")
	time.Sleep(5 * time.Second)

}
func trace(s string) func() {
	start := time.Now()
	log.Printf("enter %s", s)
	return func() {
		log.Printf("exit %s (%s)", s, time.Since(start))
	}
}
func main() {
	bigSlowOperation()
}
