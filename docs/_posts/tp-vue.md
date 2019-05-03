---
title: 基于php后端渲染的thinkphp-vue
date: 2019-05-04 00:40:51
tags:
  - vuejs
  - thinkphp
category: php
---
[[toc]]

![Thinkphp logo](https://box.kancloud.cn/5a0aaa69a5ff42657b5c4715f3d49221) 
<img width="100" src="https://vuejs.org/images/logo.png" alt="Vue logo">
<img width="100" src="https://twig.symfony.com/images/logo.png" alt="Twig logo">
# tp-vue
## 介绍
基于**php后端渲染**的thinkphp+vue框架。

采用**twig**模板引擎渲染。

框架版本为thinkphp5.1和vue@2.6.10。

php版本要求7.0+。
## 快速开始
```bash
git clone git@github.com:yuchanns/tp-vue.git
cd tp-vue
composer install
php think run
```
![](https://i.imgur.com/IXRmzhL.jpg)
![](https://i.imgur.com/jetkVOT.jpg)
## 使用方法
### 继承模板
你的主页面需要使用继承的方式进行编写：
```twig
{% extends library('base/base') %}

{% macro template() %}
  <div>
    <!-- 此处编写你的html -->
  </div>
{% endmacro %}

{% block default %}
  {% from _self import template %}
  <script>
    const App = {
      // 此处编写你的js
      template: '${tiny(template())}'
    }
  </script>
{% endblock %}

{% block style %}
  <style>
    /* 此处编写你的style */
  </style>
{% endblock %}
```
### 编写组件
使用twig模板引擎的宏功能编写你的组件：
```twig
{% macro template() %}
  <div>
    <!-- 此处编写你的组件html -->
  </div>
{% endmacro %}

<script>
  export default {
    {% macro default() %}
    {% from _self import template %}
    
    // 此处编写你的组件js
    template: '${tiny(template())}'
    {% endmacro %}
  }
</script>
{% macro style() %}
  <style>
    /* 此处编写你的组件style */
  </style>
{% endmacro %}
```
### 调用组件
使用twig模板引擎的import功能引入组件和组件样式：
```twig
{% extends library('base/base') %}

{% macro template() %}
  <div>
    <p>在下面使用组件</p>
    <HelloWorld msg="Welcome to Your Vue.js App"></HelloWorld>
  </div>
{% endmacro %}

{% block default %}
  {% from _self import template %}
  <!-- 在下面引入组件的js -->
  {% from library('index/components/HelloWorld') import default as HelloWorld %}
  <script>
    const App = {
      components: {
        // 在下面注册组件，使用tinyj模板函数压缩代码
        HelloWorld: {${tinyj(HelloWorld())}}
      },
      template: '${tiny(template())}'
    }
  </script>
{% endblock %}

{% block style %}
  <!-- 在下面引入组件的样式并使用 -->
  {% from library('index/components/HelloWorld') import style as HelloWorld %}
  ${HelloWorld()}
  <style>
    #app {
      font-family: 'Avenir', Helvetica, Arial, sans-serif;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
      text-align: center;
      color: #2c3e50;
      margin-top: 60px;
    }
  </style>
{% endblock %}
```
*注意：引用组件需要在块或宏的作用域内进行才有效果！*

支持组件嵌套引用组件。
### 模板函数
* `${tiny()}`：压缩html和css代码
* `${tinyj()}`：压缩js代码
* `${static_url()}`：引用静态资源
* `${library()}`：获取模板文件/组件路径

*可在`./extend/twig/Functions.php`中自行添加自定义模板函数*
