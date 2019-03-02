package main

import "fmt"

func main()  {
	z:=[]int{1,2,3,4,7,5,6,7,8,89,43,434,34,2,34,235,23456}
	m:=  odd(sum, z...)
	fmt.Println(m)
}
func sum(x ...int) int{
	s:=0;
	for _, v:=range x {
		s+=v
	}
	return s
}
//calback
func odd(f func(x ...int) int,y ...int) int{
	var z []int
	for _,v:=range y{
	if(v%2 != 0){
	z=append(z,v)
	}

}
	fmt.Println(z,"odd number")
	return f(z...)
}