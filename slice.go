package main

import "fmt"

func main(){
	//slice allows you to group togerther value of the same type
	x:=[]int{1,2,3,4}
	fmt.Println(x,"\t",len(x))
	// range will print with values
	for i, v :=range x{
		fmt.Println(i,v)
	}

	x=append(x,66,90);
	fmt.Println(x[1:len(x)],`slicing a slicesss`)
	y:=[]int{89,990,766}
	x=append(x,y...)
	fmt.Println(`appended using variadic parameter`,x)
	x=append(x[:2],x[5:]...)
	//breaking a slice
	fmt.Println(`appended using variadic parameter`,x)
}