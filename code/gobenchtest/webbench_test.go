package gotest

import (
	"database/sql"
	"github.com/Dwarfartisan/pgears"
	_ "github.com/lib/pq"
	"testing"
)

func Benchmark_Exec(b *testing.B) {
	db, _ := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	stmt, _ := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	for i := 0; i < b.N; i++ { //use b.N for looping
		stmt.Exec("astaxie", "研发部门", "2012-12-09")
	}
	db.Close()
}

func Benchmark_InsertMerge(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间

	var reg Data
	engine, _ := pgears.CreateEngine("postgres://postgres@127.0.0.1/postgres?sslmode=disable")
	engine.MapStructTo(&reg, "userinfo")

	username := "astaxie"
	departname := "研发部门"
	created := "2012-12-09"

	var data Data = Data{}
	data.Username = username
	data.Departname = departname
	data.Created = created
	for i := 0; i < b.N; i++ { //use b.N for looping
		engine.InsertMerge(&data)
	}
}
