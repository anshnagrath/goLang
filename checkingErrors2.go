package main
import (
		"fmt"
		"os"
		"strings"
		"io"
		)
func main(){
	f,err:= os.Create("name.txt")
	if err != nil{
		fmt.Println(err)
		return 
	}
	defer f.Close()
	r:= strings.NewReader("James Bond")
	io.Copy(f,r)
}		