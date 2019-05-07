package main

import "fmt"
import "encoding/json"

var p = fmt.Println

type response struct{
	name string `json:name`
	age int `json:age`
}

func main(){
	res, _ := json.Marshal(true)
	p(string(res))

	resobj := &response{name:"Adesh Gautam", age:22}
	r, _ := json.Marshal(resobj)
	p(string(r))

	str := `{"name":"Adesh", "age":22}`
	res = response{}
    json.Unmarshal([]byte(str), &res)
    p(res)

}