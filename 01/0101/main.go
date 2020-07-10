package main

import "fmt"

func main() {
	str := "パタトクカシーー"
	runes := []rune(str)
	toRunes := []rune{runes[0], runes[2], runes[4], runes[6]}
	fmt.Printf("%s\n", string(toRunes))
}
