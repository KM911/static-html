---
title: java基础答辩
mathjax: false
date: 2023-03-25 17:14:30
tags:
categories:
---

# 泛型

## 解决了什么问题?

* 提高了代码复用性
* 增强了类型安全

## 具体的例子

>  学生的考试成绩,语文课的评分为 “优秀” ,”良好”,”及格”: 数学课的评分为: 0-100 ;英语课的评分为 A,B,C,D ,
>
> 假设该学生一次只能选择一门课

```java
public class Student_test {
  public static void main(String[] args) {
    // 为什么需要使用泛型 可以让我们的代码更加的灵活
    Student s1 = new Student(100);
    Student s2 = new Student("优秀");
    Student s3 = new Student('A');
    s1.ShowScore();
    s2.ShowScore();
    s3.ShowScore();
  }
}

class Student<T> {
  private String class_name;
  private T score;
  public Student(T score) {
    this.score = score;
    if (score.getClass() == Integer.class) {
      this.class_name = "数学";
    } else if (score.getClass() == String.class) {
      this.class_name = "语文";
    } else if (score.getClass() == Character.class) {
    this.class_name = "英语";
    }
  }
  public void ShowScore() {
    System.out.println("这个学生的" + this.class_name + "成绩是" + this.score);
  }
  public  T getScore() {
    return this.score;
  }
  public void setScore(T score) {
    this.score = score;
  }
  public  Object getScore() {
    return this.score;
  }
}
```

这样我们就不用为每一们课都创造一个类了,提高了代码的复用性.

类型安全

```java

```



## java泛型实现原理

简单来讲就是用Object代替泛型参数.当实例化对象时,根据不同的参数,去创建不同的对象.

我们可以将我们上面的Student类写成它原始的类型

```java
class Student{
  private String class_name;
  private Object score;
  public Student(Object score) {
    this.score = score;
    if (score.getClass() == Integer.class) {
      this.class_name = "数学";
    } else if (score.getClass() == String.class) {
      this.class_name = "语文";
    } else if (score.getClass() == Character.class) {
    this.class_name = "英语";
    }
  }
}
```

## java泛型存在的问题

* 无法使用基本的数据类型 只能使用包装类 因为是Object实现的
* 无法区分两个类型参数不同的对象.
* 无法创建数组. (数组的数据类型必须要确定)
* 静态方法无法使用泛型 

## 泛型的高阶用法

### 限制参数类型



# 反射







# 数据结构

* 知道基本数据类型的实现原理 
* 时间复杂度和空间复杂度的分析
* 在何时选取什么样的数据结构

## Benchmark 

* 在头部或者尾部插入删除 链式结构 
* 随机访问 顺序结构
* 奇怪的hashmap

1. 对比ArrayList和LinkedList在插入删除上的性能差异

   不用想因为ArrayList是通过动态数组实现的,到一定大的时候就需要重新开辟内存空间然后拷贝,性能表现就是没有用链表实现的好.
