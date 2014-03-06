package main

import (
	"fmt"
	"github.com/Dwarfartisan/pgears"
	"os"
	"runtime/pprof" // 引用pprof package
)

type Data struct {
	Id      string                 `field:"id" pk:"true" dbgen:"true"`
	Meta    map[string]interface{} `field:"meta" jsonto:"map"`
	Content map[string]interface{} `field:"content" jsonto:"map"`
}

var a string

func main() {
	var reg Data
	f, _ := os.Create("profilefile.prof")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile
	engine, err := pgears.CreateEngine("postgres://postgres@127.0.0.1/postgres?sslmode=disable")
	checkErr(err)
	engine.MapStructTo(&reg, "tt4")
	var content = map[string]interface{}{
		"type":    "chapters",
		"title":   "测试数据",
		"content": "只是测试一下不要想太多",
	}
	var meta = map[string]interface{}{
		"author":  "刘鑫",
		"local":   []string{"中国", "广东", "珠海"},
		"summary": "作者应该放到单独的字段",
	}
	var data Data = Data{}
	var data0 Data = Data{}
	data.Meta = meta
	for i := 0; i < 20; i++ {
		data.Content = content
		err = engine.InsertMerge(&data)
		checkErr(err)

		data0.Id = data.Id
		engine.Fetch(&data0)
	}
	fmt.Println(data)
	fmt.Println(data0)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
