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





# String

* string 只有同一个环境空间 如果存在两个相同的字符串 就会 出现这个问题 就是 

```

```



# ADT

* dict 一般是 一对一
* map 可以是一对多
* arraylist 数据类型可变的  数组
* venctro 数据类型不可变的  数组
* set 内部数据不可重复 

```java
ArrayList：基于动态数组实现，支持快速随机访问，插入和删除元素效率较低。

LinkedList：基于双向链表实现，支持快速插入和删除元素，随机访问元素效率较低。

Vector：与ArrayList类似，但是是线程安全的，因此效率较低。

Stack：基于Vector实现，是一种后进先出的数据结构。

CopyOnWriteArrayList：基于可重入锁实现，是一种线程安全的List，支持快速随机访问，但是插入和删除元素效率较低。

```

```java
HashMap：基于哈希表实现，支持快速随机访问，插入和删除元素效率较高，但是不保证元素的顺序。

TreeMap：基于红黑树实现，支持按照元素的自然顺序或者指定的比较器顺序遍历元素。

LinkedHashMap：基于哈希表和双向链表实现，保证元素的插入顺序和访问顺序一致。

WeakHashMap：基于哈希表实现，但是键是弱引用，当键不再被引用时，会被自动从Map中删除。
```

关键是我们需要知道什么时候该使用何种数据类型 你懂吧 



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

  ## 参数传递
  
  * 获取线程运行结果
  
  线程的常用方法
  
  ![image-20230328142413350](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230328142413350.png)
  
  当然了线程的优先级并不是绝对的,也就是说就算你的优先级很高,也并一定就是一直在运行的,可以理解为就是提高中奖概率

# 泛型

其实就是在c++中模板函数 作用是一样的 我目前还不知道具体的使用场景,这里就先知道如何使用

主要有两类 分别是泛型类和泛型方法

```java
public class tem {
  public static void main(String[] args) {
    // 为什么需要使用泛型 可以让我们的代码更加的灵活 我们这里的学生分数可以是各种不同的形式
    Student math_student = new Student();
    math_student.MathTeacher(100);
    Student chinese_student = new Student();
    chinese_student.ChineseTeacher("优秀");
    Student english_student = new Student();
    english_student.EnglishTeacher("A");
  }
}

class Student<T> {
  public String Teacher;
  private T score;

  public void MathTeacher(T score) {
    this.Teacher = "Math";
    this.score = score;
  }

  public void ChineseTeacher(T score) {
    this.Teacher = "Chinese";
    this.score = score;

  }

  public void EnglishTeacher(T score) {
    this.Teacher = "English";
    this.score = score;
  }

  public Student() {
  }

  public Student(T score) {
    this.score = score;
  }

  public T getScore() {
    return score;
  }

  public void setScore(T score) {
    this.score = score;
  }

  public String toString() {
    return "Student [Teacher=" + Teacher + ", score=" + score + "]";
  }
}
```



# 反射

反射可以让你获取/修改类的私有属性和方法 这里我们还是只给出如何使用的例子

* 实例化一个对象

* 获取属性 (无论私有还是公有都是一样的)
* 调用方法 并获取其返回值 

```java
public class refla {
  public static void main(String[] args) throws InstantiationException, IllegalAccessException, NoSuchFieldException, NoSuchMethodException, InvocationTargetException {
    Class DogClass = Dog.class;
    // 实例化一个对象 这里还是会调用构造方法
    Dog newDog = (Dog) DogClass.newInstance();
    // 获取类的属性
    Field Echo_Field = DogClass.getDeclaredField("echo");
    // private属性就必须要设置为可访问的
    Echo_Field.setAccessible(true);
    // 这里就是获取某一个对象的 private 属性
    String Echo = (String) Echo_Field.get(newDog);
    System.out.println("echo is " + Echo);
    // 调用返回值为空的方法
    Method Void_Echo_Method = DogClass.getDeclaredMethod("Echo");
    Void_Echo_Method.setAccessible(true);
    Void_Echo_Method.invoke(newDog);
    // 调用返回值不为空的方法 这里需要声明返回值的类型
    Method String_Echo_Method = DogClass.getDeclaredMethod("Echo", String.class);
    String_Echo_Method.setAccessible(true);
    String echo = (String) String_Echo_Method.invoke(newDog, "wangwangwang");
    System.out.println("echo is " + echo);
  }
}
class Dog {
  private Integer age;
  private String echo = "wangwang";

  public Dog() {
    Echo();
  }

  private void Echo() {
    System.out.println(echo);
  }

  private String Echo(String echo) {
    this.echo = echo;
    System.out.println(echo);
    return this.echo;
  }

  public Dog(Integer age) {
    Echo();
    this.age = age;
  }

  public void setAge(Integer age) {
    this.age = age;
  }

  public Integer getAge() {
    return age;
  }
}

```











# lambda

学过python了 λ表达式 其实就是更加简单的一个写函数的方法.

# 为什么需要使用接口

> 一句话回答: 对于不存在继承关系,但是存在相同功能的类就需要接口.

例子A:

动物 --> 猫

动物 --> 狗

​      存在继承关系的两个类,其内部的方法可以是重载基类,也可以通过实现接口.这种时候你使用继承和接口都是一样的.

​      但是接口是无法定义数据成员的,所以无法确定这两个类存在哪些具体的属性.如果该方法是需要获取类的属性的时候,特别是共同的属性的时候,我还是觉得使用继承的方式更好.

```java
class Animal {
    // 动物的属性 比如说生活在水里,陆地上,或者是两栖
    // 属于是哪一个物种
    // ......
    public String Species;

    // 动物们都可以移动
    public void move(){}
}
class Cat extends Animal{
    Cat() {
        this.Species = "小猫";
    }
    public void move() {
        System.out.println(this.Species+"画梅花");
    }
}
class Dog extends Animal{
    Dog(){
        this.Species = "小狗";
    }
    public void move() {
        System.out.println(this.Species+"画月牙");
    }
}
```

例子B

动物 --> 狗

机器 --> 机器狗

两个不存在继承关系,但是存在相同功能的,我们就应该使用接口了.

当然了你还是可以让我们的机器和动物都继承一个基类,该基类有move这个方法.但是是否是所有的机器类都是可以移动的呢?不见得吧,这个时候使用接口就是最合理的选择.

至于你说只有一个类,并且也不会对其进行扩展,就直接定义该方法就好了.

## 题外话

很多人说接口是一种规范,定义了该方法的参数和返回值,这样的回答没有答到点子上,就是为什么用接口,因为使用继承同样可以实现上面的效果.

## 总结

- 只是为了实现一个类,直接在类中写实现就好了.
- 当两个类存在相同的属性和方法时,使用继承.
- 当两个类不存在继承关系,但是具有相同的功能,使用接口.

​       接口提供了一种更强的抽象能力,不再依靠相同的基类而存在.如果你的需求完全可以通过继承实现,就没有必要逼着自己去实现接口.



这里其实我们应该开始就是使用mybitas
