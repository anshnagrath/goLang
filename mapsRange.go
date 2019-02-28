package main

import "fmt"

func main(){
	m:=map[string][]string{
		"bond_james":[]string{"strong man","hard to get"},
		"no_dr":[]string{"evil man","icecream"},
		"missMoney":[]string{"nice laddy"},
	}
	m["test"]=[]string{"setting a new value"}
	delete(m,"no_dr")
	for k,v:=range m{
		fmt.Println(k);
		for _,j:=range v{
		fmt.Println("\t",j)
		}
	}
}