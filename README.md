## 一、项目

基于go-zero v1.3.0开发

相关组件：

- MySQL
- Redis
- Etcd
- Prometheus 
- Grafana 
- Jaeger
- DTM

## 二、服务

- user
    - api服务，端口 8000:8000
    - rpc服务，端口 9000:9000
- product
    - api服务，端口 8001:8001
    - rpc服务，端口 9001:9001
- order
    - api服务，端口 8002:8002
    - rpc服务，端口 9002:9002

## goctl

### 1. 安装goctl

```
# Go 1.15 及之前版本
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/zeromicro/go-zero/tools/goctl@latest

# Go 1.16 及以后版本
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest
```

### 2. 使用goctl生成代码

- 生成model代码

```
goctl model mysql ddl -src ./model/user.sql -dir ./model -c
```

- 生成api代码

```
goctl api go -api ./api/user.api -dir ./api
```

- 生成rpc代码

```
goctl rpc proto -src ./rpc/user.proto -dir ./rpc
//goctl rpc protoc order.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## 二、系统运行

- 运行docker

```
cd deploy
docker-compose up -d
```

- 运行服务

```
cd service/user/rpc
go run user.go
其他服务类似
```

### 文件
```
file文件夹下有mysql.sql和postman文件
```

### Prometheus

```
http://127.0.0.1:3000

宿主机访问Docker的Prometheus服务的IP：
docker inspect {prometheus_container_id} 中的 IPAddress
```

### Jaeger

```
http://127.0.0.1:5000
```

### DTM

```
本项目采用DTM SAGA协议的分布式事务
```

### Reference
- https://juejin.cn/post/7051205679217901599
- https://dtm.pub/ref/gozero.html