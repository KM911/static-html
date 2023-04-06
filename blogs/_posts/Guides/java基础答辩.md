---
title: java基础答辩
mathjax: false
date: 2023-03-25 17:14:30
tags: java
categories:
---
# 泛型

> 泛型是运行时确定函数/类的类型.

## 解决了什么问题?

* 提高了代码复用性
* 增强了类型安全

## 具体的例子

> 学生的考试成绩,语文课的评分为 “优秀” ,”良好”,”及格”: 数学课的评分为: 0-100 ;英语课的评分为 A,B,C,D ,
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

类型安全 :会在语法检查阶段发现错误

```java
public class Ser {

  public static void main(String[] args) {
      // 创建一个ArrayList去存储学生的成绩
    ArrayList Score_obj = new ArrayList();
    // add 的参数是一个对象
    Score_obj.add(99);
    Score_obj.add(78.5);
    Score_obj.add("100");
    ShowScore(Score_obj);
//    System.out.println(Score_obj.get(2).getClass());

    // 你可以利用泛型去限制我们的参数类型
    ArrayList<Integer> Score_int = new ArrayList<Integer>();
    Score_int.add(99);
//    Score_int.add("60");
//    Score_int.add(89.5);
//    存在的问题 ?  这里首先我们需要说明就是实现原理 就是说 根据其原理就可以知道有可能存在哪些问题了
//    1. 无法使用基本数据类型
//    primitive type 这里是涉及到它的实现方式的问题
    // 因为泛型其实是利用Object 去替换一个原始的类型 该类必须是对象
//     ArrayList<int> score  = new ArrayList<int>();
//    2.只有原始类型 class java.util.ArrayList
    System.out.println(Score_int.getClass());
    System.out.println(Score_obj.getClass());
    Integer a = (Integer) (10);
    if (Score_int instanceof ArrayList) {
      System.out.println((Score_int instanceof ArrayList));
    }
  }

  public static void ShowScore(ArrayList Scores) {
    for (int i = 0; i < Scores.size(); i++) {
      if ((Integer) Scores.get(i) > 60) {
        System.out.println("该学生及格");
      }
    }
  }
}
```

## java泛型实现原理 – 类型擦除

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

* 无法使用基本的数据类型
* 无法区分两个类型参数不同的对象的类.
* 无法创建数组. (数组的数据类型必须要确定)
* 静态方法无法使用泛型

## 泛型的高阶用法

* 限制类型参数的访问 super和extends
* ? 通配符

一个歧义: 同时限制类和接口时,如果该类并没有实现该接口,可以通过语法检查,但是会出现问题.

```java
public class generic_bug {
  public static void main(String[] args) {
    Cat cat = new Cat();
    Zoo zoo = new Zoo(cat);
    zoo.show();
  }
}
interface Runnable {
  void run();
}
class Animal implements Runnable {
  public void run() {
    System.out.println("Animal is running");
  }
}
interface Flyable {
  void fly();
}
// 其之类都是可以作为参数
class Cat extends Animal {
//    public void run() {
//        System.out.println("Cat is running");
//    }
}
class Zoo<T extends Animal & Flyable> {
  public T animal;

  public Zoo(T animal) {
    this.animal = animal;
  }

  public void show() {
    animal.fly();
  }

  public void run() {
    animal.run();
  }

  public Zoo() {
  }
}
```

# 反射

> 反射提供了一种方式,让你可以自由地操作对象.

因为面向对象的封装性,你无法直接调用一个对象的私有方法/获取私有属性.但是利用反射就可以在程序运行时获取该对象的全部信息.

其实反射的内容就只有这么一点.

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

# 数据结构

* 知道基本数据类型的实现方式
* 时间复杂度和空间复杂度的分析
* 在何时选取什么样的数据结构

## 分析对比

1. 对比ArrayList和LinkedList在插入删除上的性能差异

   不用想因为ArrayList是通过动态数组实现的,到一定大的时候就需要重新开辟内存空间然后拷贝,性能表现就是没有用链表实现的好.
2. Arraylist从头部插入的性能要比从尾部插入的性能低.因为从头部插入,你每次都需要将数组中的数据向后移动一位.删除同理
3. 对于没有顺序要求的数据,直接使用hashmap就好了,在绝大部分情况下,我们的hashmap是比他们要快的,尤其是在大数据量上.
4. …

