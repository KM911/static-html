---
title: Windows
mathjax: false
date: 2023-04-03 13:07:34
tags:
categories:
---

# Windows

不出意外的话 我日后应该只会使用我们的Win了,如果公司发mac或者可以尝试一下,但是正常情况下,应该是不会使用其他的了.



# 需要的软件列表

* pigcha
* clash for windows

# 环境变量

* nodejs
* vscode



# 软件源和安装位置修改



# NPM

```shell
npm config ls

npm config set registry https://registry.npmmirror.com/

npm config set prefix "F:\nodejs\node_global"

npm config set cache "F:\nodejs\node_cache"   

path add F:\nodejs\node_global
path add 


npm install cnpm -g for test

PNPM_HOME and add it to path
```

just need to use pnpm oK ?



# GO

只需要设置我们的goPath 和 module就OK了

```

```



## VScode

```
--extensions-dir "D:\Software\Microsoft\.vscode\extensions"
--user-data-dir D:\0.SOFT\Code\data
```





## pip

# 不要开vpn

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

