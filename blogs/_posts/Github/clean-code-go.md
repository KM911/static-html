---
title: clean-code-go
mathjax: false
date: 2023-04-27 07:59:29
tags:
categories:
---

## Clean code

首先,我并不认为存在银弹,真的可以让所有人的代码都变得容易阅读,可维护性好.但是就像大家在学习C语言默认都会将 i j k 作为循环的标识符一样,共同遵守一些准则,将会有助于大家理解代码,而不是你的是 z x y 他的是 a b c.

## 实践案例

### 有意义的循环标识符

尽管我们经常使用 `i` 作为循环的标识符,我还是推荐大家使用更有意义的标识符. 原因很简单, 我们知道 `index` 肯定是一个

```go
var numList = []int{5, 4, 3, 2, 1}
for i := 0; i < len(numList); i++ {
    if numList[i]>numList[i+1] {
        numList[i], numList[i+1] = numList[i+1], numList[i]
    }
}
```

```go
var numList = []int{5, 4, 3, 2, 1}
var numListLens = len(numList)
for index := 0; index < numListLens ; index++ {
    if numList[index] > numList[index+1] {
        numList[index], numList[index+1] = numList[index+1], numList[index]
    }
}
```

如果你觉得不明显的话,可以看看其他的例子.





### 没有必要强制else

倒不是说 我们需要就是节省代码,这里其实只有 AB两个结果,如果你使用 `else` 肯定是想说明还有第三种结果,这样不容易引发歧义.

```go
func IsTIUFormat(file_ string) bool {
	filename := oslib.FileName(file_)
	if len(filename) == 12 && strings.HasPrefix(filename, "KM") {
		return true
	} else {
		return false
	}
}
```

```go
func IsTIUFormat(file_ string) bool {
	filename := oslib.FileName(file_)
	if len(filename) == 12 && strings.HasPrefix(filename, "KM") {
		return true
	}
	return false
}
```

## 利用switch替换同质化的if

```go


```

```go

	
```

