package main

import "fmt"

func main(){
	x:= struct{
		first string
		last string
		age int
	}{
		first:"James",
		last:"bound",
		age:32,
	}
	fmt.Println(x,"\t",x.first)
	}
