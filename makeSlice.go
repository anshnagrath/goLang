package main

import "fmt"

func main(){
	x:=make([]int,50,50)
	x = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	fmt.Println(len(x))
	for i:=0;i<len(x);i++{
		fmt.Println(x[i],i)
	}


}