package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var pattern string
	var concatPattern string
	splitPattern := make(map[string]int)

	flag.StringVar(&pattern, "pattern", "", "Pattern")
	flag.Parse()

	for _, symbol := range pattern {
		if _, ok := splitPattern[string(symbol)]; ok {
			splitPattern[string(symbol)] = splitPattern[string(symbol)] + 1
		} else {
			splitPattern[string(symbol)] = 1
		}
	}

	for key, value := range splitPattern {
		concatPattern = concatPattern + key + strconv.Itoa(value)
	}

	fmt.Println(concatPattern)
}
