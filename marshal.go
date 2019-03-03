package main

import (
	"encoding/json"
	"fmt"
)

type person struct{
	First string
	Last string
	Age int
}
func main(){
	p1:=person{
		First: "james",
		Last:"bond",
		Age:32,
	}
	p2:=person{
		First:"ansh",
		Last:"me",
		Age:34,
	}
	z:=[]person{p1,p2}
	bs,err:=json.Marshal(z)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs));
}