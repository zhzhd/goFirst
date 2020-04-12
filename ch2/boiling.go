package main

import "fmt"

//包级别的声明
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling poibt = %gF or %gC\n", f, c)
}
