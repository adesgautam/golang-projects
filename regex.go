package main

import "fmt"
import "regexp"

var p = fmt.Println

func main(){
	// Directly
	match, _ := regexp.MatchString("t([a-z])+ng", "testing")
	p(match)

	// Using Compile
	r, _ := regexp.Compile("t([a-z])+ng")
	p(r.MatchString("testing"))

	p(r.FindString("testing ting"))

	p(r.FindAllString("testing ting tng", -1))

	p(r.FindStringIndex("testing"))

	r = regexp.MustCompile("t([a-z]+)ng")
    p(r.ReplaceAllString("testing blah blah", "tested"))
}