package main
import (
		"log"
		"os"
		"fmt"
		
)
func main(){
	defer foo()
	f,err:=os.Open("fatcsaal.txt")
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(f)

}
func foo(){
	fmt.Println("when os.exit is called, defered functions dont run")
}
