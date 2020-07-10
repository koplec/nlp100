package main

import "fmt"

func main(){
	s := "stressed"
	l := len(s)
	//fmt.Printf("len->%d\n", l)
	ret := ""
	for i:=l-1; i>=0; i--{
		//fmt.Printf("i: %d\n", i)
		ret += string(s[i])
	}
	fmt.Printf("reversed: %s", ret)
}