借鉴项目pear admin golang编写


#### 项目安装

```bash
# 下 载
git clone https://gitee.com/pear-admin/pear-admin-golang

# 修 改 配 置
cp config.toml.example config.toml


```

#### 运行项目

```bash
go mod tidy
go run main.go
```
#### 备注: 因数据库使用sqlite3，使用及编译时需要gcc。如果不希望额外配置gcc，可以使用发行版。
