package main

import (
	"database/sql"
	"fmt"
)

func checkErr(err error)  {
	if err!=nil{
		panic(err)
	}
}
func main()  {
	//db, err := sql.Open("mysql", "astaxie:astaxie@/test?charset=utf8")
	db, err :=sql.Open("sqlite3","./foo.db")
	checkErr(err)

	//stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie","研发部门","2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=?where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate","id")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next(){
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect,err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
