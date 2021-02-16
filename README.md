# 360 Quake API / Fofa API

![](https://img.shields.io/badge/Version-0.1-blue.svg)
![](https://img.shields.io/badge/Golang-1.15-blue.svg)

## 项目介绍

对于渗透测试工程师来说，一定熟悉Fofa，但大部分人应该不了解Quake

其实360公司的Quake是与Fofa、Shodan类似的搜索引擎，而且效果更好，全名是360网络空间测绘系统

该项目起初是为了做QuakeAPI，后续发现实际工作中也有对FofaAPI的需求，所以最终打算兼容FofaAPI

具体如何使用可以参考下文和图片内容

- Quake：https://quake.360.cn/
- Fofa：https://fofa.so/

为什么选择Golang：

- 直接生成多个平台可执行文件，无需安装各种复杂的环境（Python，JDK等等）
- Golang拥有C++的性能和Python的简洁
- 近年来安全工具的主流语言（Xray,Goby,Kunpeng...）

## 快速上手

在github的release页面下载可执行文件：[下载地址](https://github.com/EmYiQing/QuakeAPI/releases)

下面以MySQL为例：

- Windows

```bash
QuakeAPI.exe --fofa --key [your_key] --email [your_email] --search "mysql" --total 1000 --output result.txt
```

```bash
QuakeAPI.exe --quake --key [your_key] --search service:mysql --total 1000 --output result.txt
```

- Linux / MacOS

```bash
./QuakeAPI --fofa --key [your_key] --email [your_email] --search "mysql" --total 1000 --output result.txt
```

```bash
./QuakeAPI --quake --key [your_key] --search service:mysql --total 1000 --output result.txt
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
- --email：当使用fofa时需要提供email
- --fofa：使用fofa引擎
- --quake：使用quake引擎
- --help：查看帮助
- --key：输入你的API Key
- --search：输入你的查询字符串
- --output：设置输出文件（默认是result.txt）
- --total：查询总条数（默认是100条，建议设置为100的倍数。如果数量较多，自动使用协程）
- --userinfo：是否查询用户相关的信息（用户名邮箱等信息）