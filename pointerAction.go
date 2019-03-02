package main

import "fmt"

func main(){
	x:=0
	fmt.Println("x address bf",&x)
	fmt.Println("x bf",x)
	foo(&x)
}
func foo(y *int){
fmt.Println("x after",&y)
fmt.Println("x after",*y)
*y = 43
fmt.Println("x after",&y)
fmt.Println("x after",*y)
}
