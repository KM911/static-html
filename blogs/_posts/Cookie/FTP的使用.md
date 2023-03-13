---
title: FTP的使用

date: 2022-12-12 14:17:43
tags:
categories:
---

# FTP的使用

<details><summary>什么是FTP</summary>
    FTP (File Transfer Protocol) is used to communicate and transfer files  between computers on a TCP/IP (Transmission Control Protocol/Internet  Protocol) network, aka the internet. Users, who have been granted  access, can receive and transfer files in the File Transfer Protocol  server (also known as FTP host/site).
</details>



一种文件的传输协议 我们不必探究底层的原理 只需要学会使用就好了 

## 服务器端口

ftp默认使用 21 20 两个端口 计算机网络也讲了这个

39000-40000 这个是我们PureFTPd的默认端口访问 

我们需要在服务器的防火墙开发端口

## 软件

[FileZilla](https://filezilla-project.org/) 免费软件 也不必我们去Windos控制面板里设置什么 开箱即用 简单小巧

[PureFTPd]() 这个是宝塔面板默认安装的软件 我们不必配置

## 创建ftp账户

在我们的宝塔面板中 选择ftp –>添加ftp

![image-20221021123908597](https://i0.hdslb.com/bfs/album/a877f6480cc8ab1ccb10bd2f9636ec807a009331.png)



## 连接设置

主机就是服务器的ip 

加密这里只能选择明文FTP

用户和密码是你刚刚设置的

![image-20221021123344061](https://i0.hdslb.com/bfs/album/af3a4f73bd3d90b8cb98d6fea23cb0e3dd011ec6.png)

**重点来了**

传输设置 –> 被动 

之前网络上好多教程说是主动传输 结果反而是错的 不能理解

当然了 我也不过是 一个初学者 只是解决了问题 并不懂得背后的原理 甚至很多时候 能够解问题 都是凭借运气

## 大功告成 

开始使用ftp进行文件传输吧