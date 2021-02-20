# 360 Quake API / Fofa API

![](https://img.shields.io/badge/Version-0.1-blue.svg)
![](https://img.shields.io/badge/Golang-1.15-blue.svg)

## 项目介绍

360公司的Quake是与Fofa、Shodan类似的搜索引擎，而且效果更好，全名是360网络空间测绘系统

该项目起初是为了做QuakeAPI，后续发现实际工作中也有对FofaAPI的需求，所以最终打算兼容FofaAPI

具体如何使用可以参考下文和图片内容

- Quake：https://quake.360.cn/
- Fofa：https://fofa.so/

简介：

- 基于命令行，支持QuakeAPI和FofaAPI的查询
- 使用协程技术提高大量数据查询的效率
- 支持纯命令行方式和YAML配置文件的方式
- 支持普通的TXT输出和保存到MySQL的方式

为什么选择Golang：

- 直接生成多个平台可执行文件，无需安装各种复杂的环境（Python，JDK等等）
- Golang拥有C++的性能和Python的简洁
- 近年来安全工具的主流语言（Xray,Goby,Kunpeng...）

## 快速上手

在github的release页面下载可执行文件：[下载地址](https://github.com/EmYiQing/QuakeAPI/releases)

下面以MySQL为例：

- Windows

```bash
QuakeAPI.exe --fofa --key [your_key] --email [your_email] --search mysql --total 1000 --output result.txt
```

```bash
QuakeAPI.exe --quake --key [your_key] --search service:mysql --total 1000 --output result.txt
```

- Linux / MacOS

```bash
./QuakeAPI --fofa --key [your_key] --email [your_email] --search mysql --total 1000 --output result.txt
```

```bash
./QuakeAPI --quake --key [your_key] --search service:mysql --total 1000 --output result.txt
```

- 如果命令行太麻烦，可以直接使用配置文件：

使用`--config`参数创建配置文件

```bash
QuakeAPI.exe --config
```

当使用配置文件的方式时，支持保存到MySQL，设置好相关的参数即可

```yaml
login:
  email: "your@email.com"       # 如果使用Fofa需要提供Email
  key: "your-key"               # 需要提供API Key
  userinfo: false               # 是否查询用户相关的信息

search:
  query: "service:http"         # 查询字符串
  output: "result.txt"          # 输出文件（使用mysql将忽略此选项）
  total: 1000                   # 查询个数（建议100的倍数）

use:
  quake: false                  # 是否使用quake引擎
  fofa: true                    # 是否使用fofa引擎
  
mysql:
  use: false                    # 如果需要保存MySQL设置为True
  server: localhost             # 服务器IP
  port: 3306                    # 端口
  username: your-username       # 用户名
  password: your-password       # 密码

```

编辑完配置文件后，再次运行相同的命令执行程序

```bash
QuakeAPI.exe --config
```

## 截图

### 说明
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/quake-0.png)

### Fofa
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/fofa-1.png)

### Quake
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/quake-1.png)

### 结果
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/result.png)

## 参数说明：
- --config：使用配置文件模式
- --email：当使用fofa时需要提供email
- --fofa：使用fofa引擎
- --quake：使用quake引擎
- --help：查看帮助
- --key：输入你的API Key
- --search：输入你的查询字符串
- --output：设置输出文件（默认是result.txt）
- --total：查询总条数（默认是100条，建议设置为100的倍数。如果数量较多，自动使用协程）
- --userinfo：是否查询用户相关的信息（用户名邮箱等信息）