
项目结构
```
project/
    main.go  # 包含服务器的主入口文件
    database # 数据库相关的代码
    routers
        handlers # 路由处理函数
        middlewares # 路由中间件
    config/  # 配置相关的代码
    utils/  # 通用的工具函数和库
    docs/  # 项目文档
    tests/  # 测试用例
```





权限表
用户管理
    账号管理 账号列表 创建账号 删除账号 编辑账号 修改账号 角色分配
    角色管理 角色列表 角色添加 角色删除 角色编辑

如果一个页面 没有单独的接口 就不单独提供code

tip: 白名单功能 可以先不写


### 接口权限/菜单权限 关联性

先思考 什么样的接口需要添加权限 数据类 操作类

假设现在有 A页面
A页面需要调用下面几个接口:
/api/v1/user/list
/api/v1/user/detiail
/api/v1/user/show