package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func readDir() {
	dir := "C:\\Users\\Administrator\\Desktop\\100EOS5D"
	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	names := readFile()
	for i := range fileInfoList {
		// fmt.Println("i: ", i)
		os.Rename(dir+"\\"+fileInfoList[i].Name(), dir+"\\"+names[i]+".JPG")
		fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
	}
}

func readFile() []string {
	filename := "E:\\gowwwroot\\xu\\go-rgb\\hex.txt"
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	var res []string
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		res = append(res, string(a))
		// fmt.Println(string(a))
	}
	return res
}

func hex(dec uint8) string {
	return fmt.Sprintf("%02x", dec)
}
