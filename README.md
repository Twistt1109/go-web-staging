# go-web-staging
集成gin, sqlx, go-redis, zap, viper的go web开发脚手架， 方便快速部署
使用雪花算法生成id是为了防止注册用户可以推断出网站多少人， 也是为了防止分库分表后id重复

## 项目结构
project/
|-- cmd
|   |-- main.go
|-- config/
|   |-- app.yml
|-- .env

如果开发环境需要使用真实地址, 又不想上传到仓库, 可以在`.env`中配置, 复制`.env.example`文件, 修改名称为`.env`

参考[golang-standards](https://github.com/golang-standards/project-layout/blob/master/README_zh-CN.md)

## Go目录

### `/cmd`
项目主要的应用程序。
对于每个应用程序来说这个目录的名字应该和项目可执行文件的名字相匹配（例如，`/cmd/myapp`）。

### `/internal`
私有的应用程序代码库。这些是不希望被其他人导入的代码。请注意：这种模式是Go编译器强制执行的。
`/internal/app`下面放一下系统启动时需要的配置(如, `consul`, `sentry`)

### `/pkg`
外部应用程序可以使用的库代码（如，`/pkg/mypubliclib`）。其他项目将会导入这些库来保证项目可以正常运行，所以在将代码放在这里前，一定要三思而行。

## 服务端应用程序的目录

### `/api`

OpenAPI/Swagger规范，JSON模式文件，协议定义文件。

更多样例查看[`/api`](https://github.com/golang-standards/project-layout/blob/master/api/README.md)目录。

## Web应用程序的目录

### `/web`

Web应用程序特定的组件：静态Web资源，服务器端模板和单页应用（Single-Page App，SPA）。

## 通用应用程序的目录

### `/configs`
配置文件模板或默认配置。
将`confd`或者`consul-template`文件放在这里。

### `/build`
打包和持续集成。

将云（AMI），容器（Docker），操作系统（deb，rpm，pkg）软件包配置和脚本放在`/build/package`目录中。

将CI（travis、circle、drone）配置文件和就脚本放在`build/ci`目录中。请注意，有一些CI工具（如，travis CI）对于配置文件的位置有严格的要求。尝试将配置文件放在`/build/ci`目录，然后链接到CI工具想要的位置。

### `/deployments`

IaaS，PaaS，系统和容器编排部署配置和模板（docker-compose，kubernetes/helm，mesos，terraform，bosh）。请注意，在某些存储库中（尤其是使用kubernetes部署的应用程序），该目录的名字是`/deploy`。

### `/test`

外部测试应用程序和测试数据。随时根据需要构建`/test`目录。对于较大的项目，有一个数据子目录更好一些。例如，如果需要Go忽略目录中的内容，则可以使用`/test/data`或`/test/testdata`这样的目录名字。请注意，Go还将忽略以“`.`”或“`_`”开头的目录或文件，因此可以更具灵活性的来命名测试数据目录。

更多样例查看[`/test`](https://github.com/golang-standards/project-layout/blob/master/test/README.md)。