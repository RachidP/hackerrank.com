package main

import "fmt"

func main() {
	var s []int
	//s := []int{}
	fmt.Printf("%v \n %v", s, s)
	if s == nil {
		fmt.Println("is nil")
	}
	filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}
	fmt.Println(filter)
}
