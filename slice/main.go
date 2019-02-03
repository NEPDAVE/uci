package main

import (
	"fmt"
	"sort"
	"strconv"
)

var s []int
var scan string

func main() {
	for {
		fmt.Printf("enter an integer to add it to the slice. enter 'x' to quit: ")
		_, err := fmt.Scan(&scan)

		if err != nil {
			fmt.Println(err)
		}

		if scan == string('x') || scan == string('X') {
			break
		}

		i, err := strconv.Atoi(scan)

		if err != nil {
			fmt.Println(err)
			fmt.Println("please enter an integer")
		}

		s = append(s, i)

		sort.Ints(s)
		fmt.Println(s)
	}
}
