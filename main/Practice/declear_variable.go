package main

import "fmt"

func main(){
	var a = 1
	var b, c = 2, 3

	fmt.Println(a,b,c)

	x := 1			// var x = 1と同じ変数の宣言
	y, z := 2, 3	// var y, z = 2, 3 と同じ変数の宣言

	fmt.Println(x,y,z)
}
