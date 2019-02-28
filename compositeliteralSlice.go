package main

import "fmt"

func main(){
	x:= [] int {42,43,44,45,46,47,48,49,50,51}
	for c,v:= range x{

		fmt.Println(c,v)

	}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
}