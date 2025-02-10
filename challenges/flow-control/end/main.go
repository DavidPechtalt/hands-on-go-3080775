package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error! ", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("Error in reading the file!")
	}
	// convert the bytes to a string
	data := string(bs)
	// initialize a map to store the counts
	analizes := make(map[string]int)
	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(data)
	// capture the length of the words slice
	analizes["words"] = len(words)
	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, char := range word {
			if unicode.IsLetter(char) {
				analizes["chars"]++
			} else if unicode.IsNumber(char) {
				analizes["numbers"]++
			} else {
				analizes["symboles"]++
			}
		}
	}
	// dump the map to the console using the spew package
	spew.Dump(analizes)
}
