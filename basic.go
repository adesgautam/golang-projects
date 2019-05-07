package main

import "fmt"
import "math"

const x = 10

func main(){
	fmt.Println("Hello Adesh")

	// math
	a := 1
	b := 2
	fmt.Println(a,b, math.Sin(x))
	
	// for loop
	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}

	// if statement
	if num := 9; num > 0 {
        fmt.Println(num, "is negative")
	}

	// switch statement
	i := 2
	switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    // Arrays
    var x [5]int
    y := [5]int{1,2,3,4,5}
    // 2-D array
    var two[2][3]int
    fmt.Println("dcl:", x)
    fmt.Println("dcl:", y)
    for i:=0;i<2;i++{
    	for j:=0;j<3;j++{
    		fmt.Print(two[i][j])
    	}
    }

    // Strings
    s := "hey hey hey"
    fmt.Println("s:", s[1:5])

    // Dictionary
    dict := make(map[string]int)
    dict["a"] = 1
    //in, t := dict["a"]
    fmt.Println(dict["a"])

    // Range
    nums := []int{1,2,3,4}
    for in, num := range nums{
    	fmt.Println(in, num)
    }
    rdict := map[int]string{1:"a", 2:"b"}
    for k,v := range rdict{
    	fmt.Println(k,v)
    }
}