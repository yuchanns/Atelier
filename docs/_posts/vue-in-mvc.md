---
title: vue在传统mvc框架下的使用
date: 2019-03-17 16:41:27
top: 2
tags:
  - vuejs
  - aiohttp
category: python
---
[[toc]]

当我们手上有前后端不分离的传统mvc框架下的项目，又想要摆脱繁琐的Dom操作地狱，就可以使用[vuejs](https://cn.vuejs.org)。在github上我看到很多前后端不分离的项目源码中对于`vuejs`的整合都是直接在项目根目录放置一个frontend文件夹，内置`vue-cli`生成的开发脚手架，在`webpack`下以`es6`语法的形式编写，再将编译生成的`dist`页面放到后端mvc的view中。

有时候（当编写小项目的时候）我觉得这样也显得过于复杂和巨大，幸好`vuejs`支持直接用`<script>`引入使用。本文记录的便是我在这种背景下，利用模板引擎的特性对`vuejs`的使用心得。
<!-- more -->
## 准备工作
本文所使用的web框架基于`python3.7`，使用到的库包括：
* aiohttp
* aiohttp-jinja2
* uvloop

你也可以随便选择一个web框架，无论是`django`或者`flask`或者`tornado`，甚至是`php`也没问题，只需要模板引擎支持`模板继承`以及`导入上下文`的功能就行。本文只是出于搭设便捷的目的做出选择。
### 安装环境
使用`miniconda3`创建一个虚拟环境并安装库：
```shell
conda create -n aiohttp python=3.7
source activate aiohttp
(aiohttp) pip install aiohttp aiohttp-jinja2 uvloop
```
当然`aiohttp-jinja2`还依赖于`jinja2`这个库，不过这些关联依赖都会由pip自行解决，用户不必关心细节。
### framework结构
这里仅使用典型性而简洁的mvc框架结构：
```shell
.
├── handlers
│   ├── __init__.py
│   └── index.py
├── main.py
└── templates
    └── base
        └── base.html
```
下面给出各文件源码：
```python
# main.py

import uvloop
from aiohttp_jinja2 import setup
from jinja2 import FileSystemLoader
from aiohttp import web
from handlers.index import routes


def make_app():
    app = web.Application()
    # 注意！为了避免vue和jinja2的Mustache语法冲突
    # 我将jinja2的模板变量改成{$variable}的形式 
    setup(app, loader=FileSystemLoader('./templates'),
          variable_start_string='{$', variable_end_string='}')
    app.add_routes(routes)
    return app

try:
    uvloop.install()
    web.run_app(make_app(), host='127.0.0.1', port=9000)
except KeyboardInterrupt:
    print('see you!')

```
```python
# handlers/index.py

from aiohttp_jinja2 import template
from aiohttp import web


routes = web.RouteTableDef()


@routes.get('/')
@template('base/base.html')
async def index(request):
    return

```
```html
<!-- templates/base/base.html -->

<!DOCTYPE>
<html>
  <head>
    <title>aiohttp-vue</title>
  </head>
  <body>
    <div id="app">
      <p>{{ greeting }}</p>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/vue@2.6.9/dist/vue.js">
    </script>
    <script>
      const app = new Vue({
        data () {
          return {
            greeting: 'hello world'
          }
        }
      }).$mount('#app')
    </script>
  </body>
</html>
```
启动aiohttp服务器，访问 http://127.0.0.1:9000/ 可正常见到`hello world`
## 模板继承和vue组件
在前端框架下我们编写代码通常是以组件的形式：
```vue
<template>
 <div>
 <p>{{ greeting }}</p>
 </div>
</template>

<script>
export default {
  data () {
    return {
      greeting: 'hello world'
    }
  }
}
</script>
```
在`jinja2`模板中，我们可以使用`{% block title %}{% endblock %}`这一模板继承功能达到相似写法。
### 文件变动
首先在templates文件夹下创建一个index文件夹，在该文件夹下创建一个index.html文件。然后我们将index.py中，`@template('base/base.html')`改成`@template('index/index.html')`
### 代码变动
然后修改templates/base/base.html，加入组件（[对组件不了解？](https://cn.vuejs.org/v2/guide/components.html)）。
```diff
<!-- templates/base/base.html -->

<!DOCTYPE>
<html>
  <head>
    <title>aiohttp-vue</title>
  </head>
  <body>
    <div id="app">
-     <p>{{ greeting }}</p>
+     <main-content></main-content>
    </div>
    <script src="https://cdn.jsdelivr.net/npm/vue@2.6.9/dist/vue.js">
    </script>
+   {% block script %}
+   <script>
+     let mainContent = {}
+   </script>
+   {% endblock %}
    <script>
+     mainContent.template = `<div>{% block template %}{% endblock %}</div>`
+     Vue.component('main-content', mainContent)
-     const app = new Vue({
+     const app = new Vue().$mount('#app')
-       data () {
-         return {
-           greeting: 'hello world'
-         }
-       }
-     }).$mount('#app')
    </script>
  </body>
</html>
```
接着编写templates/index/index.html
```html
<!-- templates/index/index.html -->

{% extends "base/base.html" %}

{% block template %}
<p>{{ greeting }}</p>
{% endblock %}

{% block script %}
<script>
let mainContent = {
  data () {
    return {
      greeting: 'hello world'
    }
  }
}
</script>
{% endblock %}
```
重新启动aiohttp服务器，再次访问同一网址，可以得到不变的结果。
### 这样写有什么好处 ？
由于我们嫌弃在前后不分离的小项目中使用纯前端框架过于复杂和巨大而选择导入script，所以我们也使用不到vuejs的单文件组件服，转而使用字符串模板定义组件。这样做的缺点是：html模板缺乏**语法高亮**，以及另外几个[缺点](https://cn.vuejs.org/v2/guide/single-file-components.html#%E4%BB%8B%E7%BB%8D)。

当我们像上述方式利用模板引擎的模板继承功能，就获得了**完整语法高亮**。

并且我们也可以在`Vue DevTools`完全观察到整个html页面的组件结构。

![Vue DevTools](https://i.imgur.com/vegX2ny.jpg)

只需要写一次base.html，任何继承base模块的模板都将自动成为其组件，而我们只需要放心地在模块里享受语法高亮的服务，写自己的业务代码。

遗憾的是我们依旧面临着**全局作用域的CSS**（即使你在`style`标签里添加`scoped`的属性，也只有`firefox`支持）、**每个component中的命名不得重复**的问题。

当然，我们在开头就说过，之所以采用script引入而不使用前端框架编译开发的原因是目标项目小而简单，因此全局css和组件命名不得重复并不是什么太大的问题。
## 对组件命名重复的解决方案
接下来这部分是针对模板引擎带有**宏定义**功能的组件命名重复解决方式。

结合`jinja2`的`macro`功能和上下文引入可以如此使用。
### 文件变动
首先在templates/index文件夹下创建一个components文件夹，然后在该文件夹下面创建一个test.html文件。
### 代码变动
我们编写一个名为`test`的组件：
```html
<!-- templates/index/components/test.html -->

{% macro template() -%}
<div>
    <p>{{ title }}</p>
</div>
{%- endmacro %}

<script>
    const test = // {% macro script() -%}
    {
        data () {
            return {
                title: 'test'
            }
        },
        template: `{$ template() }`
    }
    // {%- endmacro %}
</script>
```
对index.html做出如下修改：
```diff
<!-- templates/index/index.html -->

{% extends "base/base.html" %}

{% block template %}
<p>{{ greeting }}</p>
+ <test></test>
{% endblock %}

{% block script %}
+ {% from 'index/components/test.html' import script %}
<script>
  let mainContent = {
+   components: {},
    data () {
      return {
        greeting: 'hello world'
      }
    }
  }
+ mainContent.components.test = {$script()}
</script>
{% endblock %}
```
再次访问网址，可以看到新添加的组件生效了。

此时 我们在templates/index/components文件夹下再创建一个test2.html文件，内容则和test.html一模一样。
然后对index.html进行修改：
```diff
{% extends "base/base.html" %}

{% block template %}
<p>{{ greeting }}</p>
<test></test>
+ <test2></test2>
{% endblock %}

{% block script %}
{% from 'index/components/test.html' import script %}
+ {% from 'index/components/test2.html' import script as script2 %}
<script>
  let mainContent = {
    components: {},
    data () {
      return {
        greeting: 'hello world'
      }
    }
  }
  mainContent.components.test = {$script()}
+ mainContent.components.test2 = {$script2()}
</script>
{% endblock %}
```
访问网址，确认正常。
### 为什么要使用组件？
事实上，我最开始这样做的时候，采用的是`Vue.extend(mainContent)`的形式。但是`Vuejs`的组件的一个好优势就是组件内的一切只在组件当中起作用，我们就不用担心各种名字重复的问题。

上述所说对组件命名重复的解决方案，实际上是通过`宏`功能为组件命名来规避名字重复（如同引用单文件组件的`import test as test2 from './components/test'`）。解决的方法并不是特别方便（编写过程有感知），幸好写起来并不麻烦。

是否使用此方法以解决组件名称重复？无论哪种选择，只要在编写代码的过程中注意相应的问题，我们就可以畅快地在前后端不分离的传统mvc项目中使用`vuejs`了！
## 总结
文中我提到了两种组件的写法，一种**模块继承**，一种**宏定义**。我建议两种写法结合使用。

**模块继承**用于继承布局之后书写主体业务代码使用，和常规的模板继承写法一致，实际渲染起来则是作为`Vue`的文章组件;**宏定义**则作为在文章组件中使用组件的解决方案。当然，并不是说没有宏定义就不能在组件中使用组件了，只不过这样你将失去html代码高亮功能，转而只能书写丑陋的字符串模板，未免令人有些不满～

## 本文源代码
https://github.com/yuchanns/vue-in-mvc
