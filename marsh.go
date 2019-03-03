package main
import (
	"fmt"
	"encoding/json"
)
type  devLife struct{
	Code string
	Eat string
	Sleep string
	Repeat string
}
func main(){
s:=devLife{
	Code:"eachday",
	Eat:"less",
	Sleep:"more",
	Repeat:"twice a day",
}
g:=devLife{
	Code:"in all languages",
	Eat:"allkind of foods",
	Sleep:"adiquate",
	Repeat:"thrice a day",
}
sl:=[]devLife{s,g}
fmt.Println(sl)
x,err:=json.Marshal(sl)
if err !=nil{
	fmt.Println(err)

}
fmt.Println(string(x))
}
