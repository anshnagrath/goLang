package main

import "fmt"

type person struct{
	first string
	last string
	age int
}
type secretagent struct {
	person
	ltk bool
}
func main(){
sa1:= secretagent{
	person:person{
		first:"james",
		last:"bond",
		age:32,
	},
	ltk:true,
}
	p2:= person{
	first:"Miss",
	last:"Monney",
	age:24,
	}
	fmt.Println(sa1);
	fmt.Println(p2);
	fmt.Println(sa1.first,sa1.last);
	fmt.Println(p2.first)

}