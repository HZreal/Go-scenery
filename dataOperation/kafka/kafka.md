# Kafka

## 一、kafka介绍

### 1.介绍

Kafka是Apache旗下的一款分布式流媒体平台，Kafka是一种高吞吐量、持久性、分布式的发布订阅的消息队列系统。

主要用于处理消费者规模网站中的所有动作流数据。动作指(网页浏览、搜索和其它用户行动所产生的数据)。

### 2. 相关概念

术语

1. Topic(话题)：Kafka中用于区分不同类别信息的类别名称。由producer指定
2. Producer(生产者)：将消息发布到Kafka特定的Topic的对象(过程)
3. Consumers(消费者)：订阅并处理特定的Topic中的消息的对象(过程)
4. Broker(Kafka服务集群)：已发布的消息保存在一组服务器中，称之为Kafka集群。集群中的每一个服务器都是一个代理(Broker)。消费者可以订阅一个或多个话题，并从Broker拉数据，从而消费这些已发布的消息。
5. Partition(分区)：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
6. Message(消息)，是通信的基本单位，每个producer可以向一个topic（主题）发布一些消息。消息由一个固定大小的报头和可变长度但不透明的字节阵列负载。报头包含格式版本和CRC32效验和以检测损坏或截断 

四个核心API：

1. Producer API：生产者API允许应用程序将一组记录发布到一个或多个Kafka Topic中。
2. Consumer API：消费者API允许应用程序订阅一个或多个Topic，并处理向他们传输的记录流。
3. Streams API：流API允许应用程序充当流处理器，从一个或者多个Topic中消费输入流，并将输出流生成为一个或多个输出主题，从而将输入流有效地转换为输出流。
4. Connector API：连接器API允许构建和运行可重用的生产者或消费者，这些生产者或消费者将Kafka Topic连接到现有的应用程序或数据系统。例如：连接到关系数据库的连接器可能会捕获对表的每次更改。

### 3. 消息系统

1. 常见的消息系统有Kafka、RabbitMQ、ActiveMQ等

2. 消息模式如下两种：
   * Peer-to-Peer  (Queue) PTP队列模式/点到点模式  采用同步模型
     * 例如：
    
       单发邮件，我发送一封邮件给小徐，我发送过之后邮件会保存在服务器的云端，当小徐打开邮件客户端并且成功连接云端服务器后，可以自动接收邮件或者手动接收邮件到本地，当服务器云端的邮件被小徐消费过之后，云端就不再存储(这根据邮件服务器的配置方式而定)
     * 工作原理

       1. 消息生产者Producer1生产消息到Queue，然后Consumer1从Queue中取出并且消费消息。
       2. 消息被消费后，Queue将不再存储消息，其它所有Consumer不可能消费到已经被其它Consumer消费过的消息。
       3. Queue支持存在多个Producer，但是对一条消息而言，只会有一个Consumer可以消费，其它Consumer则不能再次消费。
       4. 但Consumer不存在时，消息则由Queue一直保存，直到有Consumer把它消费。

   * Publish/Subscribe（Topic）发布/订阅模式  采用异步模型
     * 例如：
    
       我微博有30万粉丝，我今天更新了一条微博，那么这30万粉丝都可以接收到我的微博更新，大家都可以消费我的消息。
     * 工作原理：

       1. 消息发布者Publisher将消息发布到主题Topic中，同时有多个消息消费者 Subscriber消费该消息。
       2. 和PTP方式不同，发布到Topic的消息会被所有订阅者消费。
       3. 当发布者发布消息，不管是否有订阅者，都不会报错信息。
       4. 一定要先有消息发布者，后有消息订阅者。
     * 发布订阅模式实现P-T-P的模式
       
       即将多个消费者加入一个消费者组中，这样这个消费组只能有一个消费者接收消费

3. Kafka所采用的就是发布/订阅模式，被称为一种高吞吐量、持久性、分布式的发布订阅的消息队列系统

