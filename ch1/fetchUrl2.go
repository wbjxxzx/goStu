package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		resp, err := http.Get(addPrefixHTTP(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("resp status: ", resp.Status)
		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%7d %s\n", nbytes, url)
	}
	end := time.Now()
	fmt.Println("elapsed time: ", end.Sub(start))
}
func addPrefixHTTP(s string) string {
	if strings.HasPrefix(s, "http://") {
		return s
	}
	return "http://" + s
}
