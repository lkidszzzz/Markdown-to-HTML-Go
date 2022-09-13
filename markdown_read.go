package main

import (
	"bufio"
	"fmt"
	"github.com/88250/lute"
	"log"
	"os"
)

// 读取Markdown文件并转换为字符串形式的HTML源码
func markdownToHTML(input string) string {
	// 打开 md 文件
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	// 读取 md 内容
	var md string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		md += scanner.Text() + "\n"
	}
	if scanner.Err() != nil {
		log.Panicln(err)
	}

	luteEngine := lute.New()
	luteEngine.SetCodeSyntaxHighlight(true)
	luteEngine.SetCodeSyntaxHighlightDetectLang(true)
	luteEngine.SetCodeSyntaxHighlightLineNum(true)
	output := luteEngine.MarkdownStr("", string(md))
	fmt.Println("Markdown成功读取并转换为HTML格式！")
	return output
}
