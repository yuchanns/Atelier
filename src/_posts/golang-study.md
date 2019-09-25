---
title: 《Go语言学习笔记》要点随记
date: 2019-09-26 00:56:00
category: golang
---
阅读《Go语言学习笔记》随手摘抄——
<!-- more -->

* 退化赋值操作：前提条件：至少有一个新变量被定义，且必须是同一作用域
  ```go
  func main() {
    x := 100  // 定义新变量
    println(&x)

    x, y := 200, "abc"  // 只有y是被定义的新变量
    
    println(&x, x)
    println(y)

    {
      x, y := 200, 300  // 不同作用域，是新变量
      println(&x, x, y)
    }

    x := 300  // 错误
  }
  ```

* 全局变量无**未使用错误**

* 符号名字首字母大小写决定了作用域。首字母大写的为导出成员，可被包外引用

* 常量无**未使用错误**

* 常量组中如不指定类型和初始化值，与上一行非空常量右值相同
  ```go
  func main() {
    const (
      x uint16 = 120
      y  // 与x类型相同
      s = "abc"
      z  // 与s类型相同
    )
  }
  ```

* 使用iota实现一组自增常量值来实现枚举类型。自增默认数据类型为int，可显式指定
  ```go
  const (
    x, a = iota, iota * 10  // 0, 0 * 10
    y, b                    // 1, 1 * 10
    z, c                    // 2, 2 * 10
  )
  ```

* 常量无法读取地址

* 无类型声明的常量不会做强类型检查
  ```go
  const x = 100
  const y byte = x  // 相当于 const y byte = 100

  const a int = 100  // 显式指定类型
  const b byte = a  // 错误
  ```

* `byte`是`uint8`的别名，`rune`是`int32`的别名，但`int`不是`int64`的别名

* slice、map、channel是引用类型

* 语法歧义：转换的目标是指针、单向通道或没有返回值的函数类型，要用括号
  ```go
  (*int)(p)
  (<-chan int)(c)
  (func())(x)
  (func()int)(x)
  ```

* 函数返回局部指针是安全的

* 函数不支持有默认值的可选参数，不支持命名实参，必须安签名顺序传递指定类型和数量的实参

* 变参本质上是切片。将切片作为变参，需要展开操作。如果是数组，要转换成切片。切片可修改原数据

* 命名参数可由return隐式返回
  ```go
  func div(x, y int) (z int, err error) {
    if y == 0 {
      err = errors.New("division by zero")
      return
    }
    z = x / y
    return
  }
  ```

* 普通函数和匿名函数都可以作为结构体字段或经通道传递

* go交叉编译：
  ```sh
  GOOS=linux GOARCH=amd64 go build -gcflags "-N -l" -v  
  ```

* 慎用`defer`，在main中循环读取文件时如果使用defer只会在main结束时调用。应把循环内部逻辑独立成函数，在内部调用defer，这样才函数结束时就会调用。defer降低性能

* 不建议使用`panic`，除非是不可恢复性的错误

* 拼接动态字符串可用`strings.Join`或`bytes.Buffer`
  ```go
  stringA := strings.Join([]string{"a", "a", "a"}, "")

  var b bytes.Buffer
  b.Grow(3)  // 预先准备足够的内存

  for i := 0; i < 3; i++ {
    b.WriteString("a")
  }

  stringB := b.String()
  ```
