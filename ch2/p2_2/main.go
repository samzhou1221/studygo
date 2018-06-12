//练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，
//然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/samzhou1221/studygo/ch2/p2_2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}
		ft := unitconv.Foot(t)
		m := unitconv.Meter(t)
		kg := unitconv.Kilogram(t)
		p := unitconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n", ft, unitconv.FTtoM(ft), m, unitconv.MtoFT(m))
		fmt.Printf("%s = %s, %s = %s\n", kg, unitconv.KGtoP(kg), p, unitconv.PtoKG(p))
	}
}
