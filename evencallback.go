package main

import "fmt"

func main(){
	d:=[]int{1,2,1,21,12,1,2332,4,34,3,544,5,4,54}
	s := even(sum, d ...);
	fmt.Println(s,"sum of even numbers");

}
func sum(x ...int)int{
	s:=0
	for _, v:=range x{
		s+=v
	}
	fmt.Println(s)
	return s
}
func even(f func(x ...int)int, x ...int) int{
	var s []int;
	for _, v:=range x{
		if v%2 == 0{
			s = append(s,v)
		}
	}
	fmt.Println(s,len(s),"sum of even numbers");
	return f(s...)
}