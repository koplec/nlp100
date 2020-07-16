package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	tokens := strings.Split(msg, " ")
	pi := []int{}
	for _, t := range tokens {
		if strings.HasSuffix(t, ",") || strings.HasSuffix(t, ".") {
			pi = append(pi, len(t)-1)
		} else {
			pi = append(pi, len(t))
		}
	}
	for _, p := range pi {
		fmt.Printf("%d ", p)
	}
}
