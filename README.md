# blog-sync
这是一个自动拉取博客到机器上的一个服务，提供了一个 http请求服务，只要请求这个地址，就会自动拉取git更新。

## 编译
在项目当前目录下执行
```sh
go build 
```
会生成可执行文件 `blog-sync`

## 运行
```
nohup ./blog-sync  2>&1 &
```
默认启动端口为 8080   ， 请求地址 `http://x.x.x.x:8080/sync`  
```json
{
    "message": "repo sync OK"
}
``` 
