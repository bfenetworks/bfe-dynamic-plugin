# bfe_pl_logid

### Describe
用于 bfe 日志生成唯一 ID 的动态插件

### Build
```
go build -buildmode=plugin -o bfe_pl_logid.so bfe_pl_logid.go
```

### How to use
编辑主配置文件 conf/bfe.conf

```
Plugins = $your-dynamic-plugins-path/bfe_pl_logid.so
```