package main
import ("encoding/json"
		"fmt"
		)
 type  jas struct {
	Code   string `json:"Code"`
	Eat    string `json:"Eat"`
	Sleep  string `json:"Sleep"`
	Repeat string `json:"Repeat"`
}
func main(){
	um:=`[{"Code":"eachday","Eat":"less","Sleep":"more","Repeat":"twice a day"},{"Code":"in all languages","Eat":"allkind of foods","Sleep":"adiquate","Repeat":"thrice a day"}]`
	var s []jas
	err:=json.Unmarshal([]byte(um),&s);
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(s)
}