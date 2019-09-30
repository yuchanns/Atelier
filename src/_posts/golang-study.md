---
title: 《Go语言学习笔记》要点随记
date: 2019-09-26 00:56:00
category: golang
---
阅读《Go语言学习笔记》随手摘抄——
<!-- more -->

[[TOC]]

<details>
<summary>kmpIndex by golang</summary>

```go
package main

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

func getNext(s string, next []int) {
	j, i := -1, 0
	length := len(s)
	next[0] = -1
	for i < (length - 1) {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			if s[i] == s[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}
}

func kmpIndex(s, sub string) (int, error) {
	fmt.Println("Comparing: ", s, " with ", sub)
	length := len(s)
	slength := len(sub)
	next := make([]int, slength)
	getNext(sub, next)
	i, j := 0, 0
	var a, b bytes.Buffer
	a.Grow(length)
	b.Grow(length)
	for (i < length) && (j < slength) {
		if j != -1 {
			a.WriteString(string(s[i]))
			b.WriteString(string(sub[j]))
			fmt.Println("Comparing: ", a.String(), " with ", b.String())
		}
		if j == -1 || s[i] == sub[j] {
			if j == -1 {
				a.Reset()
				b.Reset()
			}
			j++
			i++
		} else {
			j = next[j]
		}
		time.Sleep(time.Second)
	}
	if j == slength {
		return i - j, nil
	}
	return 0, errors.New("Not found")
}

func main() {
	result, err := kmpIndex("ABABCABCABABA", "ABABA")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Match completed on index:", result)
	}
}

```
</details>

## 数据
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

* 访问不存在的键值，使用ok-idiom模式判断
  ```go
  func main() {
      m := map[string]int{
          "a": 1,
          "b": 2,
      }

      m["a"] = 3  // 修改
      m["c"] = 4  // 新增

      if v, ok := m["d"]; ok {
          println(v)
      }

      delete(m, "d")  // 删除键值对
  }
  ```

* 字典`no addressable`，不能直接修改value成员（结构或数组），应该返回整个value，修改完毕后设置字典键值，或使用指针类型（指值是指针）
  ```go
  type user struct{
      name string
      age byte
  }

  func main() {
      m := map[int]user{
          1: {"Tom", 19},
      }

      u := m[1]
      u.age += 1
      m[1] = u

      m2 := map[int]*user{
          1: &user{"Jack", 20},
      }

      m2[1].age++
  }
  ```

* 字典初始化为空，未初始化则为nil，nil无法赋值，可以读
  ```go
  func main() {
      var m map[string]int
      println(m["a"])  // nil
      m["a"] = 1  // panic
      
      m2 := map[string]int{}

      fmt.Println(m == nil, m2 == nil)  // true false
  }
  ```
* 字典迭代期间增删键值是**安全**的。不能对字典进行并发操作，会导致进程崩溃
  * 启用`data race`检查问题
    ```sh
    go run -race test.go
    ```
  * 使用`sync.RWMutex`实现同步，避免读写操作同时进行
    ```go
    import (
      "sync"
      "time"
    )

    func main() {
      var lock sync.RWMutex
      m := make(map[string]int)

      go func() {
        for {
          lock.Lock()
          m["a"] += 1
          lock.Unlock()

          time.Sleep(time.Microsecond)
        }
      }()

      go func() {
        for {
          lock.RLock()

          _ = m["b"]
          lock.RUnlock()

          time.Sleep(time.Microsecond)
        }
      }()

      select {}
    }
    ```

