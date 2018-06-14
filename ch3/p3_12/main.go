//练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，
//也就是说它们有着相同的字符，但是对应不同的顺序。
package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(stringCompare("AFGH_+BC", "GH_AF+BC"))
}

func stringCompare(s1, s2 string) bool {
	var ss1, ss2 []string

	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		ss1 = append(ss1, string(v))
	}

	for _, v := range s2 {
		ss2 = append(ss2, string(v))
	}

	sort.Strings(ss1)

	sort.Strings(ss2)

	for i, v := range ss1 {
		if ss2[i] != v {
			return false

		}
	}

	return true
}
