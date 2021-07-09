package main

import (
	"fmt"
	"strings"
)

// Prefix 为拥有一个或多个单词的链马尔可夫链的前缀。
type Prefix []string

// String returns the Prefix as a string (for use as a map key).

// String 将 Prefix 作为一个（用作映射键的）字符串返回。
func (p Prefix) String() string {
	return strings.Join(p, " ？")
}

func (p Prefix) Print() string {
	fmt.Println(p)
	return "ok!"
}

func main() {
	// p := make(Prefix, 4, 5)
	p := Prefix{"i", "love", "you"}
	fmt.Println(len(p), cap(p))
	fmt.Println("hello!")
	fmt.Println(p.String())
	p.Print()
}
