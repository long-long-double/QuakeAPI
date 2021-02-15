# 360网络空间测绘系统（Quake）API

- Quake与Fofa、Shodan类似的搜索引擎，功能更强大
- Golang编写，生成EXE，无需安装各种复杂的环境（Python，JDK等等）
- 使用高并发技术，获取大量数据时的效率提升
- 代码相对完善，考虑到后续拓展性

## 参数说明

- --help：查看帮助
- --key：输入你的API Key
- --search：输入你的查询字符串
- --output：设置输出文件（默认是result.txt）
- --total：查询总条数（默认是100条，如果数量较多，自动使用协程）
- --userinfo：是否查询用户相关的信息

## 图片

### 使用说明
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/quake0.png)

### 实际使用
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/quake1.png)

### 查询结果
![](https://xuyiqing-1257927651.cos.ap-beijing.myqcloud.com/quake/quake2.png)

# 360 Cyberspace Surveying And Mapping System (Quake) API

-Like fofa and Shodan, Quake has more powerful functions
-Golang write, generate exe, no need to install a variety of complex environment (Python, JDK, etc.)
-Use high concurrency technology to improve the efficiency of obtaining large amounts of data
-The code is relatively perfect, considering the subsequent expansibility

## Parameter description

- -- help: view help
- -- key: enter your API key
- -- Search: enter your query string
- -- output: set the output file (default is result.txt ）
- -- total: the total number of queries (the default is 100, if the number is large, the goroutine will be used automatically)
- -- userinfo: query user related information