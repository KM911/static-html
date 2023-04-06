---
title: java高级答辩
mathjax: false
date: 2023-04-01 13:02:43
tags:
categories:
---

# 流 和 文件



# 技术栈

xml 是配置文件

jdb存储数据

telnet 进行http / UPD请求的发送

stream 进行传输



## 私有方法无法获取到注解

准确的说是`private` 的方法是无法通过 `Class.class.getMethods()` 获取注解的,一开始我还以为是因为我没有设置 `method.setAccess(true)` 经过测试就可以发现, 

当然了你肯定还是可以通过 `invoke("methodname")`的方式去调用这个方法,只是你都已经可以在编译前就知道调用哪个方法了,为什么还需要就是使用反射呢?总不能是为了去强行调用私有方法吧.

并且这里案例也告诉我们, `Class.class.getMethods()`是会获取到诸如 `toString`这样的方法,为了调用正确的方法,你需要便利该对象全部的方法,肯定是会有一定的性能损失在其中的.

```java
public class api {
  // 直接传递值 或者说是默认值
  @RUNWITHVALUE("http://www.baidu.com")
  private void ECHO(){
    System.out.println("ECHO");
  }
  
  @RUN
  public void Hello(){
    System.out.println("Hello");
  }
  
  @RUN
  public void World(){
    System.out.println("World");
  }
}

```

```java
package gen;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

public class check {
  public static void main(String[] args) throws InvocationTargetException, IllegalAccessException {
    Method[] methods = api.class.getMethods();
    for (Method method : methods) {
      // 判断是否有RUN注解
      RUNWITHVALUE a = method.getAnnotation(RUNWITHVALUE.class);
      // 也可以就是说 private
      method.setAccessible(true);
      if (a != null) {
        System.out.println("这个方法有注解" + method.getName());
        // 这里还是需要一个对象哈哈哈哈 其实不是很困难就是说 哈哈啊哈
        // 获取到注解的值
        String value = a.value();
        System.out.println(value);
        method.invoke(new api(), null);
      } else {
        System.out.println("这个方法没有注解" + method.getName());
      }
    }
  }
}

```

