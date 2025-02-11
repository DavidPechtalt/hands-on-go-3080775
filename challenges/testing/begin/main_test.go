// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetter(t *testing.T){
	tests := map[string]struct{
		input string
		want int
	}{
		"empty": {input: "", want:0},
		"one" : {input: "a2 32 ^ &/", want:1},
		"two": {input: "812 %6ab//"},
	}
	lc := letterCounter{}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T){
			if got := lc.count(tc.input); got != tc.want{
				t.Errorf("got %v want %v", got, tc.want)
			}
		})
	}
}
// write a test for numberCounter.count
func TestNumber(t *testing.T){
	got := numberCounter{designation: "number"}.count("12 4 5")
	want := 4
	if got != want{
		t.Errorf("Error in numberCounter.count. want %v but get %v", want, got)
	}
}
// write a test for symbolCounter.count
func TestSymbol(t *testing.T){
	got := symbolCounter{label: "number"}.count("&&")
	want := 2
	if got != want{
		t.Errorf("Error in symbolCounter.count. want %v but get %v", want, got)
	}
}
