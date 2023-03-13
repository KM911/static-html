---
title: go语言的包管理
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-11 14:17:43
tags:
categories:
---

# go语言的包管理

go mod init packagename

会在当前目录下创建go.mod 和 go.sum 文件

然后你就可以是用 go get 进行安装了 但是其实还是不是很好用不是吗



如果你想创建一个自己的包 

第一种是全局的包 任何go的项目都可以使用 还有就是 当前项目使用的包

## 全局package

我们平常待用一些 比如fmt net/http只需要写如下的语句就好了

```go
import (
	"fmt"
	"net/http"
)
```

我现在写了一个自己的包 名字叫做mylib

```go
package mylib

import (
)

func MyFunc() {
	println("Hello, world!")
}
```

如果你想直接调用它 你需要 

这里需要你打开你的go path文件夹



## 本地项目接口 







