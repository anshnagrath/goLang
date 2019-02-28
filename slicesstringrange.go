package main

import "fmt"

func main(){
	xs1:= []string{"James","bond","shaken,not stired"}
	xs2:=[]string{"Miss","MoneyPeney","Hellooooooo Bond"}
	xss:=[][]string{xs1,xs2}
	for i,xs:=range xss{
		fmt.Println("record",i)
		for j, val:=range xs {
			fmt.Printf("\t index position: %v value: \t %v \n",j,val)
		}
	}
}