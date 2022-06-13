# 一、项目介绍

## **项目背景说明**

> [【完整手册】第三届字节跳动青训营-后端专场](https://bytedance.feishu.cn/docs/doccnFRB1TXYJPK6yprPETHLXgd)
>
> [极简抖音App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)
>
> [题目一：抖音项目【青训营】](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)

## 项目模块介绍

| 服务名称       | 模块介绍     | 技术框架           | 传输协议   | 注册中心 | 链路追踪        |
| ---------- | -------- | -------------- | ------ | ---- | ----------- |
| api        | API服务    | Gorm、Kitex、Gin | http   | etcd | opentracing |
| userbase   | 用户登录注册管理 | Gorm、Kitex     | thrift | etcd | opentracing |
| useraction | 用户行为数据管理 | Gorm、Kitex     | thrift | etcd | opentracing |
| video      | 视频数据管理   | Gorm、Kitex     | thrift | etcd | opentracing |

## 项目服务调用关系

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/e080ae4398fb444eaf622ae2b5452c25~tplv-k3u1fbpfcp-zoom-1.image)

## 项目模块功能介绍

![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/73a5d6a7875240bcb6818dec807308c9~tplv-k3u1fbpfcp-zoom-1.image)

## 项目技术栈
![](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/04de7fb56ec74be294f877ef0d61771b~tplv-k3u1fbpfcp-zoom-1.image)

# 二、项目运行
## 基础环境
- Golang
- Linux
- Docker
- Mysql
## 运行基础依赖

```
docker-compose up
```

执行上述命令启动Etcd、Jaeger 的 docker 镜像

## 运行 userbase 服务

```
cd cmd/userbase
sh build.sh 
sh output/bootstrap.sh
```

## 运行 useraction 服务

```
cd cmd/video
sh build.sh 
sh output/bootstrap.sh
```
## 运行 video 服务

```
cd cmd/useraction
sh build.sh 
sh output/bootstrap.sh
```
## 运行 api 服务

```
cd cmd/api 
chmod +x run.sh 
./run.sh
```
## 简易http资源服务器
```
cd src
python -m http.server 8001
```