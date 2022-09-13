package main

import (
	"fmt"
	"io"
	"os"
)

func htmlCopy(filename string) {
	// 打开源文件
	file1, err1 := os.Open(osPath() + "/templates/template.html")
	if err1 != nil {
		fmt.Println(err1)
	}

	// 创建目标文件
	file2, err2 := os.OpenFile(osPath()+"/"+filename+"/"+filename+".html", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
	}
	//使用结束关闭文件
	defer file1.Close()
	defer file2.Close()
	_, e := io.Copy(file2, file1)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("模板拷贝成功！")
	}
}
