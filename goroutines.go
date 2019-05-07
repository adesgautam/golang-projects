package main

import "fmt"

func function(num int, from string){
	for i:=0;i<num;i++ {
		fmt.Println(from, i)
	}
}

func main(){
	function(5, "Normal")

	go function(5, "Goroutine")

	go func(s string){
		fmt.Println(s)
	}("Random !!!")

	fmt.Scanln()
	fmt.Println("DONE !")
}