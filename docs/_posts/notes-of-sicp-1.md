---
title: 学习记录-计算机程序的构造和解释-第一章
date: 2019-05-04 14:14:00
tags:
  - scheme
  - sicp
category: 学习笔记
---
随性做一些学第一章的记录。
<!-- more -->
## 程序设计的基本元素

## 过程和它们所产生的计算
### 练习 1.19
> 使用尾递归计算斐波那契数列的过程中，状态变量a和b有以下变换规则，a←a+b和b←a，现在将这种变换称为T变换。从1和0开始将T反复应用n次，将产出一对数Fib(n+1)和Fib(n)。
>
> 换句话说，斐波那契数可以通过将T<sup>n</sup>应用于对偶(1, 0)而产生出来。
>
> 现在将T看做是变换族T<sub>pq</sub>中p=0且q=1的特殊情况，其中T<sub>pq</sub>是对偶(a, b)按照a←bq+aq+ap和b←bp+aq规则的变换。请证明，如果我们应用变换T<sub>pq</sub>两次，其效果等同于应用同样形式的一次变换T<sub>p'q'</sub>，其中p'和q'可以由p和q计算出来。这就指明了一条求出这种变换的平方的路径。将所有的这些集中到一起，形成下面的过程，其运行只需要对数的步数：
> ```scheme
> (define (fib n)
>   (fib-iter 1 0 0 1 n))
> (define (fib-iter a b p q count)
>   (cond ((= count 0) b)
>         ((even? count)
>           (fib-iter a
>                     b
>                     <??>      ; compute p'
>                     <??>      ; compute q'
>                     (/ count 2)))
>         (else (fib-iter (+ (* b q) (* a q) (* a p))
>                         (+ (* b q) (* a q))
>                         p
>                         q
>                         (- count 1)))))
> ```

已知a<sub>2</sub>←b<sub>1</sub>q+a<sub>1</sub>q+a<sub>1</sub>p和b<sub>2</sub>←b<sub>1</sub>p+a<sub>1</sub>q，这是a<sub>2</sub>, b<sub>2</sub>和a<sub>1</sub>, b<sub>1</sub>的关系。所以要求出这种变换的平方的路径就是求出a<sub>3</sub>, b<sub>3</sub>和a<sub>1</sub>, b<sub>1</sub>的关系。
应用关系可得：
> a<sub>3</sub>=b<sub>2</sub>q+a<sub>2</sub>q+a<sub>2</sub>p
> 
> b<sub>3</sub>=b<sub>2</sub>p+a<sub>2</sub>q

将a<sub>1</sub>和b<sub>1</sub>代入计算可得：

> a<sub>3</sub>=(b<sub>1</sub>p+a<sub>1</sub>q)q+(b<sub>1</sub>q+a<sub>1</sub>q+a<sub>1</sub>p)q+(b<sub>1</sub>q+a<sub>1</sub>q+a<sub>1</sub>p)p
>> a<sub>3</sub>=b<sub>1</sub>qp+a<sub>1</sub>q<sup>2</sup>+b<sub>1</sub>q<sup>2</sup>+a<sub>1</sub>q<sup>2</sup>+a<sub>1</sub>qp+b<sub>1</sub>qp+a<sub>1</sub>qp+a<sub>1</sub>p<sup>2</sup>
>>> a<sub>3</sub>=b<sub>1</sub>(q<sup>2</sup>+2pq)+a<sub>1</sub>(2qp+2q<sup>2</sup>+p<sup>2</sup>)
>>>
>>> 由于a<sub>3</sub>=b<sub>1</sub>q'+a<sub>1</sub>q'+a<sub>1</sub>p'，所以：
>>> 
>>> q' = q<sup>2</sup>+2pq
>>>
>>> p' = 2qp+2q<sup>2</sup>+p<sup>2</sup>-q'=q<sup>2</sup>+p<sup>2</sup>

所以答案为：
 ```scheme
 (define (fib n)
   (fib-iter 1 0 0 1 n))
 (define (fib-iter a b p q count)
   (cond ((= count 0) b)
         ((even? count)
           (fib-iter a
                     b
                     (+ (* q q) (* p p))      ; compute p'
                     (+ (* q q) (* 2 q p))    ; compute q'
                     (/ count 2)))
         (else (fib-iter (+ (* b q) (* a q) (* a p))
                         (+ (* b q) (* a q))
                         p
                         q
                         (- count 1)))))
 ```
 ## 用高阶函数做抽象