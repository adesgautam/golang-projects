package main

import "fmt"
import "time"

var p = fmt.Println

func main(){
	now := time.Now()
	p(now)

	then := time.Date(2019, 02, 17, 20, 34, 58, 651387237, time.UTC)
	p(then.Weekday())

	p(then.After(now))

	diff := now.Sub(then)

	p(diff)

	p(then.Add(diff))
}