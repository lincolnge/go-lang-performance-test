package main

import (
	"fmt"
	"os"
	"runtime/pprof" // 引用pprof package
)

var a string

func main() {
	f, _ := os.Create("profilefile.prof")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile

	f1()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func f1() {
	a := "O"
	fmt.Println(a)
	f2()
	var k = 0
	for i := 0; i < 100000000; i++ {
		k = k + i
	}
}
func f2() {
	fmt.Println(a)
	var k = 0
	for i := 0; i < 100000000; i++ {
		k = k + i
	}
}
