

* WebSocket是一种在单个TCP连接上进行全双工通信的协议
* WebSocket使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
* 在WebSocket API中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输
* 需要安装第三方包：
    * cmd中：go get -u -v github.com/gorilla/websocket


聊天室demo运行   go run server.go hub.go data.go connection.go  之后执行local.html文件




gorush：Go 编写的通知推送服务器      https://github.com/appleboy/gorush
gotify：基于 WebSocket 进行实时消息收发的简单服务器      https://github.com/gotify/server



