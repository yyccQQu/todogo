package main

import (
	"strconv"
	"fmt"
)

func main()  {
	//最常见的数字转换是Atoi (string to int)和Itoa (int to string)。
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i,s)

	// ParseBool, ParseFloat, ParseInt和ParseUint将字符串转换为值:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	is, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println("将字符串转换为值",b, f, is, u)

	// parse函数返回最宽的类型(float64、int64和uint64)，
	// 但是如果size参数指定了更窄的宽度，那么结果可以转换为更窄的类型，而不会丢失数据:
	ss := "2147483647" // biggest int32
	i64, _ := strconv.ParseInt(ss, 10, 32)
	isn := int32(i64)
	fmt.Println("64->32",ss, i64, isn) //2147483647 2147483647 2147483647

	// FormatBool、FormatFloat、FormatInt和FormatUint将值转换为字符串:
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)
	fmt.Println("将值转换为字符串", s1,s2,s3,s4)

	q1 := strconv.Quote("Hello, 世界")
	q2 := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println("Quote",q1,q2)



}


