package gotest

import (
	"database/sql"
	"github.com/Dwarfartisan/pgears"
	_ "github.com/lib/pq"
	"testing"
)

type Data struct {
	Username   string `field:"username"`
	Departname string `field:"departname"`
	Created    string `field:"created"`
}

func Test_Exec_1(t *testing.T) {
	db, _ := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	stmt, _ := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	if _, e := stmt.Exec("", "", "2012-12-09"); e != nil { //try a unit test on function
		t.Error("函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_InsertMerge_1(t *testing.T) {
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

	if e := engine.InsertMerge(&data); e != nil { //try a unit test on function
		t.Error("函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}
