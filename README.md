# 切图服务
WEB服务对服务器图片进行切割处理，按指定将大图切成若干份小图。

# 安装
    go get github.com/ysqi/splitimage
    cd $GOPATH/github.com/ysqi/splitimage
    go build

# 配置
配置文件在文件夹conf下 app.conf

```ini
    # 服务运行端口，默认为8080
    httpport = 8080 
    runmode = prod
    # 缩略图大小控制，为0或负数时表示无限制
    thubMaxLen = 2000
```

# 运行服务
在Liunx下运行：`$ ./splitimage` ,在 Windown 下双击运行` spliteimage.exe ` 即可。

**注意**
因为服务需要访问服务器图片资源，请确保服务有权限访问图片资源

# 服务调用
本机测试，浏览器访问 http://localhost:8080/image/split?src=XXX&length=XXX&save=XXX

**参数说明**
+ src : 待切片图片路径
+ length： 切后图片最大高宽
+ save： 切图保存位置

**实例**
>http://localhost:8080/image/split?src=C:\new.png&length=400&save=C:\NEW\
 
切片成功返回消息： 

```json
{
    "success":true,
    "msg":"",
    "data":{
        "x":2,
        "y":3
    }
}
```
如果切片失败则返回：

```json
{
    "success":false,
    "msg":"文件不存在"
}
