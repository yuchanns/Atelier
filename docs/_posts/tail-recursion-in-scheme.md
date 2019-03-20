---
title: Scheme学习笔记（一）——尾递归
date: 2019-02-09 22:17:00
tags:
- scheme
- 尾递归
category: lisp
---
本文为`SCIP`课堂作业思考总结。
<!-- more -->

![The Scheme
Programming Language](https://www.scheme.com/tspl4/canned/medium-cover.png)

[[toc]]

## 尾递归的定义

`尾递归`是函数式编程中，递归函数的一种优化手段。

根据[Wikipedia](https://en.wikipedia.org/wiki/Tail_call)的定义：

> In computer science, a tail call is a subroutine call performed as the final action of a procedure. If a tail call might lead to the same subroutine being called again later in the call chain, the subroutine is said to be tail-recursive, which is a special case of recursion. 
> 在一个程序中，如果一个子程序在该程序的最后一步被调用，就称之为尾调用。如果这个子程序在执行末尾再次调用了自身，就可以称作是尾递归。尾递归是递归调用的一种特殊形式。

## 普通递归的实现

下面以**斐波那契数列**为例。
> 斐波那契数列（Fibonacci sequence）指的是这样一个数列 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233，377，610，987，1597，2584，4181，6765，10946，17711，28657，46368........
> 这个数列从第3项开始，每一项都等于前两项之和。

根据[`SCIP`第一节视频](https://www.bilibili.com/video/av8515129/?p=2)我们可知，`普通递归`实现费氏数列的方法如下：

```scheme
; 普通递归
(define (fib n)
    (if (< n 2)
        n
        (+ (fib (- n 1))
            (fib (- n 2)))))
```

实现起来十分简单，在函数末尾不停调用自身即可。然而我们分析函数执行模式就会发现，这样的递归调用开枝散叶，是树状的——在执行过程中会如滚雪球一样越滚越大。

```
fib(100)
fib(98)       +       fib(99）
fib(96) + fib(97)     fib(98) + fib(97)
...
```

在上面的执行模式分析中，我们可以看到，每一次递归调用执行，它都会进行一次相加操作，而相加操作的参数又来自于递归调用，因此需要等待递归调用返回结果，并且每次调用结果都需要入栈保存。

其中很多地方的计算是重复浪费的。当n非常大的时候（实际上`n=50`就很明显），如果程序按从左到右的顺序执行一遍所有的树状计算，消耗的时间难以想象，而且还会出现栈溢出现象。

这时候就需要使用`尾递归`进行优化了，即每次进行递归调用的时候，直接把当前的结果当做参数放入下次递归调用，无需等待返回结果，也无需保存所有的调用结果。
那么如何实现尾递归呢？

## 尾递归的分析与实现

首先分析一下费氏数列的规律：

```
; fib(100)
; fib(100 - 2) + fib(100 - 1) => fib(98) + fib(99)
; fib(100 - 3) + 2 * fib(100 - 2) => fib(97) + 2 * fib(98)
; 2 * fib(100 - 4) + 3 * fib(100 - 3) => 2 * fib(96) + 3 * fib(97)
; ...
;
; fib(4)
; fib(2) + fib(3)
; fib(1) + 2 * fib(2)
; 2 * fib(0) + 3 * fib(1)
```

根据上述例子，我们可以归纳出以下规律：

* 当n大于等于2时，fib(n) = fib(n - 2) + fib(n - 1) = fib(n - 3) + 2 * fib(n - 2) = 2 * fib(n - 4) + 3 * fib(n - 3) = ... = a * fib(0) + b * fib(1)
* fib(n)的调用次数为n-1次

其中`a`和`b`的含义为：a是上一轮递归中的b，而b是上一轮递归中的a与b之和。b也是个斐波那契数。并且我们可以知道，第一轮递归的a和b必定都为1。
以这个规律，我们可以得出下面的代码片段：

```scheme
; 尾递归

(define (fib_tail a b n)
    (if (> 2 n)
        a
        (fib_tail b (+ a b) (- n 1))))
(define (fib n)
    (fib_tail 1 1 n))
(display (fib 1000))
(display #\newline)
(time (fib 1000))
(display #\newline)
```

运行代码可以轻松得出fib_tail(1000)的值，在我的电脑上消耗时间如下：

```shell
(time (fib 1000))
    no collections
    0.000100112s elapsed cpu time
    0.000099000s elapsed real time
    57744 bytes allocated
```

利用scheme的`named let`语法糖（我在[《初识scheme》](https://www.yuchanns.xyz/posts/2019/01/15/scheme-at-first-sight.html)中提到过），我们可以把代码片段写得更为赏心悦目：

```scheme
; 使用named let定义函数并循环自调用
(define (fib_new number)
    (let fib_tail_new((a 1) (b 1) (n number))
        (if (> 2 n)
            a
            (fib_tail_new b (+ a b) (- n 1)))))
(display (fib_new 1000))
(display #\newline)
```
### 补充一个交互型代码片段
```scheme
; 交互型
(define (fib_read number)
    (if (or (not (integer? number)) (< 100000 number) (< number 0))
        `error
        (if (< number 2)
            number
            (let fib_tail_new((a 1) (b 1) (n number))
                (if (> 2 n)
                    a
                    (fib_tail_new b (+ a b) (- n 1)))))))
(display "请输入一个小于100000的整数：")
(display #\newline)
(display (fib_read (read)))
(display #\newline)
(exit)
```
## 总结
尾递归的`中心思想`在于：**仅在程序的末尾进行调用，并且不等待返回结果**。