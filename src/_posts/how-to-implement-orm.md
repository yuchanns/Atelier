---
title: 如何实现一个orm
date: 2020-01-14 21:35:00
category: 简单易懂的现代魔法
tags:
  - orm
  - gorm
---
:::warning 草稿
本文未完成，还在写作中...
:::
> 创造一样新的事物固然了不起，拥有毅力、勇气和热情去推动它的发展更为难能可贵。

<!-- more -->
![](/images/xorm.jpg)
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
:::tip Active Record Pattern
这是一种软件设计模式，由Martin Fowler[^6]2003年在《企业应用架构模式》一书中提出的概念，用于处理关系数据库在内存中存储的对象数据。
:::

简单来说，Active Record将一张表用类来表示，所有字段描述为类的属性。然后每条查询出来的语句都是类的一个实例。

用户只需要使用面向对象的方式对这些对象作出改动，就可以增删改查数据库的数据。

使用这种模式有其好处：
* （上文提到的）节省编程工作中重复性的sql编码劳动，提高开发效率。
* 进行一定层次的抽象，将各种数据库的语法统一为oop风格的操作方式，抹消方言差别。
### 常见的orm
目前市面上有很多orm都采用了Active Record模式，包括但不限于：
> * Java：MyBatis、Hibernate
> * PHP：Doctrine2、Eloquent
> * Python：SqlAlchemy、Peewee
> * Golang：Prisma、Gorm
> * Ruby：ActiveRecord

对于动态语言来说，使用这类ORM有着**天生**的优势，后面会提到。
## 第一类函数模式
ORM适用于面向对象语言，而众所周知，Golang并不是一门面向对象的语言，尽管拥有内嵌、封装等特性，但并不具备继承特性。因此强行使用orm不是那么合适——注意，这里是说不太合适，并不是不能用，事实上现在很火的gorm正是Golang使用orm的一个最佳选择之一。

就像文章开头列出的那篇文章一样，人们（非正式地）讨论了一种函数式的orm可能性。在以函数为第一等公民的语言（如Golang）中，使用函数式风格的orm可能会更为符合语言的语法习惯。

我们使用函数创造出一个个结构，这些结构都实现了某个规定的接口，然后组装成列表，再经由另一个函数处理成sql语句，接着用于和数据库交互，再将结果返回。
```go
var persons []*person
db.Select(
  WhereLike("name", "yuch%"),
  Field("id, name, age"),
  Limit(10),
  Offset(1)
).Query(&persons)
```
### 产品案例
* [sqlboiler](https://github.com/volatiletech/sqlboiler)
* sqlx + squirrel
## 对比
排除了最初的冲动，在我看来，这两种模式所针对的实际上只是sql语句的构造方式——它们最终的目标都是由代码自动生成sql语句，并将查询结果组装到对象（结构）中。

对于Gorm这类orm来说，缺点无非是实际业务中过于依赖链式操作，在大量重复类似场景的时候使得书写出来的代码显得十分繁琐：
```go
q := db.Where("name like ?", "yuch%")
if age !== "" {
  q = q.Where("age = ?", age)
}
q.Take(&person)
```
而使用函数式orm，用户可以通过函数自行封装，将大量重复的条件封装实现成符合约定的接口，然后使用起来更为简洁：
```go
// Option为约定的接口
func Eq(col, val string) Option {
  return func (q Query) Query {
    if val == "" {
      return q
    }
    
    return WhereEq(col, val)
  }
}

db.Select(
  WhereLike("name", "yuch%"),
  Eq("age", age)
).Query(&person)
```
换言之，当我们全部使用这种方式来封装，可以节省条件判断的重复书写。

关于条件甄别，根据我的经验，其实在Eloquent中已经有一个很好的解决方案`when`：
```php
User::whereLike("name", "yuch%")->when($age != '', function ($query) {
  $query->where("age", $age);
})->select();
```
## orm标准接口
经过上面对于两种模式的orm的分析，到最后的看破（或者也许是我的理解错误）本质，我对于自己的雄心壮志突然变得兴致缺缺。

但是转念一想，抛开这一些，我觉得我还是能整理出一些有用的东西，那就是实现一个orm的标准接口——
## 存在的问题和解决方案
### n+1问题

[^1]: [ORMs and Query Building in Go](https://andrewpillar.com/programming/2019/07/13/orms-and-query-building-in-go/)
[^2]: [jmoiron/sqlx](https://github.com/jmoiron/sqlx)
[^3]: [Active Record](https://guides.rubyonrails.org/active_record_basics.html)
[^4]: [jinzhu/gorm](https://github.com/jinzhu/gorm)
[^5]: [Object-relational mapping](https://en.wikipedia.org/wiki/Object-relational_mapping)
[^6]: [Martin Fowler](https://en.wikipedia.org/wiki/Martin_Fowler_(software_engineer))