本项目采用 gin+gorm框架实现简单的用户注册、用户认证、文章发布、修改、查询、评论创建、查询等功能


入口函数为main.go

data目录下放的sqlite数据文件

middleware目录中放了jwt解析、权限校验、全局异常捕获的处理逻辑

model目录下存在user、post、comment的数据库操作方法

router包下配置的请求地址与路由映射

controller包下放的路由处理方法

tests包下放的单元测试用例
