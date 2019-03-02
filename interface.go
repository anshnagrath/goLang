package main

import "fmt"

type human interface {
	// other type of type is also of type human
	speak()
}
type person struct{
	first string
	last string
	age int
}
type agent struct {
	person
	ltk bool
}
func main(){
	sa1:= agent{
	person:person{
		first:"test",
		last:"revision",
		age:32,
	},
	ltk:true,
	}
	sa1.speak()
}
func (a agent) speak(){
	fmt.Println("I can speak")
}