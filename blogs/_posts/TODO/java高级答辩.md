---
title: java高级答辩
mathjax: false
date: 2023-04-01 13:02:43
tags:
categories:
---

# JAVA高级答辩

我们全部的学习内容为

1. 流
2. 文件
3. XML
4. 网络编程
5. JDBC
6. 国际化 *没有体现*  可以移除就是说 
7. 注解
8. 安全 *没有体现* 

## 通过XML进行文件配置

xml解析器

```java
public class XMLParer {
  private Element root;
  private NodeList nodeList;
  
  public XMLParer(String __FILE__) {
    try {
      DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
      DocumentBuilder builder = factory.newDocumentBuilder();
      // 读取XML文件
      Document doc = builder.parse(__FILE__);
      this.root = doc.getDocumentElement();
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
  
  public void setNodeList(String __TAG__) {
    this.nodeList = this.root.getElementsByTagName(__TAG__);
  }
  
  public String getNodeValue(String __TAG__) {
    if (this.nodeList == null) {
      System.out.println("Error: nodeList is null");
      return null;
    }
    return ((Element) this.nodeList.item(0)).getElementsByTagName(__TAG__).item(0).getTextContent();
  }
}

```

##  Http服务器

### 亮点

* 线程池进行内网ip查找 提高查找效率

* 利用注解+反射实现了服务的自动注册

* mysql软删除 提高IO性能

* RespondHeaderFactory 工厂类

  

  













# String的源码

从接口 `Constable` 和 `Stable`注解就可以知道, `String`底层是利用一个 `byte`数组去存储数据的,并且因为数组的长度是固定的,理论上我们对字符串进行一点的修改,都会创建一个新的 `String`对象.

很多语言都是这样的,这样不就会造成

```java
public final class String implements Serializable, Comparable<String>, CharSequence, Constable, ConstantDesc {
  @Stable
  private final byte[] value;
  ......
}
```

## 为什么不能利用 == 进行字符串的判断

我们先看一段代码,这里的结果两次都是true

```java
int a = 10;
int b = 10;
System.out.println(a == b);
//
Integer c = 10;
Integer d = 10;
System.out.println(c == d);
```

我们将其的值改为一个不是那么常见的数,我们再看看结果.这里的结果就不一样了,第一次的判断还是true,第二次的就变成false了.

```java
int a = 10134;
int b = 10134;
System.out.println(a == b);

Integer c = 10134;
Integer d = 10134;
System.out.println(c == d);
```

`java ==` 对普通的数据类型时,是比较两者的值,但是对于对象,就是比较其地址,判断其实是否是同一个对象. 

这里是因为 `java` 将常见的数字提前构建了一个对象,当你 写下 `Integer c = 10` 不会重新创建一个对象,而是将提前创建好的对象的引用赋值给 `c` ,这里就出现了一种歧义性,两个值相同的对象,利用 `==`判断的结果不一定相同.

如果你想判断两个 `Interger` 是否是在数学上相等,你需要使用 `equals` 方法去比较他们的值.

```java
String str1 = "a";
String str2 = new String("a");
System.out.println(str1 == str2);
System.out.println(str1.equals(str2));
```

结论就是,java中对对象进行相等判断,都应该使用 `equals` 方法,





# 疑问

## 使用事务

* 使用场景?

数据库操作失败的原因? sql语句有问题? 还是?

对于单条sql语句有必要开启事务吗

特别是查询语句,根本就不存在“失败”.





















## ACID

### Atomicity 原子性

原子性是指事务是一个不可再分割的工作单位，事务中的操作要么都发生，要么都不发生。





### Consistency 一致性





### Isolation 隔离性

多个事务互不影响.









### Consistency 一致性

一致性是指在事务开始之前和事务结束以后，数据库的完整性约束没有被破坏。这是说数据库事务不能破坏关系数据的完整性以及业务逻辑上的一致性。









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

```js
1、早期的前端数据提交都是http://hostName:port/path?key1=value1&key2=value2，如果是这种key=value格式的，则在controller里面只需要声明形参即可，不需要加任何注解。如果是post请求，请求体是key=value这种格式的，其实接收方法是一样的。与此同时这种格式在请求头里面体现为：content-type=application/x-www-form-urlencoded。
2、后来人们提出了json格式，如果请求为post且格式为json，则你声明的形参为entity或map时需要添加@RequstBody注解。与此同时这种格式在请求头里面体现为：content-type=application/json。
3、后来人们又提出了restful风格的api，他的格式为http://hostName:port/path/value1/value2，此时你声明的形参需要添加@PathVariable。
现在主流的参数提交就这几种格式，希望能够帮助大家不要混淆。
```

