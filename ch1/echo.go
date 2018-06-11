package ch1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Echo1() {
	start := time.Now()

	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ","
		//fmt.Println(i, " "+os.Args[i])
	}
	fmt.Println("echo1: " + s)

	end := time.Since(start)
	fmt.Println("time used: ", end)
}

func Echo2() {

	start := time.Now()

	s, sep := "", ""

	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = ","
		//fmt.Println(i, " "+arg)
	}
	fmt.Println("echo2: ", s)

	end := time.Since(start)
	fmt.Println("time used: ", end)
}

func Echo3() {

	start := time.Now()

	fmt.Println("echo3: ", strings.Join(os.Args, ","))

	end := time.Since(start)
	fmt.Println("time used: ", end)
}
