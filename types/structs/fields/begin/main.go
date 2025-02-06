// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
	first, last string
}

func main() {
	// intialize author
	a := author{
		first: "Radave",
		last:  "Tote",
	}

	// print the author
	 fmt.Printf("%#v\n%q\n", a, a)
}
