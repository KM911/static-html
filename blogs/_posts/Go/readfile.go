package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("baidu.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	println(reader.Size())

	// 输出文件内容

}
