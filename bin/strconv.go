package main

import (
	"strconv"
	"fmt"
	"os"
	"log"
	"runtime"
	"strings"
	"errors"
)

var _log *log.Logger

func init() {
	_log = log.New(os.Stderr, "js106 ", log.LstdFlags)
}

// print object
func Print(objs ...interface{}) {
	if len(objs) == 0 {
		return
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = `N/A`
		line = -1
	}

	var fmtStrs []string
	for i := 0; i < len(objs); i++ {
		fmtStrs = append(fmtStrs, "%+v")
	}
	_log.Printf(fmt.Sprintf("%s:%d: ", file, line)+strings.Join(fmtStrs, " ∫ "), objs...)
}

// print object attr
func Printf(fmtStr string, objs ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = `N/A`
		line = -1
	}
	_log.Printf(fmt.Sprintf("%s:%d: ", file, line)+fmtStr, objs...)
}

// simple check error, and print one recent call stack
func CheckError(err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			file = `N/A`
			line = -1
		}
		msg := fmt.Sprintf("PANIC %s:%d: %s", file, line, err.Error())
		_log.Println(msg)
		panic(errors.New(msg))
	}
}

func main() {
	//最常见的数字转换是Atoi (string to int)和Itoa (int to string)。
	i, err := strconv.Atoi("-42")
	CheckError(err)
	s := strconv.Itoa(-42)
	fmt.Println(i, s)

	// ParseBool, ParseFloat, ParseInt和ParseUint将字符串转换为值:
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	is, err := strconv.ParseInt("-42", 10, 64)
	CheckError(err)
	Print(is,"isssss")

	u, ers := strconv.ParseUint("42", 10, 64)
	if ers != nil {
		panic(ers)
	}
	fmt.Println("将字符串转换为值", b, f, is, u) // true 3.1415E+00 -2a 2a

	// parse函数返回最宽的类型(float64、int64和uint64)，
	// 但是如果size参数指定了更窄的宽度，那么结果可以转换为更窄的类型，而不会丢失数据:
	ss := "2147483647" // biggest int32
	i64, _ := strconv.ParseInt(ss, 10, 32)
	isn := int32(i64)
	fmt.Println("64->32", ss, i64, isn) //2147483647 2147483647 2147483647

	// FormatBool、FormatFloat、FormatInt和FormatUint将值转换为字符串:
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)
	fmt.Println("将值转换为字符串", s1, s2, s3, s4) //true 3.1415 -42 42

	q1 := strconv.Quote("Hello, 世界")
	q2 := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println("Quote", q1, q2)

}