## 二、kafka特点

### 1.高吞吐量

可以满足每秒百万级别消息的生产和消费。

### 2.持久性

有一套完善的消息存储机制，确保数据高效安全且持久化。

### 3.分布式

基于分布式的扩展；Kafka的数据都会复制到几台服务器上，当某台故障失效时，生产者和消费者转而使用其它的Kafka。

## 三、应用场景

### 1. 监控

主机通过Kafka发送与系统和应用程序健康相关的指标，然后这些信息会被收集和处理从而创建监控仪表盘并发送警告。

### 2. 消息队列

应用程度使用Kafka作为传统的消息系统实现标准的队列和消息的发布—订阅，例如搜索和内容提要（Content Feed）。比起大多数的消息系统来说，Kafka有更好的吞吐量，内置的分区，冗余及容错性，这让Kafka成为了一个很好的大规模消息处理应用的解决方案。消息系统 一般吞吐量相对较低，但是需要更小的端到端延时，并尝尝依赖于Kafka提供的强大的持久性保障。在这个领域，Kafka足以媲美传统消息系统，如ActiveMR或RabbitMQ

### 3. 站点的用户活动追踪

为了更好地理解用户行为，改善用户体验，将用户查看了哪个页面、点击了哪些内容等信息发送到每个数据中心的Kafka集群上，并通过Hadoop进行分析、生成日常报告。

### 4. 流处理

保存收集流数据，以提供之后对接的Storm或其他流式计算框架进行处理。很多用户会将那些从原始topic来的数据进行阶段性处理，汇总，扩充或者以其他的方式转换到新的topic下再继续后面的处理。例如一个文章推荐的处理流程，可能是先从RSS数据源中抓取文章的内容，然后将其丢入一个叫做“文章”的topic中；后续操作可能是需要对这个内容进行清理，比如回复正常数据或者删除重复数据，最后再将内容匹配的结果返还给用户。这就在一个独立的topic之外，产生了一系列的实时数据处理的流程。

### 5. 日志聚合

使用Kafka代替日志聚合（log aggregation）。日志聚合一般来说是从服务器上收集日志文件，然后放到一个集中的位置（文件服务器或HDFS）进行处理。然而Kafka忽略掉文件的细节，将其更清晰地抽象成一个个日志或事件的消息流。这就让Kafka处理过程延迟更低，更容易支持多数据源和分布式数据处理。比起以日志为中心的系统比如Scribe或者Flume来说，Kafka提供同样高效的性能和因为复制导致的更高的耐用性保证，以及更低的端到端延迟

### 6. 持久性日志

Kafka可以为一种外部的持久性日志的分布式系统提供服务。这种日志可以在节点间备份数据，并为故障节点数据回复提供一种重新同步的机制。Kafka中日志压缩功能为这种用法提供了条件。在这种用法中，Kafka类似于Apache BookKeeper项目。

## 四、架构原理流程

### 1. 架构

* Producer：消息和数据的生产者，主要负责生产Push消息到指定Broker的Topic中，是消息的⼊口。
* kafka cluster：kafka集群，一台或多台服务器组成
  * Broker：Broker是指部署了Kafka实例的服务器节点。负责创建Topic，存储Producer所发布的消息，记录消息处理的过程，现是将消息保存到内存中，然后持久化到磁盘
  * Topic：消息的主题，可以理解为消息的分类，kafka的数据就保存在topic。在每个broker上都可以创建多个topic，一个topic也可以在多个broker上。一个Topic包含一个或者多个Partition分区，实际应用中通常是一个业务线建一个topic。
  * Partition：Topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞吐量。同一个topic在不同的分区的数据是不重复的，partition的表现形式就是一个一个的⽂件夹！
  * Replication:每一个分区都有多个副本，副本的作用是做备胎。当主分区（Leader）故障的时候会选择一个备胎（Follower）上位，成为Leader。在kafka中默认副本的最大数量是10个，且副本的数量不能大于Broker的数量，follower和leader绝对是在不同的机器，同一机器对同一个分区也只可能存放一个副本（包括自己）
