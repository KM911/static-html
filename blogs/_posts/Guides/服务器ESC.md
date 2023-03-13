---
title: 服务器ESC
date: 2022-10-04 19:58:31
tags: 教程
categories: 服务器
---

# 服务器的环境搭建

##  夹带私货

我推荐你使用centos 不是它很好 而是它很老 就是说 支持的版本比较长 centos7.6是2011年的 就内存占用方面领先其他的系统好多 比如ubuntu 和 debain 其实没别的理由 发展了几年 需要加些新的功能 说对吧 其实大部分的功能我们根本就不需要 :smile:

理解了为什么那么多服务器都是使用centos了 其实这个也是没有办法的事情 

> 很有意思的一点 国内服务器更多的是centos 但是全球访问来看 Ubuntu的数量是前者的2-4倍[^数据来源]

[^数据来源]:[Ubuntu VS CentOS，谁才是更好的 Linux 版本？](https://zhuanlan.zhihu.com/p/145812549 )

不出意外的话 我现在也大概率使用centos了 作为主要的服务器



# 远程连接

### 获取服务器root用户密码

![image-20221016122635966](https://i0.hdslb.com/bfs/album/71bacf6018e92d780f6aef67128086a955accac2.png)

因为我们是第一次使用服务器 所以其实没有密码 需要我们重置(设置)新的密码

### 下载 shh

[openssh 下载](https://www.mls-software.com/files/setupssh-8.2p1-1.exe?spm=5176.ecscore_overview.0.0.1b494df5dpIOCC&file=setupssh-8.2p1-1.exe)

```
ssh -V 
```

如果显示了版本信息 我们就可以开始远程连接了

```
ssh root@公网ip
```

然后会让我们输入上面的服务器密码



## 配置宝塔面板

不要自己安装任何的软件先 可能会破坏宝塔的安装环境 特别是自己安装`nginx` 或者其他的`web`相关软件

我们需要先在服务器上安装宝塔面板

```
// centos
yum install -y wget && wget -O install.sh http://download.bt.cn/install/install_6.0.sh && sh install.sh
// ubuntu/debian
wget -O install.sh http://download.bt.cn/install/install-ubuntu_6.0.sh && sudo bash install.sh
```

这里会出现 8888 端口没有开启的情况

我们需要前往安全组之类的地方进行开启端口 

### 小插曲

这里宝塔面板对于 Linux 支持版本不一样 上面更新的就是面板 6 系

### 安装网站环境

打开宝塔面板会自动进入软件安装页面

创建网站 记得要把我们的 ip 给添加进去 这样才可以用 ip 访问

## 购买域名

其实就算 没有域名 我们还是可以通过ip直接的访问 没有什么很大的影响 对于我们学习来说



### http 不安全的问题

其实并不是不安全 而是 http 没有加密 所以容易被劫持 比如钓鱼 wifi

这里我们需要安装 ssl 证书

阿里云有免费的证书申请

### 如何部署 ssl 证书

用宝塔部署要简单太多了

下载证书 解压缩 然后 将内容复制进去 哈哈哈 🆗了

# 利用服务器

买了一个服务器 有什么用呢？

我连电脑都觉得没有用 你还想让我做什么呢？

## 博客

理解博客 其实就是 HTML 文件罢了 (你这样说 什么东西不是 HTML 文件) 我的意思是 你只需提供 html 文件的更新 就好了 

## 图床

[图床搭建html](./Guides/图床搭建.html)

[图床搭建md](./Guides/图床搭建.md)

## ftp (网盘)

这个体验真的很不错我觉得 哈哈哈哈

主要是带宽限制了 我们的服务器的速度 不过 1.3MB/s 的速度其实是很不错了的 呢！！！

[FTP的使用html](./ftp的使用.html)

[FTP的使用_md_](./ftp的使用.md)

## 跑代码

这里为了可以自由地编写和运行代码 我搭建了一个jupyter 还可以远程访问

[Jupyter远程连接md](./Jupyter远程连接.md) [Jupyter远程连接html](./Jupyter远程连接.html)

我写了一个爬虫 每天的凌晨1:30分自动爬取P站的日榜 不过有一说一 P站的日榜 总是有很多奇奇怪怪的东西 (我是指比如说漫画什么的 关键是它的漫画还就一张图片 不明所以 )

![image-20221022101747695](https://i0.hdslb.com/bfs/album/cc5e3f1a23a6c0d0149b96593eeb5afc27066035.png)

![image-20221022101810021](https://i0.hdslb.com/bfs/album/5842d5b75bcece5c35b15245998420e0b6b0e440.png)



## api 接口

主要是我需要去学一点点最简单的后端内容 

其实不会很复杂的 你不可能就是说 需要很高的安全性 我们简单一点点 就好 比如说 
