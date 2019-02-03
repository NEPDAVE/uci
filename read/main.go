package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	fileName string
	names    []person
)

func main() {
	fmt.Printf("enter a filename: ")
	_, err := fmt.Scan(&fileName)

	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		appendName(scanner.Text())
	}

	fmt.Println(names)
}

type person struct {
	fname string
	lname string
}

func appendName(name string) {
	nameSlice := strings.Split(name, " ")
	p := person{
		fname: nameSlice[0],
		lname: nameSlice[1],
	}

	names = append(names, p)
}
