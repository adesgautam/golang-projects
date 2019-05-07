package main

import "fmt"

func plus(a int,b int) int{
	return a+b
}

// Variadic function
func vfunc(nums ...int){
	fmt.Println(nums)
}

// Closure
func seq() func() int{
	a := 1
	return func() int {
		a += 5
		return a 
	}
}

func main(){
	res := plus(1, 2)
	fmt.Println(res)
	vfunc(1,2,3,4,5,6)
	val := seq()
	fmt.Println(val())
	fmt.Println(val())
}