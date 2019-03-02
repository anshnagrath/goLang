package main

import "fmt"

func main(){
f1:= factorial(4)
fmt.Println(f1);
f2 := factorialloop(4)
fmt.Println(f2)

}

//recursion way
func factorial(n int) int{
	if n == 0 {
		return 1
	}
	return  n * factorial(n-1);
}

func factorialloop(n int)int{
	total:=0;
	for ; n>0 ;n-- {
	total *= n
	}
	return total;

}