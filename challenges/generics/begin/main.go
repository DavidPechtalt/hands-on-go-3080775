// challenges/generics/begin/main.go
package main

import "fmt"

// Part 1: print function refactoring
func print[T string | int | bool](input T) {
	fmt.Println(input)
}

// Part 2 sum function refactoring
type numeric interface {
	~int | ~float64
}

// sum sums a slice of any type
// func sum(numbers []interface{}) interface{} {
// 	var result float64
// 	for _, n := range numbers {
// 		switch n.(type) {
// 		case int:
// 			result += float64(n.(int))
// 		case float32, float64:
// 			result += n.(float64)
// 		default:
// 			continue
// 		}
// 	}
// 	return result
// }
func sumAny[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {

	// call generic printAny function for each value above
	print("Hello")
	print(42)
	print(true)
	// call sum function
	// fmt.Println("result", sum([]interface{}{1, 2, 3}))

	// call generics sumAny function
	fmt.Println(sumAny(1, 2, 3,5.00000000000001, 3.3))
}
