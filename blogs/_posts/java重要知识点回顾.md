---
title: java重要知识点回顾
mathjax: false
date: 2023-03-18 11:48:31
tags:
categories:
---

# 对象

## 构造函数

首先我们需要知道,java中是不存在手动的内存管理机制,依赖于jvm的gc,所以java中就不存在析构函数这个概念–因为你无法决定这个对象内存空间的释放.

* 无参构造函数

主要的意义是为了帮助对象自身的属性进行初始化.这里的第二行就会报错了,原因是i没有初始化.

```java
int i ;
System.out.println(i);

int j = 10;
System.out.println(j);
```

而对于一个对象来说,即使你没有创建构造函数,编译器会自动为你生成一个构造函数进行对象属性的初始化.通常会将对象的默认值设置为null,int则为0,float为0.0.

| type         | default_value |
| ------------ | ------------- |
| object       | bull          |
| int long     | 0             |
| float double | 0.0           |
| char         | U0000         |

每一个类都会自动生成一个无参的构造函数,前提是你没有对构造函数进行重载.

如果你创建了一个有参的构造函数,但是没有编写无参构造函数,就会报错.



## this 和 super

this 集合本身 super 超集 这样其实就好理解和区分了.

和C++一样,优先执行基类或者说是父类的构造函数.(我这里默认大家就是有一定的C++语言基础的)

```java
如何用代码去体现自己的思考我看来才是最复杂的不是吗?
```

### 通过this去调用构造函数





Object类 

java中任何一个类都直接或者间接继承于Object类

## 重要的方法





#  枚举类

有限的对象

数据来源于stackoverflow 这里的人不说每个人都是技术大牛吗,但是也是正在在从事计算机相关行业的人了.可以见的,全世界范围内,我们这帮臭写代码的人,工作时间就是比较久,真的是服了,早知道还是得去读医啊,毕竟一位印度老哥说过

![image-20230319175903730](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230319175903730.png)



# 包装类 

warpper

| 基本数据类型 | 对象数据数据类型 |
| ------------ | ---------------- |
| int          | Interg           |
| boolen       | Boolen           |
|              |                  |
|              |                  |
|              |                  |

# 数据类型转换

对象数据类型 ———字符串

这里其实和我们



# 异常处理

其实你是会异常处理的,不过没有怎么见过java中的异常类型

* 





# 多线程

并发编程

## 如何使用

* 继承一个线程类 Thread

```java
public class theadertest {
    public static void main(String[] args) {
        PrintNum t = new PrintNum();
        t.start();
        // 这里是下面的hell先输出 说明不再是阻塞的了 我们成功的并发了 哈哈哈
        System.out.println("hell");
    }
}
class PrintNum extends Thread {
  // 这个重载是必要的 并且不要携带参数 这里的通信是一个问题
    @Override
    public void run ()
    {
        for (int i = 0; i < 100; i++) {
            System.out.println(i);
        }
    }
}

```



* 实现runable 接口

  ```java
  public class Runtest
  {
      public static void main(String[] args) {
          Run t = new Run();
          Thread newt = new Thread(t);
          newt.start();
          for (int i = 0; i <100 ; i++) {
              System.out.println("hell");
          }
      }
  }
  
  
  class Run implements Runnable{
      @Override
      public  void run(){
          for (int i = 0; i < 100; i++) {
              System.out.println("i = "+i);
          }
      }
  }
  
  ```

  



# 泛型

其实就是在c++中模板函数

其实我只是希望可以对于字符串的需要一个泛型,可以很轻松的将去

