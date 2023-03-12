---
title: Jupyter远程连接
mathjax: true
date: 2022-10-18 20:09:44
tags: 教程
cover: http://81.68.91.70/api/cover
categories: 服务器
---



# 不是第一次相遇 而是第一次认识

之前在我们学习python的时候 其实就已经接触过了 jupyter了 但是但是根本就没有 真正意义上理解她 直到现在才开始认识她

下面介绍一下 如何在服务器上 搭建jupyter环境 并且可以远程访问

# 环境搭建

很有意思的一件事 – 像Ubuntu centos 都拥有python的环境 

```shell
# 升级pip3
pip3 install --upgrade pip
# 安装jupyter
pip3 install jupyter 
# 安装jupyter-c-kernel 不重要 这样只是为了 让其可以编译c文件
pip3 install jupyter-c-kernel
# 加载 c-kernel
install_c_kernel
# 查看
jupyter kernelspec list
# 启动
jupyter notebook
```

# 设置密码

```shell
ipython
from notebook.auth import passwd
# 会要求你设置密码
passwd()
// 'argon2:$argon2id$v=19$m=10240,t=10,p=8$J41dCWh0gGt27p2cY6bAbg$fUa2Ah37+pLC4efly/f/i6V4KuBe/RQakzFR1cN0ZAM'
```

 这里会得到一串密钥 保留下来

## 生成配置文件

```shell
jupyter notebook --generate-config
```

找到刚刚生成的配置文件 直接将下面的内容复制进去

我的就在`/root/.jupyter/jupyter_notebook_config.py`

```json
c.NotebookApp.ip='*'
argon2:$argon2id$v=19$m=10240,t=10,p=8$/jd3jhqfBGE1Ym7tzb3cmA$waN9ZD/onhoJnTtTXyHahuv5sitepKVjr0f5ckRajTs
c.NotebookApp.password = u'xxxxxxxxxxx（上一步生成的密钥）'
# 端口号 记得防火墙要开启该端口
c.NotebookApp.port = 5000	
c.NotebookApp.open_browser = False	#禁止自动打开浏览器
c.NotebookApp.allow_remote_access = True	#远程访问
c.NotebookApp.allow_root = True
c.NotebookApp.notebook_dir = '/www/jupyter' #设置根目录，限制访问

```

# 关闭防火墙

查看防火墙的状态

`systemctl status firewalld.service`

关闭防火墙

`systemctl stop firewalld`

查看开启的端口

`firewall-cmd --list-ports`

新增端口

`firewall-cmd --zone=public --add-port=7777/tcp --permanent`

重启防火墙

`firewall-cmd --reload`

查看端口号的使用情况

`lsof -i:port`

# 挂载我们的jupyter

通常情况下 如果直接关闭了于服务器的连接 这些还在终端运行的程序是会关闭的 为了避免这种情况 我们使用nohup 指令 

```
nohup jupyter notebook
nohup jupyter notebook >/dev/null 2>&1 &
```

好了 现在 利用 ip+端口 就可以访问的jupyter了 

![image-20221018210013405](https://i0.hdslb.com/bfs/album/1761348918a15ae56cc7af2b114c685f0a487c48.png)

等等啊 还有一种方式挂载

```shell
screen -S ll_jupyter
# 我们创建了一个窗口
jupyter notebook
// 重新进入
screen -r ll_jupyter
```

