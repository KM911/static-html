---
title: Makefile简明教程
mathjax: false
date: 2023-05-06 13:10:33
tags:
categories:
---

## 前置知识

* 基本的命令行知识,了解过 `shell cmd`
* 最简单的编译命令,例如`gcc go build`.

## 编译命令

我们假设你现在有一个 `go `的项目. 你想发布一个精简版需要以下命令. (编译优化和对exe文件进行压缩)

```
go build -ldflags="-s -w" && upx -9 *.exe
```

你有的时候还需要对文件进行调试

```
go build && app.exe 
```

如果想运行测试脚本

```
go test -v ./test
```

又或者是一个 `benchmark`

```
go test -bench=. -benchmem  ./benchmark
```

你会发现你需要重复键入大量的命名 而且很多的参数挺麻烦的,这个时候有没有一种工具可以帮助你管理这些命令呢?这个时候就需要我们的 `Makefile`

## Makefile

`makefile`的基本语法如下

```
[target]:
	[command]
```

```
build :
	go build -ldflags="-s -w"
```

`target` 是你需要的目标 ,例如当我执行命令 `make build` , makefile就会检查该目标是否达成,如果没有达成,就会执行下面的命令,达成目标.

`target`可以是一个文件名,这样就会检查文件是否存在,如果不存在就可以执行命令对其进行构建.

这里有一个问题? 如果我们只是希望执行命名,而不是检查文件,我们可以提前声明. `.PHONY` 表示这些目标都是没有完成,也就是一定会执行下面的命令,并且不用担心文件名和命令重名的情况.

```makefile
.PHONY: run build realse
```



## 高级用法

### 依赖目标

有的时候我们的目标依赖于其他的目标,这个时候就可以这样做了

```
[target]: [require]
	[command]
```

### list

这里其实是为了交叉编译 不然每次都是需要我们进行修改 感觉其实挺麻烦的 



