package main

import "fmt"

func main(){
	a:=42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n",a)
	//& with variable name tells us where this variable is stored
	fmt.Printf("%T\n",&a)
	//storing that address in a variable
	b:= &a
	fmt.Println(*&a,"variable address assigned to another variable")
	fmt.Println(b,"variable address assigned to another variable")
	fmt.Println(*b,"varible address derefrence and we get the value stored at that particular place")

}