package main

import (
	"fmt"
	"sort"
)
type person struct{
	Name string
	Age int
}
func main(){
	a:=[]int{10,23,12,13,21,31,243123,14,34534,65,64,5}
	sort.Ints(a)
	fmt.Println(a)
	b:=[]string{"james","drno","rehab","test"}
	sort.Strings(b)
	fmt.Println(b)
}
func (p person) String() string{
	return fmt.Sprintf("%s:%d",p.Name,p.Age)
}
