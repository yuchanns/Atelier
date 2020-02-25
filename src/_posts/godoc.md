---
title: godoc的正确使用方式
date: 2020-02-26 00:18:00
img: /images/godoc2.png
tags:
  - chore
  - go
---
今天看到有人问go文档阅读的问题。
<!-- more -->

由于众所周知的原因，go的一系列官网在国内都无法访问——虽然说不会翻墙的**gopher**不是好的码农，但是对于代码初学者自然是要秉持着和谐友好的相处态度。

于是我告诉对方，你可以使用`godoc`在本地直接运行一个静态网站，这样就可以快速查看了；并且我丢给对方一张图，是官方提供的godoc使用[例子](https://godoc.org/golang.org/x/tools/cmd/godoc)：
```sh
godoc -http:=6060
```

过了一会儿萌新沮丧地表示，按照截图操作无法得到结果。

我当时就很诧异，然后自己试了一下，结果发现确实运行不了:joy:
```sh
╭─yuchanns@yuchannsdeMBP.lan ~
╰─➤  godoc -http:=6060
flag provided but not defined: -http:
usage: godoc -http=localhost:6060
```
不过根据提示立刻就知道了原因：官网的示例写错了，把冒号放在了等号的左边。乍一看仿佛是给http这个变量赋值的语法，实际上应该是`-http=localhost:6060`，表示端口号，可以简写成`-http=:6060`。

再次运行，水到渠成：
![](/images/godoc.png)

> **注：**
>
> 第一次使用godoc时会自动拉取需要的代码，同样由于网络环境问题，会出现无法拉取的现象。所以我们应该先[把GOPROXY设置为国内镜像源](/posts/2019/09/30/golang-mod/)，就没有问题了。
>
> 如果找不到godoc命令，应该是没安装对应的程序(我就是这样)，只需运行`go get golang.org/x/tools/cmd/godoc`即可。