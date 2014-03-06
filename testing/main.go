package main

import "fmt"
import "github.com/lincolnge/go-lang-performance-test/math"
import "runtime/pprof" // 引用pprof package
import "os"

func main() {
	f, _ := os.Create("profilefile.prof")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile

	for i := 0; i < 100000; i++ {
		xs := []float64{1, 2, 3, 4}
		avg := math.Average(xs)
		fmt.Println(avg)
	}
}
