package main

import "fmt"

type person struct {
	first string
	last string
	favfood []string
}
func main() {
x := person{
	first:"james",
	last:"bond",
	favfood:[]string{"fruits","breads","chinnese"},
}
fmt.Println(x)
for k,v:=range x.favfood{
	fmt.Println(k,v)
}
}