* 字典对象本身就是指针包装；在创建时预先准备足够的空间有助于提升性能，减少扩张时的内存分配和重新哈希操作；对于海量小对象直接用字典存储键值拷贝数据，缩短gc时间；字典不会收缩内存，可适当替换成新对象
## 结构体
* 结构体建议使用命名初始化，否则作为字段类型时无法直接初始化；只有字段类型全部支持时，才能做相等操作；可使用指针操作结构字段，不能是多级指针
  ```go
  func main() {
      type file struct{
          name string
          attr struct{  // 匿名结构类型字段
              owner int
              perm  int
          }
      }

      f := file{
          name: "test",

          // attr: {  // 错误
          //     owner: 1,
          //     perm:  0755,
          // },
      }

      f.attr.owner = 1  // 正确方式
      f.attr.perm = 0755
  }
  ```

* 空结构自身和作为数组元素类型长度都为0。可作为通道元素类型用于事件通知
  ```go
  func main() {
    exit := make(chan struct{})

    go func() {
      println("hello, world!")
      exit <- struct{}{}
    }()

    <-exit
    println("end.")
  }
  ```
* 匿名字段隐式地以类型名作为字段名，其成员可直接引用，但初始化时需当做独立字段；如嵌入其他包中的类型，则隐式字段名不包括包名；不能将基础类型和其指针同时嵌入，因为两者隐式名字相同；如果出现重名，就无法直接引用，需显式字段引用
  ```go
  type attr struct{
    perm int
  }

  type file struct{
    name string
    attr
    os.File
  }

  func main() {
    f := file{
      name: "test",
      attr: attr{  // 显式初始化匿名字段
        perm: 0755,
      },
      File: os.File()  // 不含包名
      *int
      // int  // 不可同时嵌入
    }

    f.perm = 0644  // 直接引用匿名字段成员
    println(f.perm)
  }
  ```
* 字段标签不是注释，是描述字段的元数据，**是**类型的组成部分，**不属于**数据成员。可用反射获取，常被用于格式校验、数据库关系映射等
  ```go
  type user struct{
    Name string `string:"昵称"`
    Sex  byte   `byte:"性别"`
  }

  func main() {
    u := user{"yuchanns", 1}
    v := reflect.ValueOf(u)
    t := v.Type()

    for i, n := 0, t.NumField(); i < n; i++ {
      fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
    }
    // string:"昵称": yuchanns
    // byte:"性别": 1
  }
  ```

* 在内存分配时，字段须做对齐处理，通常以所有字段中最长的基础类型宽度为标准。如果仅有空结构类型字段或其是最后一个字段，会按1对齐，长度为0

* 结构体的方法接收一个前置参数，称作receiver，类似于类中的this；receiver可以是任何**除接口和指针以外**的类型；当它是基础类型时，在方法中被调用是以复制的形式，是指针类型时，不会被复制；指针类型的receiver必须是合法指针（包括nil）或能取得实例地址；不能用多级指针调用方法
  ```go
  type Name int

  func (receiver Name) test() {
    receiver++
    println("test:", receiver)
  }

  func (receiver *Name) testPointer() {
    *receiver++
    println("testPointer:", *receiver)
  }

  func main() {
    var t Name = 1
    t.test()  // test: 2
    println("after test:", t)  // after test: 1
    t.testPointer()  // testPointer: 2
    println("after testPointer:", t)  // after testPointer: 2
  }
  ```

* 选择receiver的类型：
  * 指针类型
    * 需要修改实例状态
    * 大对象（减少复制成本）
    * 包含Mutex等同步字段（避免锁操作无效）
    * 无法确定的状况
  * 普通类型
    * 无需修改状态的小对象和固定值
    * 引用类型、字符串、函数等指针包装对象

* 可以像访问匿名字段成员那样来调用方法，同样具有同名遮蔽问题，可以利用这点实现覆盖操作

* 方法集决定结构体是否实现了某个接口
  * 类型T方法集合包含所有的receiver T方法
  * 类型*T方法集包含所有的receiver T + *T方法
  * 匿名嵌入S，T方法集包含所有receiver S方法
  * 匿名嵌入*S，T方法集包含所有receiver S + *S方法
  * 匿名嵌入S或*S，\*T方法集包含所有receiver S + *S方法

---
未完待续 >>