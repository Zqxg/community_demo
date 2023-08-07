# 社区话题demo
**此项目使用`Gin`+`Gorm`+`mysql` ，基于`RESTful API`实现的一个社区话题评论demo。

## 需求描述
1. 支持发布帖子（增）
2. 支持删除帖子（删）
3. 帖子话题运行修改（改）
4. 展示话题（标题、文字描述）和回帖列表（查）

+ 话题和回帖数据使用mysql数据库

## 项目功能介绍
- 新增/删除/修改/查询 话题
- 新增/删除/修改/查询 评论

后续添加用户注册登录功能

## 接口文档
[接口文档](https://console-docs.apipost.cn/preview/7cdfdb30d3008cd4/8f13a469d4be6339)

## 项目依赖
GO 1.20.1
1. gin
2. gorm
3. mysql
4. ini

## 项目结构
```
├── api       //api接口：错误信息、话题模块、评论模块
├── cmd       //程序启动
├── conf      //相关配置
├── models    //初始化、结构体
├── routes    //路由逻辑处理
├── serializer//序列化
└── service   //接口函数的实现
```

## 评论模块功能实现
数据库表一对多（唯一外键）
## 开发问题
1. 数据库连接中断
+ 在设置数据库连接池前多写了个defer关闭了数据库，导致在创建方法前关闭了数据库

2. 评论模块数据库操作有问题，无法对应到相关的帖子下。怀疑是数据库操作有问题。
+ 使用string查询条件
```
var posts []models.Post  
err = models.DB.Where("topic_id = ?", tid).Find(&posts).Error
```
## 参考资料
1. [(27条消息) Golang gorm 关联关系 一对多_golang 一对多查询_富士康质检员张全蛋的博客-CSDN博客](https://blog.csdn.net/qq_34556414/article/details/130099428?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-130099428-blog-107849953.235^v38^pc_relevant_sort&spm=1001.2101.3001.4242.2&utm_relevant_index=4)
2. [查询 | GORM - The fantastic ORM library for Golang, aims to be developer friendly.](https://gorm.io/zh_CN/docs/query.html#String-%E6%9D%A1%E4%BB%B6)
3. https://github.com/CocaineCong/TodoList