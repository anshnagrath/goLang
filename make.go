package main

import "fmt"

func main(){
	x:=make([]int,10,10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// if we move the space more than the allocate one then the allocated size get doubles
	x = append(x,23,232 )
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}