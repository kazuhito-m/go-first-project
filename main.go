package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kazuhito-m/go-first-project/calc"
)

func main() {
	fmt.Println("Hello World!")
	// 文字列json化の例
	json := simplejson.New()
	json.Set("message", "Hello, World!")
	b, _ := json.EncodePretty()
	fmt.Printf("%s\n", b)
	// 計算の例
	const v1, v2 = 1, 2
	fmt.Printf("Add(%v, %v) = %v\n", v1, v2, calc.Add(v1, v2))
	fmt.Printf("Sub(%v, %v) = %v\n", v2, v1, calc.Sub(v2, v1))

}
