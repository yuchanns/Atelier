---
title: MySQL学习笔记01
date: 2019-01-13 14:36:22
tags:
- mysql
category: 学习
---
<!-- more -->
[[toc]]
## 1.sql查询语句的执行过程
**学习来源** [极客时间-MySQL实战45讲01](https://time.geekbang.org/column/article/68319)

#### MySQL逻辑架构图

@flowstart 
st=>start: 客户端
connect=>condition: 连接器：权限验证
是否开启缓存？
cache=>condition: 查询缓存
是否命中？
analyse=>operation: 分析器：词法、语法分析
optimize=>operation: 优化器：索引选择
execute=>operation: 执行器：操作引擎
result=>end: 返回结果
st->connect(yes)->cache
connect(no)->analyse
cache(yes)->result
cache(no)->analyse
analyse(right)->optimize
optimize->execute
ana=>operation: test
execute(left)->result
@flowend

MySQL 可以分为`Server`层和`存储引擎`层两部分。
连接器到执行器部分为`Server`层。
`show processlist;`可以查看连接。
#### 解决长连接占用内存过高问题：
* 定期断开长连接。使用一段时间或者执行一个占用内存的大查询之后断开连接
* MySQL5.7+通过执行`mysql_reset_connection`初始化连接资源

#### 查询缓存利弊
对更新操作频繁的表，查询缓存容易失效，不适合使用。
将`query_cache_type`设置成DEMAND，默认不使用。使用`SQL_CACHE`显式指定：
```sql
select SQL_CACHE * from table where column=1;
```
MySQL8.0+删除了查询缓存功能。
#### 权限验证
在执行器执行和命中缓存返回结果之前会先进行用户对表的权限验证。

#### 慢日志查询
在慢日志中`rows_examined`字段表明该sql语句执行过程中扫描了多少行。

## 2.sql更新语句的执行过程
**学习来源** [极客时间-MySQL实战45讲02](https://time.geekbang.org/column/article/68633)
更新流程涉及到两个重要的日志模块：`redo log`和`binlog`。
#### WAL技术-InnoDB的`redo log`
全称`Write-Ahead Logging`。
先写日志，再写磁盘。
当有记录更新时，`InnoDB`引擎先把记录写到`redo log`中，并更新内存，就算更新完成。然后在合适的时候记录到磁盘中。
`redo log`是大小固定，从头开始写，写到末尾又回到开头循环。
保证`crash safe`能力
#### Server层的日志-`binlog`
区别：
* `redo log`是`InnoDB`独有的
* `redo log`是物理日志，`binlog`是逻辑日志
* `redo log`循环写，固定大小，`binlog`追加写，不会覆盖

#### 更新流程图
执行更新语句
```sql
update T set c=c+1 where ID=2;
```

@flowstart 
fetch=>start: 取ID=2这一行
（执行器执行）
cache=>condition: 数据页是否在内存中？
（InnoDB内部）
ioread=>operation: 磁盘中读取到内存
（InnoDB内部）
returnrow=>operation: 返回行数据
（InnoDB内部）
increase=>operation: 将这行c值加1
（执行器执行）
write=>operation: 写入新行
（执行器执行）
updatecache=>operation: 更新到内存
（InnoDB内部）
prepare=>operation: 写入redo log
处于prepare阶段
（InnoDB内部）
w_binlog=>operation: 写入binlog
（执行器执行）
commit=>end: 提交事务
处于commit状态
（InnoDB内部）
fetch->cache(no)->ioread
cache(yes)->returnrow
ioread->returnrow->increase(right)->write->updatecache->prepare->w_binlog->commit
@flowend

从`prepare`到`commit`叫做`两阶段提交`，目的是让`redo log`和`binlog`状态保持一致。
#### 保证数据不丢失
在`my.cnf`中：
`innodb_flush_log_at_trx_commit`设置成1
`sync_binlog`设置成1