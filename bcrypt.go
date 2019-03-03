package main
import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)
func main(){
	
	s:="12213123123"
	z,err:=bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(z,"hashed Password")

	err = bcrypt.CompareHashAndPassword(z,[]byte(s))
	if err != nil{
		fmt.Println(err)
	}
}