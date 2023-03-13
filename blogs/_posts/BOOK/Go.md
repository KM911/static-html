---
title: go
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---


# GO的优化

## 开启编译优化

这里可以进入我们的build 

```go
go build -ldflags="-s -w" -o out.exe gofile.go
```

然后是使用[UPX](https://github.com/upx/upx) 它可以对二进制文件进行压缩 实现更加小巧

```
upx -9 out.exe
```

不显示命令行窗口

```go
go build -ldflags "-s -w -H=windowsgui"
```



## 实际测试

