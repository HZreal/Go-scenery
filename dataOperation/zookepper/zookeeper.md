# zookeeper

## 一、什么是ZooKeeper

an open-source server which enables highly reliable distributed coordination

a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services.

## 二、为什么ZooKeeper能干这么多？

* ZooKeeper主要服务于分布式系统，可以用ZooKeeper来做：统一配置管理、统一命名服务、分布式锁、集群管理。
* 使用分布式系统就无法避免对节点管理的问题(需要实时感知节点的状态、对节点进行统一管理等等)，而由于这些问题处理起来可能相对麻烦和提高了系统的复杂性，ZooKeeper作为一个能够通用解决这些问题的中间件就应运而生

ZooKeeper的数据结构，跟Unix文件系统非常类似，可以看做是一颗树，每个节点叫做ZNode。每一个节点可以通过路径来标识，Znode分为两种类型：

* 短暂/临时(Ephemeral)：当客户端和服务端断开连接后，所创建的Znode(节点)会自动删除
* 持久(Persistent)：当客户端和服务端断开连接后，所创建的Znode(节点)不会删除

ZooKeeper还配合了监听器，常见的监听场景有以下两项：

* 监听Znode节点的数据变化
* 监听子节点的增减变化

## 三、ZooKeeper是怎么做到的？

### 3.1 统一配置管理

比如我们现在有三个系统A、B、C，他们有三份配置，分别是ASystem.yml、BSystem.yml、CSystem.yml，然后，这三份配置又非常类似，很多的配置项几乎都一样。

此时，如果我们要改变其中一份配置项的信息，很可能其他两份都要改。并且，改变了配置项的信息很可能就要重启系统

于是，我们希望把ASystem.yml、BSystem.yml、CSystem.yml相同的配置项抽取出来成一份公用的配置common.yml，并且即便common.yml改了，也不需要系统A、B、C重启。

做法：我们可以将common.yml这份配置放在ZooKeeper的Znode节点中，系统A、B、C监听着这个Znode节点有无变更，如果变更了，及时响应。

### 3.2 统一命名服务

统一命名服务的理解其实跟域名一样，是我们为这某一部分的资源给它取一个名字，别人通过这个名字就可以拿到对应的资源。

比如说，现在我有一个域名www.java3y.com，但我这个域名下有多台机器：

* 192.168.1.1
* 192.168.1.2
* 192.168.1.3
* 192.168.1.4

别人访问www.java3y.com即可访问到我的机器，而不是通过IP去访问

### 3.3 分布式锁

系统A、B、C都去访问/locks节点，访问的时候会创建带顺序号的临时/短暂(EPHEMERAL_SEQUENTIAL)节点，比如，系统A创建了id_000000节点，系统B创建了id_000002节点，系统C创建了id_000001节点。

接着，拿到/locks节点下的所有子节点(id_000000,id_000001,id_000002)，判断自己创建的是不是最小的那个节点

* 如果是，则拿到锁。执行完操作后，把创建的节点给删掉
* 如果不是，则监听比自己要小1的节点变化

### 3.4集群状态

只要系统A挂了，那/groupMember/A这个节点就会删除，通过监听groupMember下的子节点，系统B和C就能够感知到系统A已经挂了。(新增也是同理)

除了能够感知节点的上下线变化，ZooKeeper还可以实现动态选举Master的功能。(如果集群是主从架构模式下)

* Zookeeper会每次选举最小编号的作为Master，如果Master挂了，自然对应的Znode节点就会删除。然后让新的最小编号作为Master，这样就可以实现动态选举的功能了。



## 四、安装搭建配置使用:

### 1.mac安装:

    brew install zookeeper

* 安装目录：/usr/local/Cellar/zookeeper/3.7.0_1    
  ./bin/目录包含zkServer、zkCli等可执行文件
* 默认配置文件: /usr/local/etc/zookeeper/zoo.cfg
* 启动server:   zkServer start
  * zkServer start zk.cfg   指定配置文件启动
  * zkServer stop/status
