---
title: spring入门教程
mathjax: false
date: 2023-03-16 19:12:03
tags: 教程
categories:
---



# 知道你在学什么

额,一开始我就把spring和springboot两个搞混淆了.

我们现在学习就是直接去学习springboot就好了,配置相对较为简单不是吗?

其他的东西你是已经有经验了对吧,就不重点开始将了

# 环境搭建

讲真的,java的环境搭建应该是我目前为止,看到最为麻烦的了,其他的语言都可以通过命令行输入几条简单的命令就可以解决,但是java不行,有的时候甚至需要你自己去修改pom文件去选择你要使用的版本,我真的不是很喜欢,不过没有办法啊,你就是需要学会这些东西不是吗?

## 使用idea

假设我们现在的技术栈是 springboot +junit+ mysql + mybatis_plus + lombok

这里我们先一步一步的来并且认识一下配置文件.

### 网络不好的我们

因为网络问题 所以我们需要使用的就是阿里云的镜像服务

```
https://start.aliyun.com
```

![image-20230317133706808](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230317133706808.png)





### spring web

在创建项目的时候选择spring web即可 

这里你只需要确保就是你的maven下载完毕即可 

这里我并不想开始讲REST风格的api就算了,你可以直接运行我们application文件.



### Junit

这个其实已经和我们的spring高度绑定了不需要你自己去配置了已经.

```xml
<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-test</artifactId>
  <scope>test</scope>
</dependency>
```



### mysql + mybatis

这两个放到一起讲的原因很简单啊 一个是数据库的驱动,一个是orm框架,两个需要配合着使用.

这里存在一个mysql的版本问题,就是大家不是很想更新我们的数据库导致对应的驱动开发和适配工作也没有很好的推进,就导致你如果使用8.0以上版本的驱动很有可能就是找不到包,就会pom就会报错的结果.



```yml

```



 

# 开始解析参数

```java
@GetMapping("/hello")
//    开始解析参数 
public String hello() {
  return "hello";
}
```

两种格式 第一种

* query
* param

body又有好几种了 就是说.

这里其实是一个反思吧 就是可以看出来你其实没有很好的学习 就是模仿才是学习的第一步  而不是你自己去抓瞎 不是吗?至少你看上去可以学习到很多东西都是可以的 他还记得我的课程表 不是吗