* Consumer：消息和数据的消费者，是消息的出口。主要负责主动到已订阅的Topic中拉取消息并消费
  * Consumer Group：我们可以将多个消费组组成一个消费者组，在kafka的设计中同一个分区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个topic的不同分区的数据，这也是为了提高kafka的吞吐量！
* ZooKeeper：ZooKeeper负责维护整个Kafka集群的状态，存储Kafka各个节点的信息及状态，实现Kafka集群的高可用，协调Kafka的工作内容。

### 2. 工作流程

1. ⽣产者从Kafka集群获取分区leader信息
2. ⽣产者将消息发送给leader
3. leader将消息写入本地磁盘
4. follower从leader拉取消息数据
5. follower将消息写入本地磁盘后向leader发送ACK
6. leader收到所有的follower的ACK之后向生产者发送ACK

### 3. 选择partition的原则

1. partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
2. 如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
3. 如果既没指定partition，又没有设置key，则会采用轮询⽅式，即每次取一小段时间的数据写入某
个partition，下一小段的时间写入下一个partition

### 4. ACK应答机制

producer在向kafka写入消息的时候，可以设置参数来确定是否确认kafka接收到数据，这个参数可设置的值为0,1,all
* 0代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效率最高。
* 1代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功。
* all代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保leader发送成功和所有follower都完成备份。安全性最高，但是效率最低。
  
最后要注意的是，如果往不存在的topic写数据，kafka会⾃动创建topic，partition和replication的数量默认配置都是1。

### 5. Topic和数据⽇志

topic 是同⼀类别的消息记录（record）的集合。在Kafka中，⼀个主题通常有多个订阅者。对于每个主题，Kafka集群维护了⼀个分区数据⽇志⽂件

每个partition都是⼀个有序并且不可变的消息记录集合。当新的数据写⼊时，就被追加到partition的末尾。在每个partition中，每条消息都会被分配⼀个顺序的唯⼀标识，这个标识被称为offset，即偏移量。注意，Kafka只保证在同⼀个partition内部消息是有序的，在不同partition之间，并不能保证消息有序。

Kafka可以配置⼀个保留期限，⽤来标识⽇志会在Kafka集群内保留多⻓时间。Kafka集群会保留在保留期限内所有被发布的消息，不管这些消息是否被消费过。⽐如保留期限设置为两天，那么数据被发布到Kafka集群的两天以内，所有的这些数据都可以被消费。当超过两天，这些数据将会被清空，以便为后续的数据腾出空间。由于Kafka会将数据进⾏持久化存储（即写⼊到硬盘上），所以保留的数据⼤⼩可 以设置为⼀个⽐较⼤的值。

### 6. Partition结构

Partition在服务器上的表现形式就是⼀个⼀个的⽂件夹，每个partition的⽂件夹下⾯会有多组segment⽂件，每组segment⽂件⼜包含.index⽂件、.log⽂件、.timeindex⽂件三个⽂件，其中.log⽂件就是实际存储message的地⽅，⽽.index和.timeindex⽂件为索引⽂件，⽤于检索消息。

### 7. 消费数据

多个消费者实例可以组成⼀个消费者组，并⽤⼀个标签来标识这个消费者组。⼀个消费者组中的不同消费者实例可以运⾏在不同的进程甚⾄不同的服务器上。

如果所有的消费者实例都在同⼀个消费者组中，那么消息记录会被很好的均衡的发送到每个消费者实例。

如果所有的消费者实例都在不同的消费者组，那么每⼀条消息记录会被⼴播到每⼀个消费者实例。

在同⼀个消费者组中，每个消费者实例可以消费多个分区，但是每个分区最多只能被消费者组中的⼀个实例消费

## 五、服务启动(开放端口9092)

