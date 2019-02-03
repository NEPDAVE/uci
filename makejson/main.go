package main

import (
	"encoding/json"
	"fmt"
)

var (
	name    string
	address string
)

func main() {
	m := make(map[string]string)

	//getting name data
	fmt.Printf("enter a name: ")
	_, err := fmt.Scan(&name)

	if err != nil {
		fmt.Println(err)
	}

	m["name"] = name

	//getting address data
	fmt.Printf("enter an address: ")
	_, err = fmt.Scan(&address)

	if err != nil {
		fmt.Println(err)
	}

	m["address"] = address

	//marshaling data into json
	barr, err := json.Marshal(m)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(barr))
}
