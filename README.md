# gin-mvc-framework
根据其他语言的个人经验基于gin搭建的简单框架结构,仅用于学习参考。

# 目录结构
- app 项目目录
    -   constants 常量
    -   entity 实体层
    -   http http相关
        -   controller 控制器
        -   middleware 中间件
    -   service 服务层
    -   utils 工具
- config 配置文件
    -  database 数据库配置
    -  router 路由配置
- main.go 启动文件

# 项目启动


```
go run main.go
```
访问locahost:8088 出现返回结果

```
{"code":200,"message":"hello","data":null}
```