## 并发安全

当使用了多线程后,普通的数据结构就不安全了.



## 用hashmap实现栈

```java
package benchmark;
import java.util.HashMap;

// 规定有如下的方法
// push 用来向栈中添加元素
// pop 返回栈顶元素 没有元素返回-1
// length 用来获取栈的长度
public class stack_hashmap {
  public static void main(String[] args)
  {
    Stack stack = new Stack();
    stack.push(1);
    stack.push(2);
    stack.push(3);
    System.out.println(stack.pop());
    System.out.println(stack.pop());
    System.out.println(stack.pop());
  }
}

class Stack {
  // 用来存储数据的hashmap
  private HashMap<Integer, Integer> map = new HashMap<>();
  private Integer length = 0;
  public int length() {
    return length;
  }
  public void push(int value) {
    length++;
    map.put(length, value);
  }
  public int pop() {
    if (length == 0) {
      return -1;
    }
    int value = map.get(length);
    map.remove(length);
    length--;
    return value;
  }
}
```

# 多线程

> 线程是操作系统的能够进行运算调度的最小单位.它被包含在进程中,是实际在运行的单位.

## 为什么需要多线程

> 提高程序的运行效率.

* 读取大文件 避免IO阻塞导致主程序一直在等待状态
* 游戏需要“同时”播放音乐和动画,还有后台发送各种请求
* 按钮组件绑定的事件 … 

## 实现多线程的方式

* 继承thread类
* 实现runable接口
* 实现callable接口 

如果需要获取到线程运行的结果,应该使用第三种方法.

## 线程的生命周期

![image-20230328200012120](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230328200012120.png)

## 线程同步

> 商品库存问题

使用关键字 synchronized 或者  ReentrantLock阻塞队列实现线程同步来实现只有一个线程在执行当前任务.

我目前比较喜欢的方式是利用synchronized

```java
package threading_demo;
public class sync_demo {
  public static void main(String[] args) {
    Thread task1 = new SyncThread();
    Thread task2 = new SyncThread();
    Thread task3 = new SyncThread();
    task1.start();
    task2.start();
    task3.start();
    // 等待线程执行完毕
    try {
      task1.join();
      task2.join();
      task3.join();
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
    System.out.println("最终剩余了" + SyncThread.products + "件商品");
  }
}
class SyncThread extends Thread {
  public static int products = 100;

  @Override
  public void run() {
    while (true) {
      synchronized (SyncThread.class) {
        if (products > 0) {
          products--;
          System.out.println("购买成功,剩余" + products + "件商品");
        } else {
          break;
        }
      }
    }
  }
}
```

## 死锁

这里两个线程都在等待对方释放锁,就陷入了卡死的地步.死锁是我们软件开发过程中,应该避免出现的情况.

多个锁嵌套.

```java
package threading_demo;
// 死锁的演示
public class loop_lock {
  public static void main(String[] args) {
    Object LockA = new Object();
    Object LockB = new Object();
    Thread_A thread_A = new Thread_A(LockA, LockB);
    Thread_B thread_B = new Thread_B(LockA, LockB);
    thread_A.start();
    thread_B.start();
  }
}
class Thread_A extends Thread {
  public Object LockA;
  public Object LockB;

  public Thread_A(Object LockA, Object LockB) {
    this.LockA = LockA;
    this.LockB = LockB;
  }

  @Override
  public void run() {
    synchronized (LockA) {
      System.out.println("Thread_A get LockA");
      try {
        Thread.sleep(1000);
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
      synchronized (LockB) {
        System.out.println("Thread_A get LockB");
      }
    }
  }
}
class Thread_B extends Thread {
  public Object LockA;
  public Object LockB;

  public Thread_B(Object LockA, Object LockB) {
    this.LockA = LockA;
    this.LockB = LockB;
  }

  @Override
  public void run() {
    synchronized (LockB) {
      System.out.println("Thread_B get LockB");
      try {
        Thread.sleep(1000);
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
      synchronized (LockA) {
        System.out.println("Thread_B get LockA");
      }
    }
  }
}
 
```
