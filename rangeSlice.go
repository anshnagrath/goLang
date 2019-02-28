package main

import "fmt"

func main(){
	x:= []int{1,2,3,4,6}
	for k,v:= range x{
		fmt.Println(k,v)
	}


}
