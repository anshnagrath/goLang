package main

import "fmt"

type person struct {
	first_name string
	last_name string
}
func main(){
p1:= person{
	first_name: "James",
	last_name: "test",
}
fmt.Println(p1.first_name)
}
