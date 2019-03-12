package main
import (

		"os"
		"fmt"
		
)
func main(){
	defer foo()
	f,err:=os.Open("fatcsaal.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(f)

}
func foo(){
	fmt.Println("when os.exit is called, defered functions dont run")
}
