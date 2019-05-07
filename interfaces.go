package main

import "fmt"
import "math"

// Interface
type geometry interface{
	area() float64
	perimeter() float64
}

// Structures
type square struct{
	side float64
}

type circle struct{
	radius float64
}

// Functions of Structures
func (sq square) area() float64 {
	return sq.side*sq.side
}

func (ci circle) area() float64 {
	return math.Pi * ci.radius * ci.radius
}

func (sq square) perimeter() float64 {
	return 4*sq.side
}

func (ci circle) perimeter() float64 {
	return 2*math.Pi*ci.radius
}

// Helper function for the Interface
func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main(){
	r := square{side: 20}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}