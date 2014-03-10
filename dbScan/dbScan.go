package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"runtime/pprof" // 引用pprof package
)

func main() {
	f, _ := os.Create("profilefile.prof")
	pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	defer pprof.StopCPUProfile() // 结束profile
	for i := 0; i < 10000; i++ {
		db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
		//插入数据
		stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
		stmt.Exec("astaxie", "研发部门", "2012-12-09")
		checkErr(err)
		db.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
