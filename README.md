一个简易的使用Gin框架开发的HTTP服务

config目录：解析config.ini配置文件
db：gorm操作mysql数据库
log:日志模块封装
model：数据模型，封装了数据库的增删查改等操作
router：http路由
tools：一些常用的工具类
vendor:依赖第三方包
controller：api具体逻辑
config.ini:配置文件，自行添加即可
main.go：程序入口

运行的时候，首选启动自己的MySQL数据库，修改对应的config.ini文件，执行
go run main.go即可
