package main

import "fmt"

func main(){
	x:=map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,
		"five":5,
	}
	x["six"] = 6;
	if v,ok:=x["six"];ok{
		fmt.Println(`six is present`,ok,v)
	}
	//to loop over a map and to print its key value pair
	for k,s:= range x{
	fmt.Println(k,s)
	}
	delete(x,"six")
	delete(x,"seven")
	fmt.Println(x)
}