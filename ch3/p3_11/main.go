//练习 3.11： 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	// fmt.Println(comma("1"))
	// fmt.Println(comma("12"))
	// fmt.Println(comma("123"))
	// fmt.Println(comma("1234"))
	// fmt.Println(comma("12345"))
	// fmt.Println(comma("123456"))
	// fmt.Println(comma("1234567"))

	// fmt.Println(comma("-1"))
	// fmt.Println(comma("-12"))
	// fmt.Println(comma("-123"))
	// fmt.Println(comma("-1234"))
	// fmt.Println(comma("-12345"))
	// fmt.Println(comma("-123456"))
	// fmt.Println(comma("-1234567"))

	// fmt.Println(comma("+1"))
	// fmt.Println(comma("+12"))
	// fmt.Println(comma("+123"))
	// fmt.Println(comma("+1234"))
	// fmt.Println(comma("+12345"))
	// fmt.Println(comma("+123456"))
	// fmt.Println(comma("+1234567"))

	fmt.Println(comma("1.1"))
	fmt.Println(comma("12.12"))
	fmt.Println(comma("123.123"))
	fmt.Println(comma("1234.1234"))
	fmt.Println(comma("12345.12345"))
	fmt.Println(comma("123456.123456"))
	fmt.Println(comma("1234567.1234567"))

	fmt.Println(comma("+1.1"))
	fmt.Println(comma("+12.12"))
	fmt.Println(comma("+123.123"))
	fmt.Println(comma("+1234.1234"))
	fmt.Println(comma("+12345.12345"))
	fmt.Println(comma("+123456.123456"))
	fmt.Println(comma("+1234567.1234567"))

	fmt.Println(comma("-1.1"))
	fmt.Println(comma("-12.12"))
	fmt.Println(comma("-123.123"))
	fmt.Println(comma("-1234.1234"))
	fmt.Println(comma("-12345.12345"))
	fmt.Println(comma("-123456.123456"))
	fmt.Println(comma("-1234567.1234567"))

}

func comma(s string) string {

	var buf bytes.Buffer
	var sign, after string

	if s[0] == '-' || s[0] == '+' {
		sign = string(s[0])
		s = s[1:]
	}

	if strings.Contains(s, ".") {
		temp := strings.Split(s, ".")

		s = temp[0]
		after = "." + temp[1]
	}

	j := 0

	for i := (len(s) - 1); i >= 0; i-- {
		buf.WriteByte(s[i])
		j++
		if j == 3 && i != 0 {
			buf.WriteByte(',')
			j = 0
		}
	}

	s1 := buf.String()

	

	var b []byte

	for i := (len(s1) - 1); i >= 0; i-- {
		b = append(b, s1[i])
	}

	return sign + string(b) + after

}
