package main

import "os"

func main() {
	var rmdirs []func()
	dirs := []string{
		"r:/a",
		"r:/b",
		"r:/c",
		"r:/1/2",
	}
	for _, d := range dirs {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
