package main

import "fmt"

func main(){
	x:=10
	for{
		if(x>=100) {break}
	//	fmt.Printf(" when %v is divided by 4 the remainder is %v\n",x,x%4)
		fmt.Printf(" when %v is divided by 4 the remainder is %v\n",x,x%4)
		x++
	}
}