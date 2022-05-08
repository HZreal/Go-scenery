# 监听日志文件，将新日志送到kafka中



### demo内容

本demo主要讲的是利用golang的tail库，监听日志文件的变动，将日志信息发送到kafka中。

注意：单独在一个goland运行，注意监听log文件路径

### 涉及的golang库和可视化工具

* go-ini：用于读取配置文件，统一管理配置项，有利于后其的维护
* sarama：是一个go操作kafka的客户端。目前我用于向kefka发送消息
* tail：类似于linux的tail命令了，读取文件的后几行。如果文件有追加数据，会检测到。就是通过它来监听日志文件
* offsetexplorer:是kafka的可视化工具，这里用来查看消息是否投递成功

### 工作的流程

1. 加载配置，初始化sarama和kafka。

2. 起一个的协程，利用tail不断去监听日志文件的变化。

3. 主协程中一直阻塞等待tail发送消息，两者通过一个管道通讯。一旦主协程接收到新日志，组装格式，然后发送到kafka中

### 环境

确保zookeeper和kafka正常运行。因为还没有使用sarama读取数据，使用offsetexplorer来查看任务是否真的投递成功

### 代码分层

serve来存放写tail服务类和sarama服务类，conf存放ini配置文件，main函数为程序入口






