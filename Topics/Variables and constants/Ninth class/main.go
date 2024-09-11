package main

import "fmt"

const (
	First = iota + 1 // Make iota start counting from 1 instead of 0 here!
	Second
	Third
	Fourth
	Fifth
	Sixth
	Seventh // 7
	Eighth  // something is missing after the Eighth class...
	Ninth
)

// DO NOT delete or modify the contents of the main function!
func main() {
	fmt.Println(check())
}
