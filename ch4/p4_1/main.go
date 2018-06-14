//练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。

package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/imroc/biu"
)

func main() {

	x, y := "dfsdssx", "yfdsfsf"

	fmt.Printf("Number of different bits between %s and %s is : %d\n", x, y, GetNumberOfDiffBit(x, y))
}

//use sha256 crypting string
func Sha256Crypt(s string) []byte {

	c := sha256.Sum256([]byte(s))

	c1 := c[:]

	return c1
}

//generate binary string
func BytesToBinaryString(b []byte) (bs string) {

	s := biu.BytesToBinaryString(b)

	bs = s[1 : len(s)-1]

	//fmt.Println(bs)
	return
}

func GetNumberOfDiffBit(s1, s2 string) uint64 {

	var sum uint64

	c1 := Sha256Crypt(s1)
	c2 := Sha256Crypt(s2)

	s3 := BytesToBinaryString(c1)
	s4 := BytesToBinaryString(c2)

	for i := 0; i < len(s3); i++ {

		if s3[i] != s4[i] {
			sum++
		}
	}

	return sum
}
