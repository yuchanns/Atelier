---
title: 用redis做队列
date: 2020-03-15 20:27:00
img: /images/redis.jpg
tags:
  - redis
  - queue
---
最近做一个小型项目，因为**rabbitMQ**等专业软件比较重，故团队决定采用**redis**实现一个轻量的队列。

---
本文还在撰写中...

<!-- more -->
## 收获
在这篇文章中，你可以获得：
* `redigo`包的基本用法
* 初步认知`context`包的作用
* 了解一个功能模块的实现思路以及如何逐步完善的步骤
* **Go**特性(`select`、`channel`和`goroutine`)的利用

最终代码量大概265行左右。

## redis队列的原生用法

## 队列、消息的设计思路

## 改进·Ack机制解决消息丢失问题

## 改进·使用Context实现安全退出
`context`包是Go语言标准库的包之一，在各种接口中被广泛使用。在本文中，将被用来使`goroutine`“安全退出”。

笔者强烈建议读者通过阅读**Go夜读**团队成员饶全成写的这篇《深度解密Go语言之context》[^1]文章来了解`context`的原理；而本文则只着重描写一个使用案例。

## 改进·支持复数消费者

[^1]: [深度解密Go语言之context](https://qcrao.com/2019/06/12/dive-into-go-context/)