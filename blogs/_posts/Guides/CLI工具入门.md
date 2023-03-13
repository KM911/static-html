---
title: CLI工具入门

date: 2023-01-12 14:17:43
tags:
categories:
---

# CLI工具入门

随着我对于键盘的喜爱程度不断上升— 希望更多的操作都通过我们的键盘去完成 对于CLI工具的要求自然就有了.

我觉得只要可以解决你的问题,图形化还是命令行都是很好的工具 qaq

# CLI 框架

```go
go get github.com/urfave/cli/v2

import ("github.com/urfave/cli/v2")
```

## hello world

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/urfave/cli/v2"
)

func main() {
    app := &cli.App{
        Name:  "boom",
        Usage: "make an explosive entrance",
        Action: func(*cli.Context) error {
            fmt.Println("boom! I say!")
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

```

## 基本概念

FLAG  –flag 

Argument  

Command

Action



# 利用参数

```go

```



# 利用FLAG



# 结合





优先级问题

```
gt argv --flag

gt --flag argv
```

等一下 这里的问题在于 我们的flag类型是否会对我们的结果造成影响

当string flag 不输入参数时 就会视为 help

但是如果时 bool flag 就不会出现这样的问题



一句话 flag 始终应该在参数的前面 

然后根据flag是否需要参数 也就是我们的bool类型不需要参数 向后增加 就是这样的一个逻辑

如果执行了 subcommand 的action 就不会执行主程序的action 只会执行一个action 如果

如果我们想在build 后运行 程序 很明显 我们应该把我们的逻辑放到我们的就是subcommand中去

flag 在我们的 subcommand 前 

flag是最高的优先级 





hexo -s 

hexo -w 这里是一个布尔类型的 其实是 







