package main

import (
	"fmt"
)

var scanString string
var s string

func main() {
	fmt.Printf("enter a floating point number: ")
	_, err := fmt.Scan(&scanString)
	
	if err != nil {
		fmt.Println(err)
	}
	
	for _, v := range scanString {
		if string(v) == "." {
			break	
		}
		s = s + string(v)
	}
	fmt.Println(s)
}
