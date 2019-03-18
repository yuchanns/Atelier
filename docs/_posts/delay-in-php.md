---
title: php延时任务
date: 2018-10-08 23:27:49
tags:
- swoole
- 延时任务
category: php
---
<!-- more -->
最近工作上用到了延时任务（订单取消等等），考虑了几种延时方法。
先记录一下使用swoole的方案，其余的等待更新。
# 使用swoole实现延时任务
## 安装swoole扩展
swoole的安装参考[官方文档](https://wiki.swoole.com/wiki/page/6.html)：
> 安装的前置条件
>php-5.3.10 或更高版本
>gcc-4.4 或更高版本
>make
>autoconf
>pcre

```shell
pecl install swoole
```

修改php.ini添加扩展：
```vim
extension=swoole.so
```

## 使用swoole延时任务
使用方法同样可以查询[官方文档](https://wiki.swoole.com/wiki/page/480.html)。
> 持续触发型

```php
//每隔2000ms触发一次
$task_id = swoole_timer_tick(2000, function ($timer_id) {
    echo "tick-2000ms\n";
});
```

> 预定时间触发一次

```php
//3000ms后执行此函数
$task_id = swoole_timer_after(3000, function () {
    echo "after 3000ms.\n";
});
```

> 清除定时器

```php
swoole_timer_clear($task_id);
```

下面给出我的例子：
```php
#! /usr/bin/php

# delay_swoole.php

<?php
swoole_time_tick(2000, function() {
    $file = '/www/delay.log';
    $content = file_get_contents($file);
    file_put_contents($file, $content.' | '.date('Y-m-d H:i:s'));
});
```

在文件开头指定了
```
#! /usr/bin/php
```
即php执行程序位置，是为了可直接执行本文件。

## 使用nohup和&后台运行delay_swoole.php
> nohup 是 no hang up 的缩写，由Command参数和任何相关的Arg参数指定的命令，忽略所有挂断（SIGHUP）信号。

简而言之，nohup可以使用户退出账号/终端之后依旧保持进程运行。

使用案例：
```shell
nohup command > myout.file 2>&1
```

> 参数解释
> 0 – stdin (standard input)，1 – stdout (standard output)，2 – stderr (standard error) ；
> 2>&1的意思就是将标准错误（2）重定向到标准输出（&1），标准输出（&1）再被重定向输入到myout.file文件中。

&和nohup具有不同的功能。

> &是指将执行程序转为后台运行。忽略所有中断（SIGINT）信号。

简而言之，&无法用Ctrl+C等方法终端执行。

通常我们可以把nohup和&结合起来使用，这样就具备了忽略挂断和中断信号的特点。

```shell
nohup delay_swoole.php > /dev/null 2>&1 &
```

这样就实现了**使用swoole实现延时任务**的目的。

|标题一|标题二|标题三|
|---|---|---|
|内容1|内容2|内容3|
|内容1|内容2|内容3|
