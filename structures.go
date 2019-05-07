package main

import "fmt"

// Structure
type rect struct{
	height int
	width int
}

// rect's function
func (r rect) area() int{
	return r.height * r.width
}

func main(){
	p := rect{10, 22}
	fmt.Println(p.height, p.width)
	fmt.Println(p.area())
}