---
title: join和+在文件路径处理上的对比
date: 2022-12-12 14:17:43
tags:
categories:
---

# 前言

我们在写代码的时候 经常会遇到文件读写 之类需要文件路径的情况 你可以使用`path.join` 或者 `+`对字符串进行拼接 我们就来分析一下两者的区别

## 通用性分析

在`windows`下我们`backslash`  在`Unix`下我们使用 `slash`  

![image-20221103074957329](http://81.68.91.70/picgo/7_49_57.webp)

搞笑的是 `backslsh \`还有转义的意思 为了避免被转义 需要写成 `\\` 如果在传输的过程中 对其进行了多次处理(你不知道这个`\`) 很可能变成这样的结果

很明显 如果你使用 `\`在windos下 不会出现的问题 到了服务器上就会出现一些意外情况

通用性上 `join `完胜

## 实用性分析

写代码的时候 难免会写很多错误 如果代码可以处理我们的错误就会很舒服

我们的`path.join` 拥有错误处理能力 

```js
import path from "path"
// 使用的是相对路径 
let path_test1 = path.join(process.cwd(),"./network","./test.jpeg")
// 有/ 并且过多
let path_test2 = path.join(process.cwd(),"/network","////test.webp")
// 有/ 并且是反斜杠
let path_test3 = path.join(process.cwd(),"\\network","\\test.webp")
console.log(path_test1)
console.log(path_test2)
console.log(path_test3)
```

他们都可以获得正确的结果 并且是可以不同平台间通用的 (我的是win 所以最后的路径使用的是`\`)

实用性分析 也是`join`的胜利



## 性能分析

其实不必 分析也可以知道 肯定是字符串拼接的效率要高 但是这一点点的性能差别真的没有太大优化的意义

两个性能差别挺大的 

<details><summary>分析设计</summary>
```js
import path from "path"
import time from "./api/time.js"
function timer(fun, times) {
    let start = time()
    for (let i = 0; i < times; i++) {
        fun()
    }
    let end = time()
    console.log("开始时间" + start)
    console.log("结束时间" + end)
    // console.log("耗时"+(end-start))
}
function test_join() {
    let path_new = path.join(process.cwd(), "network", "test.jpeg")
    let path_new2 = path.join(process.cwd(), "network", "\\test.jpeg")
}
function test_plus() {
    let path_new = process.cwd() + "/network" + "/test.jpeg"
    let path_new2 = process.cwd() + "/network/" + "test.jpeg"
}
timer(test_join, 10000000)
timer(test_plus, 10000000)
```
结果是:
开始时间8_36_19
结束时间8_36_27
开始时间8_36_27
结束时间8_36_27
</details>



## 总结

​		一个软件最关键的指标就是健硕性 先确保程序可以正常运行 我们再来讨论其他的 

所以以后无论什么情况下 我们都使用`path.join`来处理我们的文件路径问题

## 小技巧

`ES6`中我们无法使用`__dirname` 来获取当前的文件的路径  可以采用如下的方法

```js
let dirname = process.argv[1].reaplce(文件名,"")
```

当然了这个其实还是有问题的啊 就是如果是模块的话



#  模块间的信息传递

 导出函数 和 变量

```js
// core.js
let a = 10
let b = 20 
function add(x,y){return x+y}
export {a,b,add}
```

```js
// app.js
import * as core from "core.js"
console.log(core)
```

这样我们就可以获取到 `core` 中的函数和变量 

不过 其中的变量是常量 是无法修改的 我们传递一个对象就好了 

变量在自己的模块里是可以被修改的 到了其他的模块就会变成常量

核心模块 就是无法向其他模块传递信息的 我们只有通过函数的参数传递

获取当前的路径 

在其他的模块里也可以获得 就是

```js
let dirname = process.cwd()
```