### 1. kraft方式启动(Mac)

1. 为新集群生成ID，假设输出: m6gE1wP1RIqJWtZPA_A4mg

       kafka-storage random-uuid

2. 格式化存储目录: 单节点模式下执行如下(集群搭建后续)

         kafka-storage format -t <uuid> -c /usr/local/etc/kafka/kraft/server.properties

3. 启动kafka server

       kafka-server-start /usr/local/etc/kafka/kraft/server.properties

### 2. zookeeper方式启动(Mac)

可通过jps命令查看java进程

1. 先启动zookeeper

       zookeeper-server-start -daemon /usr/local/etc/kafka/zookeeper.properties &
   
2. 再启动Kafka

       kafka-server-start -daemon /usr/local/etc/kafka/server.properties &

### 3. kraft方式搭建集群(Mac)

1. 修改配置 /kraft/server.properties
   * 假设三个节点node，分别修改配置中node.id为1、2、3
    
2. 为集群生成uuid(同上)

       kafka-storage random-uuid

3. 三台主机格式化存储目录(同上)

         kafka-storage format -t <uuid> -c /usr/local/etc/kafka/kraft/server.properties

4. 启动节点服务(同上)

       kafka-server-start /usr/local/etc/kafka/kraft/server.properties

## 六、基本使用

### 1. 命令行使用

1. 创建topic
   
   * 如创建名为order的topic消息类型

         kafka-topics --create --topic order --bootstrap-server 127.0.0.1:9092 --partitions 1 --replication-factor 1
   
   * 查看topic=order的信息
   
         kafka-topics --describe --topic order --bootstrap-server 127.0.0.1:9092

2. 发送消息
   
   * 进入发送消息控制台

         kafka-console-producer --topic order --bootstrap-server 127.0.0.1:9092
   
   * 在发送控制台输入发送的内容，每一行都会被看做独立的消息发送到topic中

3. 接收消息

   * 单个消费者消费

     启动一个消费者，进入监听等待消费状态。参数--from-brginning 选项表示从该Topic的开始位置消费消息，未指定则监听等待后续接收到的消息

         kafka-console-consumer --topic order --from-beginning --bootstrap-server localhost:9092

   * 多个消费者消费消息
     
     * 启动两个消费者，并分别设置两个消费者对应的消费组为consumer-group1、consumer-group2，参数--consumer-property 指定消费者所属的消费组，注意在两个终端中执行，进入两个消费者监听状态

           kafka-console-consumer --topic order --bootstrap-server localhost:9092 --consumer-property group.id=consumer-group1
           kafka-console-consumer --topic order --bootstrap-server localhost:9092 --consumer-property group.id=consumer-group2
     
     * 查看消费组的信息

           kafka-consumer-groups --bootstrap-server 127.0.0.1:9092 --list            仅查看消费组列表
           kafka-consumer-groups --bootstrap-server 127.0.0.1:9092 --group consumer-group1 --describe      指定组名详细查看，输入结果中: CURRENT-OFFSET（表示消费组已经消费的消息的偏移量）、LOG-END-OFFSET（主题对应分区消息的结束偏移量）、LAG（表示消费组未消费的消息的数量）

4. 删除指定主题的消息

   * 需要创建一个删除消息的json文件，如命名为delete-order.json，文件内容如下

         {
             "partitions": [               // 按照此分区列表的信息进行删除
                 {
                     "topic": "order",    // 待删除的主题的名称
                     "partition": 0,      // 分区编号
                     "offset": -1         // 删除的区间，-1 表示删除所有的消息
                 }
                 // 如果有多个分区，则需要添加多个 partition 对象
             ],
             "version": 1
         }

   * 执行删除

         kafka-delete-records --bootstrap-server 127.0.0.1:9092 --offset-json-file ./delete-order.json

### 2. go连接kafka使用

#### producer.py

#### consumer.py

### 3. python连接kafka使用