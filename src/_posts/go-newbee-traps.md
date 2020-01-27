---
title: go新手常见陷阱
date: 2020-01-25 22:56:00
category: golang
tags:
  - chore
---
:::warning 未完成
本文还在编写中...
:::
节选自[《50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs》](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)，仅摘录一些作者比较在意的片段。

关联仓库[yuchanns/gobyexample](https://github.com/yuchanns/gobyexample/tree/master/newbee_traps)
<!-- more -->

[[toc]]

## 初级篇
### 未指定类型变量不能用nil初始化
支持`nil`初始化的变量类型有`interface`、`function`、`pointer`、`map`、`slice`和`channel`。所以使用nil初始化未指定类型的变量会导致编译器无法自动推断：
```go
package main

func main() {
  var x interface{} = nil
  _ = x
}
```
### 初始化为nil的map无法添加元素
应该使用<mark>make</mark>方法声明来对`map`进行实际的内存分配；slice可以使用<mark>append</mark>方法对值为nil追加元素。

当然，初始化slice时最好预估一个长度，节省重复扩容开销。
```go
package main

func main() {
  m := make(map[string]int)
  // var m map[string]int // 错误示范，初始化值为nil
  m["one"] = 1 // 如果对上述值为nil的map添加元素，会报错

  var s []int
  s = append(s, 1) // 正确的slice追加元素用法
}
```
### 初始化string不能为nil
`nil`不支持`string`类型的初始化。它的初始值应为空字符串：
```go
package main

func main() {
  var s string
  // var s string = nil // 错误示范，cannot use nil as type string in assignment
  if s == "" {
    s = "default"
  }
}
```
### range遍历slice和array时的非预期值用法
使用<mark>rang</mark>进行遍历时，第一个值固定返回索引，第二个值固定返回值。

如果只想用值，在索引位置可用`_`来接收，节省复制开销。

在大数组中最好不使用range来遍历，因为range的本质是对索引和值的复制和再赋值，开销较大；推荐使用`for i := 0; i < len(s); i++ {}`的方式进行。

```go
package main

import "fmt"

func main() {
  x := []string{"a", "b", "c"}

  for _, v := range x { // 索引不进行复制
    fmt.Println(v)
  }
}
```
### 使用独立的一维slice组装创建多维数组
分为两步：
* 创建外层slice
* 为每个元素分配一个内层slice
这样的好处是每个内层数组都是独立的，更改不影响其他内层数组。
```go
package main

func main() {
  x := 2
  y := 4
  
  table := make([][]int, x)
  for i := range table {
    table[i] = make([]int, y)
  }
}
```
### 字符串是不可改变的
字符串是只读的二进制slice，无法通过访问索引的方式更改个别字符。如果想要更改，需要转化成`[]byte`类型。

对于<mark>UTF8</mark>字符串，实际上应该转换为`[]rune`类型，避免出现字节更新错误。
```go
package main

import "fmt"

func main() {
  x := "test"
  xbytes := []byte(x)
  xbytest[0] = 'T'

  y := "s界"
  yrunes := []rune(y)
  yrunes[0] = '世'

  fmt.Println(string(xbytes))
  fmt.Println(string(yrunes))
}
```
