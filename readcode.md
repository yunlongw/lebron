# 代码阅读笔记

# API 


# RPC

## 每个 rpc service 可以启动复数个服务，只要端口不同就可以。 
> api 在请求同一个 service 的时候是通过 Etcd 寻找对应的  Key 实现的 ，如果有复数个 service 对应同一个 key ，那么 etcd 会实现自身的负载均衡，返回不同的服务实例

> 两个 user ，加载不同的配置文件，区别只是端口不同 ，都被注册到了 etcd 里面 ，API 端不用作任何更改，而且都不需要重启 ， API 也是从 etcd 拿数据 ， ETCD 就直接负载均衡， 用 POSTMAN 请求 API 接口 ， 两个 USER 就交替返回数据 ， API 完全无感知
```
// 通过加载不同的配置实现
go run user.go  -f etc/user_02.yaml
```

## ETCD 
```
etcdctl get "" --prefix --keys-only  // 查看集群中的所有键
```