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

	flag.StringVar(&pattern, "pattern", "", "incoming pattern")
	flag.Parse()

	for _, symbol := range pattern {
		if _, ok := splitPattern[string(symbol)]; ok {
			splitPattern[string(symbol)] = splitPattern[string(symbol)] + 1
		} else {
			splitPattern[string(symbol)] = 1
		}
	}

	for symbol, symbolCount := range splitPattern {
		concatPattern = concatPattern + symbol + strconv.Itoa(symbolCount)
	}

	fmt.Println(concatPattern)
}
