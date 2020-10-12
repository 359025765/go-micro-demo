# go-micro-demo 微服务案例
 
###  服务框架 
- 应用层(php所选框架laravel)，rpc框架则是go-micro

### 通信编码
- php应用与go微服务之间通过json编码，微服务与微服务之间的服务调用则通过protobuf编码 

### 服务网关
- php应用于微服务之间通过Http Api网关进行通信，微服务与为微服务之间通过grpc

### 数据库存储 
- mysql

### 服务发现 
- Etcd

### 开发环境
- macOS

### 部署
- docker
- k8s
