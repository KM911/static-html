---
title: 为go应用程序添加图标
mathjax: false
date: 2023-03-21 17:33:09
tags:
categories:
---

## 环境准备

* [rsrc可执行文件](https://github.com/akavel/rsrc)
* ico图片
* 可build的go项目

## 模板

创建一个名为 `main.exe.manifest`的文件 内容如下

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<assemblyIdentity
    version="1.0.0.0"
    processorArchitecture="x86"
    name="controls"
    type="win32"
></assemblyIdentity>
<dependency>
    <dependentAssembly>
        <assemblyIdentity
            type="win32"
            name="Microsoft.Windows.Common-Controls"
            version="6.0.0.0"
            processorArchitecture="*"
            publicKeyToken="6595b64144ccf1df"
            language="*"
        ></assemblyIdentity>
    </dependentAssembly>
</dependency>
</assembly>
```

```shell
rsrc -arch amd64 -manifest main.manifest -ico main.ico -o main.syso
```

```shell
go build
```

这样就成功了 其实还是很简单的就是说 我们可以写一个模板 将其简化

