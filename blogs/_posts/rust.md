---
layout: _posts
title: rust
date: 2023-05-14 06:34:58
tags:
---

## 好吧我还是背叛了自己

我在go中遇到了什么问题,我需要学习内存模型.

然后就是我在文件读写的时候遇到的问题

byte = u8

int = u16 i16 是

go语言中将 int to byte 是直接取余 我现在需要想想该如何调整一个

## 错误处理 

rust和`go`的错误处理逻辑有些相似 但是也不一太一样.

一句话概括就是 `rust`将函数正确的结果和错误封装到了一个 `Result`的结构体中

`unwarp` 和 `expect` 可以说是没有什么用处

不过可以减少代码 当你知道这里不会出现任何问题的时候.

```rust
let string = fs::read_to_string("hello.txt");
match string {
    Ok(string) => println!("{}", string),
    Err(error) => {
        println!("{}", error);
        // panic!("Failed to read hello.txt");
    }
}
```

## 循环语句





## 包管理

这里其实是挺麻烦的在 rust中反正我是没有看懂的

* 每个文件都是一个单独的包,没有任何的问题
* 关于 mod 和  `use` 可以理解为 一个是 import 另一个是 C++的use (用于简化作用域的名字长度)

其实说了很多 不如给一个案例

```rs
# 项目结构
main.rs
--oslib
	mod.rs
	timeclock.rs

```

```rust
# mod.rs
pub mod timeclock;
pub use timeclock::TimeClock;
```

```rust
# timeclock.rs
pub struct TimeClock {
    start: std::time::Instant,
}

impl TimeClock {
    pub fn new() -> Self {
        Self {
            start: std::time::Instant::now(),
        }
    }
}

impl Drop for TimeClock {
     fn drop(&mut self) {
        let duration = self.start.elapsed();
        println!("Time elapsed in expensive_function() is: {:?}", duration);
    }
}
```

```rust
# main.rs
mod oslib;
fn main(){
   let _clock =  oslib::TimeClock::new();
    let data = Some(1);
}
```

如果你使用了 `use` 就可以将其简化 就像这样 

这样的化就需要多写一个mod.rs文件进行包的管理,需要花费更加多的时间

```

```



## 数据类型

基本数据类型 : u8 char   

抽象数据类型 : tuple array  String Vec

rust 的slice是真的有点没让人看懂

```
5319  9387  2570  4338 04/26 946
YANHUI LIAO
3457 Kincheloe Road
Portland, 97212 

```



## 字符串处理 

* connect
* split
* index
* substr

