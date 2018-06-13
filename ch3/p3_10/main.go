//练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

package main

import (
	"bytes"
	"fmt"
)

func main() {
	// fmt.Println(comma("12"))
	// fmt.Println(comma("123"))
	// fmt.Println(comma("1234"))
	// fmt.Println(comma("1234567"))

	fmt.Println(comma2("1"))
	fmt.Println(comma2("12"))
	fmt.Println(comma2("123"))
	fmt.Println(comma2("1234"))
	fmt.Println(comma2("1234567"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {

	var buf bytes.Buffer

	j := 0

	for i := (len(s) - 1); i >= 0; i-- {
		buf.WriteByte(s[i])
		j++
		if j == 3 && len(s)%3 != 0 {
			buf.WriteByte(',')
			j = 0
		}
	}

	s1 := buf.String()

	var b []byte

	for i := (len(s1) - 1); i >= 0; i-- {
		b = append(b, s1[i])
	}

	return string(b)

}
