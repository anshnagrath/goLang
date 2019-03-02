package main

import "fmt"

func main (){
	xi:= []int{23,12,12,12,34,23,12,544,6,673,33}
	c:=sum(xi...)
	fmt.Println(c,"Total")
}

func sum(x ...int) int{
	fmt.Printf("%T\n",x)
	sum:=0;
	for _,v:=range x{
		sum+=v
	}
	return sum
}