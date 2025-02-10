// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
func whatAmI(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%v is an int ", val)
	case string:
		return fmt.Sprintf("%v is an string ", val)
	case bool:
		return fmt.Sprintf("%v is an boolean ", val)
	default:
		{
			return fmt.Sprintf("%v is not a supported type", val)
		}
	}
}
func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
