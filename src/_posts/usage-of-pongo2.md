---
title: pongo2使用笔记
date: 2020-02-07 22:16:00
category: golang
tags:
  - chore
---
:::warning 草稿
施工中...
:::
<mark>pongo2</mark>[^1]是一个`Django`风格的模板引擎，目标是模板语法完全适配Django。
<!-- more -->

[[toc]]

## 简介
虽然传统的**MVC/MTV**大势已去，日益复杂的前端技术致使如今的**Trending**是前后分离，后端模板引擎不再受到开发者的重视，但是依然有开源库还提供这一功能。

甚至gin本身也自带了一个简陋的模板引擎，虽然正式开发中基本上不会用到这个功能，但是笔者还是想要了解一下相关的事情。然而在稍微阅读了gin的文档和源码之后，稍微写一个小demo时，笔者发现gin的模板引擎过于简陋，很是不爽。于是决定找找第三方的模板引擎。

在笔者有限的经历中，用起来最强大和舒爽的模板引擎当属`Jinja2`——这是一个用python实现的著名的模板引擎，具有诸如`Macro`、`Block`、`Include`等强大功能，同时性能也很好。在[pkg.go.dev](https://pkg.go.dev/)中，笔者轻易找到了一个自称**A Django-syntax like template-engine**的模板引擎，它就是[pongo2](https://pkg.go.dev/github.com/flosch/pongo2)。这里提示一下不了解python的读者，Jinja2和Django模板引擎的语法基本一致。

pongo2的文档很简单，仅提供了Api的说明，作者表示关于模板的用法，只需要查看Django的文档就可以了[^2]。

## 接入gin
笔者并不打算浪费篇幅描述关于模板语法的内容，这些在Django文档上可以清楚地了解到。本文的重点只是研究一下如何将pongo2和gin结合起来使用。
### gin的模板引擎接口分析
### 用pongo2实现接口
```go
package main

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
	"path"
)

type PongoRender struct {
	TmplDir string
}

func New(tmplDir string) *PongoRender {
	return &PongoRender{
		TmplDir: tmplDir,
	}
}

func (p *PongoRender) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	fileName := path.Join(p.TmplDir, name)

	if gin.Mode() == gin.DebugMode {
		template = pongo2.Must(pongo2.FromFile(fileName))
	} else {
		template = pongo2.Must(pongo2.FromCache(fileName))
	}

	return &PongoHTML{
		Template: template,
		Name:     name,
		Data:     data.(pongo2.Context),
	}
}

type PongoHTML struct {
	Template *pongo2.Template
	Name     string
	Data     pongo2.Context
}

func (p *PongoHTML) Render(w http.ResponseWriter) error {
	p.WriteContentType(w)
	return p.Template.ExecuteWriter(p.Data, w)
}

func (p *PongoHTML) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}
```

[^1]: [flosch/pongo2](https://github.com/flosch/pongo2)
[^2]: [Django 文档](https://docs.djangoproject.com/zh-hans/3.0/)