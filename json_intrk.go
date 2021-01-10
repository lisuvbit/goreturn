package main

import "fmt"

func isp() {
	m := map[string]int{
		"A_key": 100,
		"B_key": 300,
	}

	for key, val := range m {
		fmt.Println(key, val)
	}
}

func main() {
	isp()
}
