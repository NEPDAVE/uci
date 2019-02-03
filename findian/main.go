package main

import (
	"fmt"
	"strings"
)

var ian string

func main() {
	fmt.Printf("Let's find ian. Enter a string: ")
	_, err := fmt.Scan(&ian)

	if err != nil {
		fmt.Println(err)
	}

	if string(ian[0]) == "i" && string(ian[len(ian)-1]) == "n" && strings.Contains(ian, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
