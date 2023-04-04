---
title: hashmap源码
mathjax: false
date: 2023-03-31 13:03:51
tags:
categories:
---

# 阅读hashmap源码

* 环境 openjdk 19.0.2

##  自己对于hashmap的理解

之前学数据结构时,一直以为hashmap是不存储key的,通过计算key的hash值然后取余,将value存储在数组中.现在看来未必.

## Node类

从构造函数就可以看出来,node存储了key和value,以及hash,其中next是指向hash值相同的node.

存储key是为了发生hash碰撞时,可以分辨不同的Node,next是用于构造一个链表,存储hash值相同的Node.

```java
static class Node<K,V> implements Map.Entry<K,V> {
        final int hash;
        final K key;
        V value;
        Node<K,V> next;

        Node(int hash, K key, V value, Node<K,V> next) {
            this.hash = hash;
            this.key = key;
            this.value = value;
            this.next = next;
        }

        public final K getKey()        { return key; }
        public final V getValue()      { return value; }
        public final String toString() { return key + "=" + value; }

        public final int hashCode() {
            return Objects.hashCode(key) ^ Objects.hashCode(value);
        }

        public final V setValue(V newValue) {
            V oldValue = value;
            value = newValue;
            return oldValue;
        }

        public final boolean equals(Object o) {
            if (o == this)
                return true;

            return o instanceof Map.Entry<?, ?> e
                    && Objects.equals(key, e.getKey())
                    && Objects.equals(value, e.getValue());
        }
    }
```

## hash

java的hash不是对象的hash值,而是做了一个异或处理. 移动16位是为了让高位和低位都参加到

```java
static final int hash(Object key) {
  int h;
  return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16);
}
```