* 启动zk客户端进行连接:     
  * zkCli -server 127.0.0.1:2181
  * zkCli -server 127.0.0.1:21811 127.0.0.1:21812 127.0.0.1:21813 127.0.0.1:21814

* 客户端连接默认端口2181、leader选举端口3888、集群节点通讯数据同步端口2888

### 2.配置环境变量:

在~/.bash_profile 文件中添加

    export ZK_HOME=/usr/local/Cellar/zookeeper/3.7.0_1 
    export PATH=$PATH:$ZK_HOME/bin

### 3.zk客户端命令行使用:

https://zookeeper.apache.org/doc/r3.8.0/zookeeperCLI.html

* 查看当前节点的子节点（可监听）:
    
  ls path路径 -w监听子节点变化/-s附加次级信息/-R递归查看

  例如 ls /
  
* 创建普通节点:     
  create -s含有序列/-e临时（客户端会话断开即删除）/-c容器节点，当该容器中没有节时，超时后被删除（60s）
  例如 create -s /node111 helloworld

* 获得指定路径下节点的值（可监听）:    
  get path路径 -w监听节点内容的变化/-s附加次级信息     get/node1
* set	设置节点的具体信息    set/node1
* stat	查看节点状态     stat /node1
* delete	删除节点    delete /node1
* deleteall	递归删除节点

### 4.集群搭建

1. 创建存放四个节点数据的文件夹.../cluster_data/，并在每个节点文件夹下创建myid文件，存放节点serverid
   
    mkdir ../cluster_data/zk1  
    mkdir ../cluster_data/zk2  
    mkdir ../cluster_data/zk3  
    mkdir ../cluster_data/zk4  
   echo 1 > zk1/myid  
   echo 2 > zk2/myid  
   echo 3 > zk3/myid  
   echo 4 > zk4/myid  
   
   
2. 四个节点的配置

* 复制/usr/local/etc/zookeeper/zoo.cfg文件四份分别命名为zk1.cfg,zk2.cfg,zk3.cfg,zk4.cfg,
* 每个配置文件都加入如下：
  * server.1=127.0.0.1:2888:3888
    server.2=127.0.0.1:2889:3889
    server.3=127.0.0.1:2890:3890
    server.4=127.0.0.1:2891:3891:observer
* 每个配置文件分别修改dataDir指向自己对应的数据文件目录如zk1.cfg修改为 ../cluster_data/zk1
* 每个配置文件分别修改客户端访问端口clientPort，如分别为21811 21812 21813 21814

3. 启动集群并查看状态，可以看到哪些节点成为leader,哪些成为follower

* zkServer start zk1.cfg
* zkServer start zk2.cfg
* zkServer start zk3.cfg
* zkServer start zk4.cfg
* zkServer status zk1.cfg
* zkServer status zk2.cfg
* zkServer status zk3.cfg
* zkServer status zk4.cfg


4. 客户端连接集群

* zkCli -server 127.0.0.1:21811,127.0.0.1:21822,127.0.0.1:21833,127.0.0.1:21844

## 五、服务流程

* 1：Client向Follower发出写操作；

* 2：Follower将收到的写请求转发给Leader处理；

* 3：Leader收到请求后分配一个全局单调递增的唯一的事务ID(即ZXID，按其先后顺序来进行排序与处理)；

* 4.1：Leader服务器会为每一个 Follower服务器都各自分配一个单独的队列,然后将需要广播的事务Proposal依次放入这些队列中去,并且根据FIFO策略进行消息发送；每一个 Follower服务器在接收到这个事务 Proposal之后,都会首先将其以事务日志的形式写入到本地磁盘中去,并且在成功写入后反馈给 Leader服务器个Ack响应。Leader自己也会将事务日志写入磁盘；

* 4.2：当 Leader服务器接收到超过半数Follower的Ack响应后,就会广播一个Commit消息给所有的 Follower服务器以通知其进行事务提交（写入内存）,同时 Leader自身也会完成对事务的提交。

* 5：由Leader将结果返回给Client1


