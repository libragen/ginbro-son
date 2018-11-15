# Ginbo RESTful APIs

## 安装依赖
	go get -u github.com/gin-contrib/cors
	go get -u github.com/gin-contrib/static
	go get -u github.com/gin-gonic/autotls
	go get -u github.com/gin-gonic/gin
	go get -u github.com/sirupsen/logrus
	go get -u github.com/spf13/viper
    go get -u github.com/go-redis/redis
    go get -u github.com/go-sql-driver/mysql
    go get -u github.com/jinzhu/gorm
    
## 使用
- [swagger DOC ](http://127.0.0.1:5555/swagger/)`http://127.0.0.1:5555/swagger/`
- [static ](http://127.0.0.1:5555)`http://127.0.0.1:5555`
- [app INFO ](http://1127.0.0.1:5555/app/info)`http://127.0.0.1:5555/app/info`
- API baseURL : `http://127.0.0.1:5555/api/v1`

## 注意
- mysql表中没有id/ID/Id/iD字段将不会生成路由和模型
- json字段 在update/create的时候 必须使可以序列号的json字符串,否则mysql会报错