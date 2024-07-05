# go-web-staging
集成gin, sqlx, go-redis, zap, viper的go web开发脚手架， 方便快速部署
使用雪花算法生成id是为了防止注册用户可以推断出网站多少人， 也是为了防止分库分表后id重复

## 项目结构
```
├── cmd
│   └── main.go
├── config
│   └── config.go
├── internal
│   ├── api
│   │   ├── api.go
│   │   └── user
│   │       ├── api.go
│   │       └── user.go
│   ├── config
│   │   └── config.go
│   ├── controller
│   │   ├── controller.go
│   │   └── user
│   │       ├── controller.go
