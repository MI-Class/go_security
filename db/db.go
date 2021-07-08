package db

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"
	"testaa/util"
	"time"
)

var DB *sql.DB

func InitDb(dbName string) error {
	DB, _ = sql.Open("sqlite3", dbName)
	err := DB.Ping()
	return err
}
func CreatDb(userName string, pwd string) error {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()
	userTable := "CREATE TABLE IF NOT EXISTS `userinfo`(" +
		"`uid` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`username` VARCHAR(20) NOT NULL," +
		"`departname` VARCHAR(40) NOT NULL," +
		"`created` INTEGER" +
		")  ;"
	versionTable := "CREATE TABLE IF NOT EXISTS `version`(" +
		"`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`version` TEXT NOT NULL," +
		"`lastVersion` TEXT  ," +
		"`creatTime` INTEGER" +
		")  ;"
	mainTable := "CREATE TABLE IF NOT EXISTS `main`(" +
		"`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`title` TEXT  ," +
		"`webUrl` TEXT  ," +
		"`userName` TEXT  ," +
		"`passWord` TEXT  ," +
		"`remark` TEXT  ," +
		"`folder` TEXT  ," +
		"`updateTime` INTEGER  ," +
		"`createTime` INTEGER  ," +
		"`hash` TEXT NOT NULL" +
		")  ;"
	folderTable := "CREATE TABLE IF NOT EXISTS `folder`(" +
		"`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`name` TEXT  " +
		")  ;"
	_, err := DB.Exec(userTable)
	util.CheckErr(err)
	_, err = DB.Exec(versionTable)
	util.CheckErr(err)
	_, err = DB.Exec(mainTable)
	util.CheckErr(err)
	_, err = DB.Exec(folderTable)
	util.CheckErr(err)
	times := time.Now().Unix()
	timeStr := strconv.FormatInt(times, 10)
	timeE, _ := util.Encrypt(timeStr)
	v := md5.Sum([]byte(pwd + timeStr))
	vSrt := hex.EncodeToString(v[:])
	vSrt, _ = util.Encrypt(vSrt)
	p, _ := util.Encrypt(pwd)
	stmt, err := DB.Prepare("INSERT INTO version(version,lastVersion,creatTime) values(?,?,?)")
	_, err = stmt.Exec(vSrt, p, timeE)
	// DB.Exec("INSERT INTO version(`version`,`lastVersion`,`creatTime`) value(?,?,?)", string(v[:]), p, strconv.FormatInt(times, 10))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("user create success")
	return nil
}
