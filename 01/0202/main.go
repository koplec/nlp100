package main

import "fmt"

func main() {
	a := "パトカー"
	b := "タクシー"
	ra := []rune(a)
	rb := []rune(b)
	ret := []rune{}
	for i := 0; i < len(ra); i++ {
		ret = append(ret, ra[i])
		ret = append(ret, rb[i])
	}
	fmt.Printf("ret:%s\n", string(ret))
}
