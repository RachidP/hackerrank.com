package main

import "fmt"



func main() {
	fmt.Println("fib di 6 = ", fib(6))
	
}

func fib(n int) int{
	if n<2{
		return n
	}
	
	return fib(n-1) + fib(n-2)

}

