// “Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can.”
// という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語
// は先頭の1文字，
// それ以外の単語は先頭の2文字を取り出し，
// 取り出した文字列から単語の位置（先頭から何番目の単語か）
// への連想配列（辞書型もしくはマップ型）を作成せよ

package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	words := strings.Split(msg, " ")
	//fmt.Println("words->", words)
	for idx, w := range words {
		if strings.HasSuffix(w, ".") {
			words[idx] = w[0 : len(w)-1]
		}
	}
	//fmt.Println("words->", words)
	ret := make(map[string]int)
	for idx, w := range words {
		if sliceContain([]int{0, 4, 5, 6, 7, 8, 14, 15, 18}, idx) {
			ret[string(w[0])] = idx
		} else {
			ret[string(w[0:2])] = idx
		}
	}
	fmt.Println("ret->", ret)
	//場所に応じて並び替えたい
	rev := make(map[int]string)
	for k, v := range ret {
		rev[v] = k
	}
	fmt.Println("rev->", rev)
}

func sliceContain(slc []int, target int) bool {
	for i := 0; i < len(slc); i++ {
		if target == slc[i] {
			return true
		}
	}
	return false

}
