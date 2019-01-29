<<<<<<< HEAD
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

=======
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

>>>>>>> 98f53c3d1c0f579d67d7ab5b49572858dd203ce6
