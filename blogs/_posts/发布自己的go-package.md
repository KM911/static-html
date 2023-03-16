---
title: 发布自己的go-package
mathjax: false
date: 2023-03-16 12:01:55
tags: 教程
categories:
---

# 理由

写代码写久了,自然就会有一些自己的东西在这里

# 流程概览

* 创建GitHub仓库 
* go mod init
* 编写package
* 上传到GitHub
* 通过tag进行发布
* go get -u url 进行使用

# 注意的细节

首先

go mod init 的地址不要带.git

最好是可以做到名称一致,仓库的名称,package的名称都是一样的,我这里就都是oslib 并且请不要使用-这个符号.比如你的GitHub仓库地址是`https://github.com/KM911/oslib.git`,你就应该执行 `go mod init github.com/KM911/oslib`

然后创建一个文件比如说oslib.go

```go
package oslib

import (
	"os"
	"path/filepath"
)

/*
将文件路径转换为go文件路径
*/
func ToGoFile(path string) string {
	return path + ".go"
}
```

然后就可以commit了,但是在push之前,你应该还需要做一件事情,就是将打上tag.

一般情况下,tag 是 vx.x.x 这样的形式,比如v0.0.1这样存在,其实就是版本号一样的存在.

# 引用

这里你就可以将remote应用了

创建了一个新的项目 ,main.go文件如下

```go
package main

import "github.com/KM911/oslib"

func main() {
   oslib.Run("dir")
}package main
```



# 如果remote更新了会发生什么结果

```go

```
