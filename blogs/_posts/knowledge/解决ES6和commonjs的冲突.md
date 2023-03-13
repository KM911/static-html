

# ES6的commonjs的区别

其实我一开始是很迷惑的 就是说 什么是ES6的内容 什么是commonjs的内容

```js
let name = { firstName: "dong", lastName: "做个",theme :"test" }
let { firstName, lastName } = name
console.log(firstName)
```

上面就是一个很标准的ES6写法 （变量的解构赋值 let 也是ES6提供的） 我们的node 就可以直接运行 没有任何的我呢提

但是如果你写了这样的代码

```js
import express from 'express'; // 使用import引入模块
let app = express();
app.listen(3000);
```

就会报错 

```
import express from 'express';
^^^^^^

SyntaxError: Cannot use import statement outside a module
```

报错的原因是 import 是将其他模块 作为module 但是我们的require不是 这样的

解决的方法就是在我们`package.json`中添加`"type": "module",` 表示你是使用ES6语法 

目前我发现的ES6和其他commonjs冲突的地方就是这样了 

## 总结

对于ES6 和 commonjs 其实是可以混用的 只有我们的模块是会出现问题的

这里其实还可以将我们的模块导出方案实现更加好 

现在需要写一些项目了 开始加油 

