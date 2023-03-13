---
title: Vue
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---

# Vue

## 项目结构

>    index.html    //vue只有一个html页面
>
>   /src    // 这里是我们的主要编写内容
>
>   >   App.vue    
>   >
>   >   main.js
>   >
>   >   /assert   //静态资源 
>   >
>   >   >   css jpg 
>   >
>   >   /components  组件
>   >
>   >   >   hello.vue
>   >   >
>   >   >   todo.vue



# 理解项目

App.vue 其实也是一个组件 就好像JAVA啊 自己也是一个对象一样

`vue`文件有三个部分 和原生的是一个样的其实是

```vue
<script setup>

// 这里我们可以写一个就是更加基础的组件 就是专为了绑定数据而生的
const props = defineProps({
    msg: {
        type: String,
        default: "按钮"
    }
})

</script>

<template>
    <button>{{ msg }}</button>
</template>

<style scope>
    button{
        color:red;
    }
</style>
```

非常理想的是 组件间的类名是不会互相影响的 这对于我们的开发很友好



## JS in Vue

我们这里是用的是api形式的书写规范

```vue
<script setup> //setup 在组件被加载的时候就会执行
 //  props 是从父组件中获取到的值
const props = defineProps({
    msg: {
        type: String,
        default: "按钮"
    }
})
// 我们使用指针进行值的传递
import { ref } from 'vue'
const count = ref(0)
```

## Html in Vue

其实没有什么区别 

最外层的template 只是为了说明这里是模板 不会真的有一个`template` 

```vue
<template>
  <!-- 使此按钮生效 -->
  <BB :msg="he"></BB>
  <button @click="count++">count is: {{ count }}</button>
</template>
```

## CSS in Vue

其实没有任何的区别  `scoped` 也是区别啊 确实 

```vue
<style scoped>
   	button{
        color:red;
    }    
</style>
```



# 组件化

开始真正地书写我们的组件吧

```js
所有的 props 都遵循着单向绑定原则，props 因父组件的更新而变化，自然地将新的状态向下流往子组件，而不会逆向传递。这避免了子组件意外修改父组件的状态的情况，不然应用的数据流将很容易变得混乱而难以理解。
```

这里有它的考虑在里面 如果子组件可以修改父组件的值 这样很麻烦

它想的是 我们直接通过父组件将子组件初始化 然后让其自己可以独立完成某项工作 





组件化的布局 我是真的写不出来啊 就是比较麻烦 

本来就不是很会布局 结果现在 还要将其划分为组件 这不是为难我吗?



# 重启我们的todo

将其作为我们的第一个项目吧 

