package util

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
)

var key []byte

func InitUtil(s []byte) {
	key = s
}
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// 加密
func Encrypt(text string) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

// 解密
func Decrypt(decrypted string) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func CheckErr(err error) {
	if err != nil {
		panic(&err)
	}
}

// defer func() {
// 	if err := recover(); err != nil {
// 		println("recover中扑火到了异常")
// 		println(err)
// 		println(fmt.Errorf("Error: %+v", err))
// 		fmt.Scanf("马上退出程序", "马上退出程序")
// 	}
// }()
// login()
// var a int
// db, err := sql.Open("sqlite3", "./foo.db")
// checkErr(err)
// err = CreateTable(db)
// println(&err, 222)
// //插入数据
// stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
// println(&err, 333)
// checkErr(err)

// res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
// checkErr(err)

// id, err := res.LastInsertId()
// checkErr(err)

// fmt.Println(id)

// //查询数据
// rows, err := db.Query("SELECT * FROM userinfo")
// checkErr(err)

// for rows.Next() {
// 	var uid int
// 	var username string
// 	var department string
// 	var created string
// 	err = rows.Scan(&uid, &username, &department, &created)
// 	checkErr(err)
// 	fmt.Println(uid)
// 	fmt.Println(username)
// 	fmt.Println(department)
// 	fmt.Println(created)
// }
// db.Close()
// fmt.Print("请输入")
// fmt.Scan(&a)
// fmt.Print(a)
// fmt.Scan(&a)
// 先进行MD5运算,取其中8位作为后续的运算
