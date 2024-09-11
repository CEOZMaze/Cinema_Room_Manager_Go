package main

import (
	"fmt"
	"reflect"
)

// nolint: gosimple // <- DO NOT delete this comment!
func solve() interface{} {
	// Correct the data type of the variable 'a' so that the desired value fits into it
    // But you should not use unnecessarily large data types
	var a int16

	a = -1000 // nolint: gomnd // <- DO NOT delete this comment!

	return a
}
