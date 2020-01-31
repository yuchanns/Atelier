---
title: go编译工具的使用之汇编
date: 2020-01-31 15:50:59
category: golang
tags:
  - chore
  - tool
---
:::warning 未完成
施工中...
:::
有时候我们想要知道写出来的代码是怎么编译执行的，这时候`go tool compile`就是一个很好用的工具。

本文相关代码[yuchanns/gobyexample](https://github.com/yuchanns/gobyexample/tree/master/assembly)
<!-- more -->

[[toc]]

## 如何输出汇编代码
有三种方法可以输出go代码的汇编代码：
* `go tool compile` 生成obj文件
* `go build -gcflags` 生成最终二进制文件
* 先`go build`然后在`go tool objdump` 对二进制文件进行反汇编

当然，具体行为还需要在这些命令后面加上具体的<mark>flag</mark>。flag的内容可以通过查阅[官方文档](https://godoc.org/cmd/compile#hdr-Command_Line "command line")获得。
:::tip 本文涉及Flags说明
-N 禁止优化

-S 输出汇编代码

-l 禁止内联
:::
### 什么是内联
如果学过<mark>c/c++</mark>就知道，通过`inline`关键字修饰的函数叫做内联函数。内联函数的优势是在编译过程中直接展开函数中的代码，将其替换到源码的函数调用位置，这样可以节省函数调用的消耗，提高运行速度。适用于函数体短小且频繁调用的函数，如果函数体太大了，会增大目标代码。是一种**空间换时间**的做法。

go编译器会智能判断对代码进行优化和使用汇编。而我们在分析学习代码调用情况的时候需要禁止掉这些优化，避免混淆理解。

以下我们使用`go build -gcflags="-N -l -S" file`来获得汇编代码。

### 获取一份简单的汇编代码
现在手上有一份关于`range`的代码，但是我们运行之后出现了一些问题[^1][^2]：
```go
package assembly

import "fmt"

func RangeClause() {
	arr := []int{1, 2, 3}
	var newArr []*int
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}
```
结果输出了三个3。

也许我们在学习过程中见过类似的错误，然后设法(或者别人告诉我们怎么)避免错误，但是仍然百思不得其解，知其然不知其所以然。

<details>
<summary>避免错误的写法</summary>

将`&v`替换成`&arr[i]`
```go{8-9}
package assembly

import "fmt"

func RangeClause() {
	arr := []int{1, 2, 3}
	var newArr []*int
	for i := range arr {
		newArr = append(newArr, &arr[i])
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}
```
</details>

这时候获取汇编代码就可以排上用场了。

执行`go build -gcflags="-N -l -S" range_clause.go`，得到下面这份输出结果：
<details>
<summary>汇编结果</summary>

<!-- <<< @/src/.vuepress/snippets/plan9_range.go -->
</details>
看着输出结果，很cool~~~但是看不懂:(

## 汇编的简单知识
go使用的汇编叫做`plan9汇编`。最初go是在plan9系统上开发的，后来才在Linux和Mac上实现。

关于plan9汇编的入门，推荐看这个视频[^3]：

-><lazy-video src="//player.bilibili.com/player.html?aid=46494102&cid=81455226&page=1" /><-

其中一些汇编知识是通用的[^4]，**GoDoc**也提供了go汇编的快速引导[^5]，另外也有一部分可以参考plan9汇编手册[^6]。
### 寄存器
寄存器是CPU内部用来存放数据的一些小型存储区域，用来暂时存放参与运算的数据和运算结果。

:::tip 分类
这里提及的寄存器可以分为三类：
* 后缀Register的寄存器属于数据类寄存器
* 后缀为Pointer的寄存器属于指针类寄存器
* 后缀为Index的寄存器属于索引类寄存器
:::
|助记符|名字|用途|
|---|---|---|
|AX|累加寄存器(AccumulatorRegister)|用于存放数据，包括算术、操作数、结果和临时存放地址|
|BX|基址寄存器(BaseRegister)|用于存放访问存储器时的地址|
|CX|计数寄存器(CountRegister)|用于保存计算值，用作计数器|
|DX|数据寄存器(DataRegister)|用于数据传递，在寄存器间接寻址中的I/O指令中存放I/O端口的地址|
|SP|堆栈顶指针(StackPointer)|指向栈顶|
|BP|堆栈基指针(BasePointer)|保存在进入函数前的栈顶基址|
|SB|静态基指针(StaticBasePointer)|go汇编的伪寄存器。`foo(SB)`用于表示变量在内存中的地址，`foo+4(SB)`表示foo起始地址往后偏移四字节。|
|SI|源变址寄存器(SourceIndex)|用于存放源操作数的偏移地址|
|DI|目的寄存器(DestinationIndex)|用于存放目的操作数的偏移地址|
### 操作指令
用于指导汇编如何进行。以下指令后缀<mark>Q</mark>说明是64位上的汇编指令。
|助记符|指令种类|用途|示例|
|---|---|---|---|
|MOVQ|传送|数据传送|`MOVQ AX, 48`表示把48传送AX中|
|LEAQ|传送|有效地址传送|`LEAQ BX, AX`表示把AX有效地址传送到BX中|
|PUSHQ|传送|栈压入|`PUSHQ AX`表示先修改栈顶指针，将AX内容送入新的栈顶位置|
|POPQ|传送|栈弹出|`POPQ AX`表示先弹出栈顶的数据，然后修改栈顶指针|
|ADDQ|运算|相加并赋值|`ADDQ AX, BX`表示BX和AX的值相加并赋值给AX|
|SUBQ|运算|相减并赋值|略，同上|
|MULQ|运算|无符号乘法|`MULQ DX`DX为乘数，被乘数总是指定为AX|
|DIVQ|运算|无符号除法|字节操作时，被除数在AX中；字操作时，被除数的高位放在DX中，低位放在AX中|
|CMPQ|运算|对两数相减，比较大小|`CMPQ SI CX`表示比较SI和CX的大小。与SUBQ类似，只是不返回相减的结果|
|CALL|转移|调用函数|`CALL runtime.printnl(SB)`表示通过<mark>printnl</mark>函数的内存地址发起调用|
|JMP|转移|无条件转移指令|`JMP 389`无条件转至389地址处|
|JEQ|转移|条件转移指令|`JEQ 389`ZF标志位=1时转至389地址处|
|JB|转移|无符号数的条件转移指令|略，同上|
### 标志位
在操作指令里提到了`ZF标志位`，这是汇编的标志位
|助记符|名字|用途|
|---|---|---|
|OF|溢出|0为无溢出 1为溢出|
|CF|进位|0为最高位无进位或错位 1为有|
|PF|奇偶|0表示数据最低8位中1的个数为奇数，1则表示1的个数为偶数|
|AF|辅助进位||
|ZF|零|0表示结果不为0 1表示结果为0|
|SF|符号|0表示最高位为0 1表示最高位为1|

这么一通信息轰炸下来，作为初学者可能已经头晕脑胀记不住了，其实是否记住这并不重要——后面分析用到了再回来查阅意思即可。

## 分析汇编代码
### 从1+1开始
> “好了，现在我们已经学会了加减乘除四则运算，接下来我们来解答一下这道微积分的题目”XD

我们先从一个简单的范例`1+1`来实践一下对汇编代码的分析：
```go
package assembly

func Add() {
	a := 1 + 1
	println(a)
}
```
汇编结果：
```go{2,23}
"".Add STEXT nosplit size=32 args=0x0 locals=0x18
        0x0000 00000 (add.go:3)      TEXT    "".Add(SB), ABIInternal, $24-0
        0x0000 00000 (add.go:3)      MOVQ    (TLS), CX
        0x0009 00009 (add.go:3)      CMPQ    SP, 16(CX)
        0x000d 00013 (add.go:3)      JLS     77
        0x000f 00015 (add.go:3)      SUBQ    $24, SP
        0x0013 00019 (add.go:3)      MOVQ    BP, 16(SP)
        0x0018 00024 (add.go:3)      LEAQ    16(SP), BP
        0x001d 00029 (add.go:3)      FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (add.go:3)      FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (add.go:3)      FUNCDATA        $2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (add.go:4)      PCDATA  $0, $0
        0x001d 00029 (add.go:4)      PCDATA  $1, $0
        0x001d 00029 (add.go:4)      MOVQ    $2, "".a+8(SP)
        0x0026 00038 (add.go:5)      CALL    runtime.printlock(SB)
        0x002b 00043 (add.go:5)      MOVQ    "".a+8(SP), AX
        0x0030 00048 (add.go:5)      MOVQ    AX, (SP)
        0x0034 00052 (add.go:5)      CALL    runtime.printint(SB)
        0x0039 00057 (add.go:5)      CALL    runtime.printnl(SB)
        0x003e 00062 (add.go:5)      CALL    runtime.printunlock(SB)
        0x0043 00067 (add.go:6)      MOVQ    16(SP), BP
        0x0048 00072 (add.go:6)      ADDQ    $24, SP
        0x004c 00076 (add.go:6)      RET
        0x004d 00077 (add.go:6)      NOP
        0x004d 00077 (add.go:3)      PCDATA  $1, $-1
        0x004d 00077 (add.go:3)      PCDATA  $0, $-1
        0x004d 00077 (add.go:3)      CALL    runtime.morestack_noctxt(SB)
        0x0052 00082 (add.go:3)      JMP     0
```
第一行是go汇编的固定开头，指定过程名字为`"".Add`，`args=0x0 locals=0x18`则对应第二行的`$24-0`是十六进制和十进制的转化。

第二行`TEXT`是一个伪操作符，以过程名的内存地址(`"".Add(SB)`)为定义过程的内存起点(回想一下`foo(SB)`是什么意思？)，然后在栈上为过程分配内存。`$24-0`其中`24`表示栈帧的大小为24字节(跟函数内部变量数据类型以及个数有关，例如这里是两个整型变量，就是2x8=16字节，然后还有一个8字节的整型用来存储BP值，所以一共24个字节)，`0`则表示调用方传入的参数大小。`ABIInternal`应该是**应用程序二进制接口内部**(Application Binary Interface Internal)的意思，不重要。

第三行的`MOVQ (TLS), CX`，我们现在可以回头查阅一下MOVQ是干什么用的——用于数据传送。可以看出来是把一个<mark>(TLS)</mark>赋值给CX(计数寄存器)。但是这个(TLS)是什么呢？它实际上也是一个伪寄存器，保存了指向当前G(保存`gorountine`的一种数据结构)的指针[^7]。

第四行则是比较当前栈顶指针和G指针正偏移16字节的地址大小。

如果左边小于右边就跳到`0x004d`(从十进制77转换为十六进制后的值)这个地址。我们先看看这个地址有什么内容：`NOP`意思是**No Operation**，无操作数。看了下这里是运行到了`add.go`文件的第六行，也就是一个`}`，所以是没有任何操作的。往下又回到了第三行，先不管。

回到第五行，如果没有达成上面的条件判断，就不会进行内存地址跳转，而是继续执行第六行的代码。

这一行代码是将栈顶地址减去**24**字节的内存容量，并把结果存到SP中。这一步的意思是，栈顶指针指向了这个过程(`"".Add`)的内存地址开头，因为第二行给过程分配内存的时候是分配在栈上的，导致栈顶指针发生了移动。

第七行把`BP`的值赋予了`16(SP)`，意思是从栈顶开始第十六个字节位置开始的那个整型变量。接着第八行把16(SP)的地址赋给了BP。

第九到十三行`FUNCDATA`和`PCDATA`是由编译器生成的，作用是告诉<mark>GC</mark>(**GarbageCollection**)区分堆栈中变量的类型。`$数字`表示变量属于什么类型(参数？本地？)，而后面的`gclocals·xxxxx(SB)`则是引用了一个隐藏的包含了GC标记的变量。注意这一行用到了`·`(middle dot)，用来代替go源文件中的`.`，因为在汇编中此符号已经被作为标点符号来解析。

第十四行，把结果2赋给变量a。这里有两个点需要注意：首先`$2`并不是表示上面那个`FUNCDATA`创建的变量，而是`1+1`的结果值。如果上面的代码改成了`1+2`那么此处会变成`$3`；`"".a+8(SP)`并不是一个加法运算，而是表示距离栈顶8字节位置的那个变量a，这只是一种go汇编语法的强制规定，必须把变量名和内存地址使用`+`连起来表示而已，没有实际意义。

第十五行，源码来到了第六行，调用了`runtime`包的<mark>printlock</mark>方法，根据名字可以看出这是打印前进行加锁的。

第十六和十七行的效果是把变量a放到AX寄存器中，然后把寄存器的地址赋给栈顶指针。

第十八行、十九和二十行则是打印栈顶指针的内容、打印换行符和解锁。

第二十一行把函数栈上记录的BP值还给BP，而二十二行将栈顶指针指向函数末尾。最后函数退出。


[^1]: [for和range的实现|Go语言的设计和实现](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-for-range/)
[^2]: [Common Mistakes|Go](https://github.com/golang/go/wiki/CommonMistakes)
[^3]: [plan9汇编入门|go夜读](https://www.bilibili.com/video/av46494102)
[^4]: [Assembly Programming|Tutorialspoint](https://www.tutorialspoint.com/assembly_programming/index.htm)
[^5]: [A Quick Guide to Go's Assembler](https://golang.org/doc/asm#introduction)
[^6]: [A Manual for the Plan 9 assembler](https://9p.io/sys/doc/asm.html)
[^7]: [teh-cmc/go-internals](https://github.com/teh-cmc/go-internals/blob/master/chapter1_assembly_primer/README.md#splits)