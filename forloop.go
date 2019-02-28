package main

import "fmt"

func main(){
	x:=1
	for{
		if(x>9){
			fmt.Println("break excecuted")
			break
		}
		fmt.Println("looping",x);
		x++;
	}
}
