package main

import (
	"fmt"
	"math"
	"os"
	"runtime/pprof"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	rectarea := 0.00
	for i := 0; i < 10000000; i++ {
		rectarea = r.height * r.height
	}
	return rectarea
}

func (c Circle) area() float64 {
	cirarea := 0.00
	for i := 0; i < 10000000; i++ {
		cirarea = c.radius * c.radius * math.Pi
	}
	return cirarea
}

func main() {
	f, _ := os.Create("proe.prof")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
