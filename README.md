# 基于GO实现千万级WebSocket消息推送服务
## 拉模式和推模式区别
### 1. 拉模式（定时轮询访问接口获取数据）
- 数据更新频率低，则大多数的数据请求时无效的
- 在线用户数量多，则服务端的查询负载很高
- 定时轮询拉取，无法满足时效性要求
### 2. 推模式（向客户端进行数据的推送）
- 仅在数据更新时，才有推送
- 需要维护大量的在线长连接
- 数据更新后，可以立即推送
- 基于WebSocket协议做推送
- 浏览器支持的socket编程，轻松维持服务端的长连接
- 基于TCP协议之上的高层协议，无需开发者关心通讯细节
- 提供了高度抽象的编程接口，业务开发成本较低

## 基于WebSocket推送
- 浏览器支持的socket编程，轻松维持服务端的长连接
- 基于TCP可靠传输之上的协议，无需开发者关心通讯细节
- 提供了高度抽象的编程接口，业务开发成本低

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200719165336765.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzQzNDQyNTI0,size_16,color_FFFFFF,t_70)
### 传输原理
- 协议升级后，继续复用HTTP的底层Socket完成后续通讯
- message底层被切分成多个frame帧传输
- 编程时只需操作message，无需关心frame
- 框架底层完成TCP网络I/O，webSocket协议解析，开发者无需关心