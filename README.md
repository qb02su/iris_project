# iris_project
一个基于iris框架的电商后台管理系统

效果如图
![image](https://github.com/qb02su/iris_project/blob/master/static/img/12.png)

#### 项目介绍

- `iris-go` 框架后台接口项目

- `xorm` 数据库模块

- 使用了`mvc`包作为对mvc架构的支持

- 数据支持 `mysql`，`redis` 配置文件在`config.json`文件下

- 前端采用了 `vue` 框架

  

#### 项目初始化

> 拉取项目

```
git clone https://github.com/qb02su/iris_project.git
```

> 加载依赖管理包使用国内镜像。
>
> 官方： https://goproxy.io/
>
> 中国：[https://goproxy.cn](https://goproxy.cn/)
>
> golang 1.13 可以直接执行：

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

> 运行项目

```
go run main.go // go 命令
```

#### 登录项目

输入地址 [http://localhost:8081](http://localhost:8081/)

具体可在在 `config.json` 内配置

