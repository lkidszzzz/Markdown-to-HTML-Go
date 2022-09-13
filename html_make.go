package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func htmlMake(filename, content string) {
	file, err := os.OpenFile(osPath()+"/"+filename+"/"+filename+".html", os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("文件打开失败: %v", err)
	}
	defer file.Close()
	// 查找文件末尾的偏移量
	n, _ := file.Seek(0, io.SeekEnd)
	// 从末尾的偏移量开始写入内容
	_, err = file.WriteAt([]byte("\n"+content+"</article>\n</body>\n</html>"), n)
	if err != nil {
		log.Fatalf("文件写入失败: %v", err)
	} else {
		fmt.Printf("HTML写入成功！")
	}
}
