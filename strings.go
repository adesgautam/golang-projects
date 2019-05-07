package main

import "fmt"
import s "strings"

var p = fmt.Println

func main(){
	p( s.Contains("test", "es") )
	p( s.Count("test", "t") )
	p( s.HasPrefix("test", "t") )
	p( s.HasSuffix("test", "t") )
	p( s.Index("test", "t") )
	p( s.Join([]string{"a", "b", "c"}, "") )
	p( s.Repeat("test", 2) )
	p( s.Replace("foo", "o", "0", -1) )
    p( s.Replace("foo", "o", "0", 1) )
    p( s.Split("a-b-c-d-e", "-") )
    p( s.ToLower("TEST") )
    p( s.ToUpper("test") )
}