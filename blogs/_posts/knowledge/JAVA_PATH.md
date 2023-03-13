# JAVA环境配置

首先 你需要下载并安装 [JDK](https://www.oracle.com/java/technologies/downloads/) (Java SE Development)  

大概的名字是`jdk-17_windows-x64` 之类的 不要下载运行时`runtime`

请牢记你的安装路径

## 配置JAVA_HOME  和 CLASSPATH

你需要记住之前的安装路径 这很正常 然后添加环境变量

新建一个变量环境变量 名字为JAVA_HOME 内容为 上面的路径 `SOFT\JAVA`

新建一个环境变量 名字为CLASSPATH 内容为 `.;%JAVA_HOME%\lib;%JAVA_HOME%\lib\tools.jar`

然后添加PATH

新建 添加如下的路径

```
%JAVA_HOME%\bin
%JAVA_HOME%\jre\bin
```

## 报错

如果有`java` 但是没有`javac` 很可能是没有安装正确的`JDK`

 ![image-20221104144002765](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20221104144002765.png)

删除我们的 `lib/ext` 文件夹就好了  

有的时候`win`会无法识别`%`里的内容 建议将上面的 `%JAVA_HOME` 写成绝对路径