package main

import "fmt"

func main(){
	switch  {
	case false:
		fmt.Println(`this should not print`)
	case (2==4):
		fmt.Println(`thisshould'nt either'`)
	case (4==4):
		fmt.Println(`this should'`)
		///to scpecify the code to run after that
		fallthrough
	case (3==3):
		fmt.Println(`this 2`)
		fallthrough
	default:
		fmt.Println("default")
		}
}