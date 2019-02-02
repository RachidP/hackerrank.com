package main

import "testing"
import "fmt"

var fibTest =[]struct{
	n int //input
	
	expected int //expected

}{
	{1,1},
	{2,1},
	{3,2},
	{4,3},
	{5,5},
	{6,8},
	{7,13},
	
}

func TestFib(t *testing.T){

	for _,v := range fibTest{
		actual:= fib(v.n)
		if actual != v.expected{
			t.Errorf("fib(%d): want %d, got %d" , v.n, v.expected, actual)
		}
	}
}

func ExampleFib(){
	fmt.Println(fib(6))

	// Output: 8


}

func BenchmarkFib10(b *testing.B){

	for n:=0; n< b.N; n++{
		fib(10)
	}
}
func BenchmarkFib( b *testing.B){

	for n:=0; n< b.N; n++{
		fib(n)
	}
}