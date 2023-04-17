---
title: 将文件拖到exe程序上,发生了什么?
mathjax: false
date: 2023-03-21 15:42:31
tags: 教程
categories: 计算机基础
---

# 如何进行推断

很简单,通过print之类的输出信息啊

错.因为你的程序会在结束后立即关闭窗口,你根本就来不及,当然了你会说 可以让窗口sleep啊,对的? 这里我用go是没有办法通过就是sleep暂停的 当然了 这里是go的问题.

```go
package main

import (
	"github.com/KM911/oslib"
	"os"
	"path"
	"runtime"
	"time"
)
func RunPath() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	return path.Dir(fullFilename)
}
func main() {
	Info := ""
	Info = Info + "cmdpath is " + oslib.CmdPath()
	Info = Info + " \nrunpath is " + RunPath()
	cwd, _ := os.Getwd()
	Info = Info + " \ncwd is " + cwd
	Info = Info + " \nos.Args is " + os.Args[0]
	Info = Info + " \nos.Args is " + os.Args[1]
	os.WriteFile("test.txt", []byte(Info), 0666)
	// 卡死
	time.Sleep(100000)
}

```

将一个其他地方的文件拖给它后,得到这样的txt文件

```txt
cmdpath is D:\SOFT\BIN 
runpath is D:/CODE/go/snap 
cwd is D:\SOFT\BIN 
os.Args is D:\CODE\go\snap\test.exe 
os.Args is D:\SOFT\BIN\video.go
```

很明显了 首先 我们的参数1 就是我们程序的绝对路径  参数2 是你拖动给我的文件 或者文件的绝对路径 

关键是 这里的cmdpath 和 runpath 和 cwd

也就是我们的拖动文件的位置 ,知道该如何进行就是一个打包了吧.



# 总结一下就是

拖动文件带来的是绝对路径,Args[0] 是接受的一方,通常是exe程序,Args[1]是被拖动的文件,可能是图片或者视频等其他任何东西.

