package main

import "fmt"

const (
	a = 1 << iota
	b = 1 << iota
	c = 1 << iota
)

func main()  {
	//2 4 8
	fmt.Println(a,b,c) //2进制 移动位数
}