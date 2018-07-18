## 基于gin框架的新闻restful-api服务


---

### 项目试用

```shell
go get github.com/xiaofang-git/gin-restful
```
- vender
```
govender fetch github.com/xiaofang-git/gin-restful
```

### 开发模块
    - JWT用户验证模块
        - token生成
        - token校核

    - API服务模块

        - 文章列表根据页数查询api
        - 文章根据ID查询api

    - 权限模块

        - 基于用户角色（分组）鉴权


