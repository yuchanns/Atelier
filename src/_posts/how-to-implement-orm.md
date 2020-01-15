---
title: 如何实现一个orm
date: 2020-01-14 21:35:00
category: 简单易懂的现代魔法
tags:
  - orm
  - gorm
---
<span style="color:red">草稿未完成</span>
> 创造一样新的事物固然了不起，拥有毅力、勇气和热情去推动它的发展更为难能可贵。

<!-- more -->
-><lazy-image title="xorm" src="/images/xorm.jpg" /><-
## 前言
事情的起因是前不久，我在网上看到的一篇文章[^1]。

:::tip ORMs and Query Building in Go
作者讲述自己使用`sqlx`[^2]过程中产生的关于更适合Go语法习惯的sql交互方式的思考：**Active Record**[^3]模式还是第一类函数？
:::

在阅读完作者给出的思路和例子之后，我受到鼓舞，开始调查了一些市面上的orm，并妄图自行实现具有这种特性的第一类函数orm——很显然，我还不是个拥有莫大毅力和勇于挑战的开源奉献者，不到两天的尝试后我就变得沮丧、灰心和失望。

`gorm`[^4]真香！

但是，在这一次野心勃勃的失败征途中，我也不是毫无收获——现在我尝试着将我所学习到的一些关于orm的认知，艰难地总结下来。

[[toc]]

## 定义
摘自维基百科[^5]

:::tip Object-relational mapping
对象关系映射(ORM)是一种将数据进行跨语言系统转换以兼容适配的编程技术。它实际上创建了语言层面的虚拟数据库对象，使得用户可以用操作对象的方式来访问和修改数据(对象属性)而不必关心sql底层的转换。
:::

这一概念是随着面向对象编程而兴起的。orm使用方便，它的存在极大地降低了业务代码编写过程中使用者与数据库交互的负担。

下面的伪代码展示了直接操作sql和通过orm间接交互的代码片段区别：
```js
// 直接操作
sql = "SELECT id, name, gender, age FROM persons WHERE name = yuchanns"
rows = db.exec(sql)
age = rows[0]["age"]
sql = "UPDATE persons SET age = 27 WHERE name = yuchanns"
rowsAff = db.exec(sql)

// 通过orm
orm.getByName("yuchanns").find(&person)
age = person.age
person.age = 27
rowsAff = orm.save(&person)
```
可以看到，使用了orm之后我们就节省了书写sql交互语句的功夫，直接以符合人类思考方式的面向对象方式轻松操纵数据~~并且降低了门槛，不需要学习sql知识~~。
## Active Record模式
### 常见的orm

## 第一类函数模式
### 产品案例

## orm标准接口

## 存在的问题和解决方案
### n+1问题

[^1]: [ORMs and Query Building in Go](https://andrewpillar.com/programming/2019/07/13/orms-and-query-building-in-go/)
[^2]: [jmoiron/sqlx](https://github.com/jmoiron/sqlx)
[^3]: [Active Record](https://guides.rubyonrails.org/active_record_basics.html)
[^4]: [jinzhu/gorm](https://github.com/jinzhu/gorm)
[^5]: [Object-relational mapping](https://en.wikipedia.org/wiki/Object-relational_mapping)