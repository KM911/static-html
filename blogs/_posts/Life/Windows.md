---
title: Windows
mathjax: false
date: 2023-04-03 13:07:34
tags:
categories:
---

# Windows

不出意外的话 我日后应该只会使用我们的Win了,并且还是win10的版本,如果公司发mac或者可以尝试一下,但是正常情况下,应该是不会使用其他的了.

# 需要的软件列表 (这表还挺大的不是吗)

| 编译器 |  IDE   | 生活 | 工具 | 代码相关 | 其他 |
| :----: | :----: | :--: | :--: | :------: | :--: |
|   GO   | GOland |      |      |          |      |
|  JAVA  |  Idea  |      |      |          |      |
|  NODE  | vscode |      |      |          |      |
|   C    |        |      |      |          |      |
|   R    |        |      |      |          |      |
| Python |        |      |      |          |      |
|  RUST  |        |      |      |          |      |
|        |        |      |      |          |      |
|        |        |      |      |          |      |

# 软件源和安装位置修改

我之前其实尝试修改过软件的默认安装路径,其实这样不好,像 `edge` 就会因为这个原因导致无法安装.



## NPM

主要是换源和修改全局npm包的下载位置

```shell
npm config ls
npm config set registry https://registry.npmmirror.com/
npm config set prefix "D:\PATH\nodejs\node_modules\global"
npm config set cache "D:\PATH\nodejs\node_modules\cache"   
// npm install cnpm -g for test
```

环境变量修改部分

```
path add F:\nodejs\node_global
path add PNPM_HOME 这里我们需要进行一个测试
```

## GO

只需要设置我们的goPath 和 module就OK了,当然了还需要换源.

`GOPATH` 理论上是我们项目的地址,可以有多个,但是可以存在一个唯一的就是说,代表着全局仓库.

%GOPATH%\bin 需要将其添加到环境变量中,这里就是为了让我们下载的包可以直接运行是吧

其实我还是不理解,是否可以就是  `GOROOT` 和 `GOPATH` ,因为他们两个的项目结构是高度相似的.

编译的版本可以新一些,其实没有什么问题,作为一个工具就是主打的一个叫做更新.

我们还需要

## VScode

主要解决的问题是,用户数据和插件安装问题,在启动方式上添加如下的参数就好了,可能位置需要调整一下,但是大体是一样的.

```
--extensions-dir "D:\SOFT\IDEA\VsCode\extensions"
--user-data-dir "D:\SOFT\IDEA\VsCode\data"
```

## pip


这里很好笑,你开启vpn后pip将会无法下载

超时默认事件是100秒 如果你超市了 可以去看看 特别是大的东西

换源 创建pip.ini文件 

这里我们可以使用就是xcopy命令

```shell
[global]
index-url = https://pypi.tuna.tsinghua.edu.cn/simple
```

超时设置

```
--timeout=100 
```

修改默认安装路径 修改lib\site.py

```python
USER_SITE = "D:/SOFT/Path/Python/Lib/site-packages"
USER_BASE = "D:/SOFT/Path/Python/Scripts"
```

## JAVA 

我们直接利用IDEA就好了让他帮我们下一个JDK不需要管太多的事情.

`MAVEN_HOME` maven的位置 maven是一个工具

`M2_HOME` maven的安装位置

我直接把两位位置设置成为一样的了 哈哈哈哈.

换



## Docker 

```
wsl 
```

导出wsl子系统镜像:

```
wsl --export docker-desktop docker-desktop.tar
wsl --export docker-desktop-data docker-desktop-data.tar
```

删除现有的wsl子系统：

```
wsl --unregister docker-desktop
wsl --unregister docker-desktop-data
```

重新创建wsl子系统：

```
wsl --import docker-desktop d:\your-install-path docker-desktop.tar
wsl --import docker-desktop-data d:\your-install-path docker-desktop-data.tar
```

##  WSL

默认会下载就是2

```
wsl --install 
```



```
wsl --export Ubuntu-20.04 d:\wsl-ubuntu20.04.tar

wsl --unregister Ubuntu-20.04

wsl --import Ubuntu-20.04 D:\SOFT\COMPUTER\Docker\wsl d:\wsl-ubuntu20.04.tar --version 2
```





## GIT

```
git config --global http.proxy http://127.0.0.1:7890
```


## Rust

首先创建两个环境变量 ,然后就是安装rust setup mscv

```R
RUSTUP_HOME:存储工具链和配置文件
CARGO_HOME:存储cargo的缓存
```

