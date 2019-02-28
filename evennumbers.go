package main

import "fmt"

func main(){
	x:=0
	for {
		x++;
		if(x>100){
			break
		}
		if(x%2 !=0){
	//			fmt.Println("values",x)
			continue
		}
		fmt.Println("values",x)
	}

}