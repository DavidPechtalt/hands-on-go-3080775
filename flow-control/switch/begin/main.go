// flow-control/switch/begin/main.go
package main

import (
	"fmt"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// os := runtime.GOOS

	// refactor with switch statement
	// if os == "linux" || os == "darwin" || os == "unix" {
	// 	fmt.Println("*nix variant")
	// } else if os == "windows" {
	// 	fmt.Println("Windows")
	// } else {
	// 	fmt.Printf("%s.\n", os)
	// }
	spew.Dump(runtime.GOOS)

	switch os := runtime.GOOS; os{
	case "linux", "darwin", "unix":
		fmt.Println("*nix variant")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s.\n", os)
	}
}
