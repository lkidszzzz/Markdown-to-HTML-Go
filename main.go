package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var mdname string

func main() {
	fmt.Println("请将您需要转换为HTML的Markdown文件放在此目录下")
	for {
		fmt.Print("请输入Markdown文件名： ")
		fmt.Scanln(&mdname)
		if Exists(mdname) == false {
			fmt.Println("请检查文件是否在此目录下！")
			fmt.Println("请检查文件名是否有拼写错误！")
		} else {
			fmt.Println("文件存在！")
			if (path.Ext(mdname) != ".md") && (path.Ext(mdname) != ".markdown") {
				fmt.Println("请检查文件后缀名是否为.md或.markdown！")
			} else {
				break
			}
		}
	}
	filename := strings.TrimSuffix(mdname, path.Ext(mdname))
	os.Mkdir(filename, os.ModePerm)
	htmlCopy(filename)
	cssUnzip(filename)
	htmlMake(filename, markdownToHTML(mdname))

	fmt.Println("转换成功，请输入任意字符退出。")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
