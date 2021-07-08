package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func read1(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 512)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 1024)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks)
}

func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd)
}

func read4(path string) string {
	fd, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fd)
}

func main() {
	bb := new([1]int)
	aa := bb[:]
	// aa := make([]int, 0)
	print(&aa)
	aa = append(aa, 2)
	print(&aa)
	aa = append(aa, 22)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	aa = append(aa, 33)
	print(&aa)
	println(aa)
}
