package main

import (
	"fmt"
)

func main() {
	x:=42;

for a:=0;a<4;a++{
	fmt.Println("looping",a);
	for i:=0;i<2;i++{
		fmt.Println("inner loop",i);
	}
}
	fmt.Printf("#b%b\t#d%d\t#d%d",x,x,x)
	j:=1
	for j < 10 {
	fmt.Println(j)
	j++;
	}
	fmt.Println("done conditional loop")
}
