package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"testaa/db"
	"testaa/services"
	"testaa/util"

	_ "github.com/mattn/go-sqlite3"
)

// //go:embed foo.db
// var f embed.FS
var key []byte

func main() {
	var ty, name, pw string

LOGIN:
	//ty, name, pw = FirstScreen()
	name = "cuiyuhui1"
	pw = "cuiyuehui"
	if strings.Trim(name, " ") == "" || strings.Trim(pw, " ") == "" {
		println("用户名或密码不能为空")
		goto LOGIN
	}

	newStr := name + "woleigeca" + pw + strings.ToUpper(name+pw)
	data := []byte(newStr)
	m := md5.Sum(data)
	key = m[8:16]
	util.InitUtil(key)
	// 根据账号查找账号对应的数据是否存在
	dbName := "./" + name + ".db"
	_, err := os.Stat(dbName)
	// if ty == "A" {
	if ty == "" {
		// 老用户登陆
		if err != nil {
			println("用户数据不存在,请重新登陆并创建用户")
			goto LOGIN
		}
		db.InitDb(dbName)
		// 查询下用户名密码是否正确

	} else {
		// 新用户登陆,先创建数据库
		if err == nil {
			println("已经有用户,请直接登陆")
			goto LOGIN
		}
		db.InitDb(dbName)
	}

	if err != nil {
		// 如果不存在就创建一下
		db.CreatDb(name, pw)
	}
	var versionService services.VersionService
	//查询数据
	aaa, _ := versionService.GetAll()
	fmt.Println(aaa)
	str := "I love this beautiful world!"
	// 对数据进行加密
	strEncrypted, err := Encrypt(str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encrypted:", strEncrypted)
	// 对数据进行解密
	strDecrypted, err := Decrypt(strEncrypted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted:", strDecrypted)

}
func Encrypt(str string) (string, error) {
	strEncrypted, err := util.Encrypt(str)
	if err != nil {
		return "", err
	}
	return strEncrypted, nil
}
func Decrypt(str string) (string, error) {
	strDecrypted, err := util.Decrypt(str)
	if err != nil {
		return "", err
	}
	return strDecrypted, nil
}

func CreateTable(DB *sql.DB) error {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()
	userTable := "CREATE TABLE IF NOT EXISTS `userinfo`(" +
		"`uid` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`username` VARCHAR(20) NOT NULL," +
		"`departname` VARCHAR(40) NOT NULL," +
		"`created` DATETIME" +
		")  ;"
	versionTable := "CREATE TABLE IF NOT EXISTS `version`(" +
		"`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`version` TEXT NOT NULL," +
		"`lastVersion` TEXT  ," +
		"`creatTime` DATETIME" +
		")  ;"
	mainTable := "CREATE TABLE IF NOT EXISTS `main`(" +
		"`id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"`title` TEXT  ," +
		"`webUrl` TEXT  ," +
		"`userName` TEXT  ," +
		"`passWord` TEXT  ," +
		"`remark` TEXT  ," +
		"`folder` TEXT  ," +
		"`updateTime` TEXT  ," +
		"`createTime` TEXT  ," +
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
	fmt.Println("CreateTable success")
	return nil
}

func OldLogin() (string, string) {
	var name, pw string
	println("请输入用户名")
	fmt.Scan(&name)
	println("请输入密码")
	fmt.Scan(&pw)
	return name, pw
}

func NewLogin() (name, pw string) {
	var pw2 string
NEW:
	println("请输入用户名")
	fmt.Scan(&name)
	if strings.Trim(name, " ") != name {
		println("用户名不能有空格")
		goto NEW
	}
	println("请输入密码")
AGAIN:
	fmt.Scan(&pw)
	println("请再次输入密码")
	fmt.Scan(&pw2)
	if strings.Trim(pw, " ") != pw {
		println("密码不能有空格")
		goto AGAIN
	}
	if pw != pw2 {
		println("两次的账号密码不同,重新输入密码")
		goto AGAIN
	}
	return name, pw
}

func FirstScreen() (ty, name, pw string) {
RESTAR:
	println("老用户登陆请输入A,新用户注册请输入B")
	fmt.Scan(&ty)
	ty = strings.ToUpper(ty)
	if ty == "A" {
		name, pw = OldLogin()
	}
	if ty == "B" {
		name, pw = NewLogin()
	}
	if ty != "A" && ty != "B" {
		println("注意输入范围,只能输入A或者B")
		goto RESTAR
	}
	return ty, name, pw
}
