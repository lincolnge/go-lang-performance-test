package main

import (
	"fmt"
	"github.com/Dwarfartisan/pgears"
	"os"
	"runtime/pprof" // 引用pprof package
)

type Data struct {
	Username   string `field:"username"`
	Departname string `field:"departname"`
	Created    string `field:"created"`
}

func main() {
	var reg Data
	f, _ := os.Create("profilefile.prof")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile
	engine, err := pgears.CreateEngine("postgres://postgres@127.0.0.1/postgres?sslmode=disable")
	checkErr(err)
	engine.MapStructTo(&reg, "userinfo")

	username := "astaxie"
	var departname = "研发部门"
	var created = "2012-12-09"

	var data Data = Data{}
	data.Username = username
	data.Departname = departname
	data.Created = created

	for i := 0; i < 10000; i++ {
		err = engine.InsertMerge(&data)
		checkErr(err)
	}

	fmt.Println(data)
}

func checkErr(err error) {
	if err != nil {
		// panic(err)
	}
}
