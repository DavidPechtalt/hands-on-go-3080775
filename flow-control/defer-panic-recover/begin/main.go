// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
		defer cleanup("first")
		defer cleanup("second")

		fmt.Println("working on main...")
	// defer recovery
	defer func(){
		if err:= recover(); err != nil {
			fmt.Println(("recovered from panic"))
		}
	}()

	// panic
	panic("something bad happended")
}
