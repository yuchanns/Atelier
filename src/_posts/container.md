---
title: 容器的意义
date: 2019-05-28 23:27:00
tags:
  - swoole
category: php
---
对于传统的php-fpm进程来说，每一个请求都是一次短暂的php框架生命周期。实例化一个类，业务处理，自动销毁，一切都是家常便饭。
<!-- more -->

::: tip
在线演示
:::

<iframe height="600px" width="100%" src="https://repl.it/@yuchanns/PHPContainer?lite=true" scrolling="no" frameborder="no" allowtransparency="true" allowfullscreen="true" sandbox="allow-forms allow-pointer-lock allow-popups allow-same-origin allow-scripts allow-modals"></iframe>

## 理由
当我们使用swoole这种内存常驻型的扩展，就可以使用单例模式来代替每一轮请求使用到的对象的创建-销毁。

我们可以借鉴其他语言中的`IOC容器`这一概念实现单例模式。一方面节省了资源的开销，另一方面便于管理对象。同时使用`反转控制`和`依赖注入`也有利于解耦：

> 控制反转（Inversion of Control，缩写为IoC），是面向对象编程中的一种设计原则，可以用来减低计算机代码之间的耦合度。其中最常见的方式叫做依赖注入（Dependency Injection，简称DI），还有一种方式叫“依赖查找”（Dependency Lookup）。通过控制反转，对象在被创建的时候，由一个调控系统内所有对象的外界实体将其所依赖的对象的引用传递给它。也可以说，依赖被注入到对象中。

## 代码
```php
<?php

use Swoole\Http\ {
    Request,
    Response
};

class Container
{
    /**
     * @var $instance Container
     */
    private static $instance = null;

    private $objPool = [];

    public function __construct()
    {
        echo "Container Initial".PHP_EOL;
    }

    public static function getInstance()
    {
        if (!self::$instance) {
            self::$instance = new self;
        }
        return self::$instance;
    }

    public function get($name)
    {
        if (!isset($this->objPool[$name])) {
            $this->objPool[$name] = new $name;
        }
        return $this->objPool[$name];
    }

    public function __destruct()
    {
        echo "Container Release".PHP_EOL;
    }
}

class Test
{
    private $value;

    public function __construct()
    {
        echo "Test initial".PHP_EOL;
        $this->value = 1;
    }

    public function get()
    {
        return $this->value;
    }

    public function add()
    {
        $this->value += 1;
    }

    public function __destruct()
    {
        echo "Test release".PHP_EOL;
    }
}

class App
{
    public function run()
    {
        $http = new swoole_http_server("127.0.0.1", 9501);

        $http->on("start", function ($server) {
            echo "Swoole http server is started at http://127.0.0.1:9501".PHP_EOL;
        });

        $http->on("request", function (
            Request $request,
            Response $response
        ) {
            $init = memory_get_usage();
            /* @var $test Test*/
            $test = Container::getInstance()->get('Test');
            $after = memory_get_usage();
            $test->add();
            $response->header("Content-Type", "text/plain");
            $response->end("usage:" . ($after - $init) . 'B' . PHP_EOL);
        });

        $http->start();
    }
}

(new App)->run();

```
在代码中我们实现了三个类：

* Container类，用作容器。使用静态属性成员`$instance`存储自身实例化对象，使用对象池数组`$objPool`存储被管理对象。
* App类，用于启动`Swoole`的http服务器，接收请求和回应。
* Test类，即被管理对象的类。

启动swoole服务器，访问链接地址，我们可以看到第一次访问时浏览器输出`usage:520B`的字样。这是实例化容器以及容器实例化Test类所消耗的内存。

当我们刷新页面，可以看到内存消耗为0。这是因为已经实例化的对象交由容器管理起来，未被销毁，在请求需要的时候返回给业务。整个过程不需要再花费开销创建和销毁对象。
## 原型Prototype和单例Singleton
---
to be continued...
