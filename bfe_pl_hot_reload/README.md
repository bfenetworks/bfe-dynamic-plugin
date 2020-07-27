# bfe_pl_logid

### Describe
用于动态加载配置文件

### Build
```
go build -buildmode=plugin -o bfe_pl_hot_reload.so bfe_pl_hot_reload.go
```

### How to use
编辑主配置文件 conf/bfe.conf

```
Plugins = $your-dynamic-plugins-path/bfe_pl_hot_reload.so
```