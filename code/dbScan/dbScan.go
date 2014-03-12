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
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	for i := 0; i < 10000; i++ {
		stmt.Exec("astaxie", "研发部门", "2012-12-09")
	}
	checkErr(err)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
