# mylinkserver-main

这是一个连连看小游戏的服务器

里面包含几个模块：

1.登陆注册

2.动态和评论

3.网络对战

4.排行榜

5.音乐获取

这几个模块很简单，代码也不难，但是没有时间优化代码结构，很多东西封装也没有做好。

启动程序请打开终端在根目录下，输入：go run server/cmd/main.go

如果上传服务器后，服务器由于代理问题，无法成功download依赖，则可以查看GOPROXY官方指示——https://goproxy.io/zh/docs/getting-started.html

使用云服务器，而不用宝塔等面板软件，可以用nohup后台运行程序，将其日志输出重定向至指定文件。
