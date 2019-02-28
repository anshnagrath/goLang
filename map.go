package  main

import "fmt"

func main(){
	x:=map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}
	fmt.Println(x,`this is the map`)
	//to check wheather a key exist in a map
	v,ok:= x["four"]
	fmt.Println(v)
	fmt.Println(ok)
	if v,ok:=x["four"]; ok{
		fmt.Println(`four exist in the map`,v)
	}else {
		fmt.Println(`four does not exist`)

	}
	if v,ok:=x[`three`];ok{
		fmt.Println(`three exists`,v)
	}
}