package main

import (
	"fmt"

)

type person struct {
	first string
	last string
	age int
}
type agent struct{
	person
	ltk bool
}

func main(){
	sa1:=agent{
		person:person{
			first:"James",
			last:"Bond",
			age:32,
		},
		ltk:true,
	}
	sa1.speak()

}
func (a agent ) speak(){
	fmt.Println("I can speak")
}