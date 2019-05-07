package main

import "fmt"
import "time"
import "math/rand"

var p = fmt.Println

func main(){
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	p(now)
	p(secs)
	p(nanos)

	p(rand.Intn(100))
	p(rand.Float64())

	// For generating random numbers with different seed
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p(r1.Intn(100))
}