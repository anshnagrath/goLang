package main

import "fmt"

// _ is to not used
// << this is used toshift the bit
//iota is for increamenting the value by 1 so we have 1 then 2 then 3 then 4 and so on
const(_ = iota
kb = 1 << (iota*10)
mb = 1 << (iota*10)
gb = 1 << (iota*10)
)
func main(){
	fmt.Printf("%d\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t%b\n",mb,mb)
	fmt.Printf("%d\t\t%b\n",gb,gb)
}
