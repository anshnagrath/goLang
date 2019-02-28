package main

import "fmt"

func main(){
	x:=[]string{"first","one","dimensional","slice"}
	y:=[]string{"secound","one","dimensional","slice"}
	// so it tells that we an slice inside an slice
	z:=[][]string{x,y}
	fmt.Println(z)
}