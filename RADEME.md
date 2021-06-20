#### 项目参考
- github.com/SananGuliyev/goddd


### Application 应用层
### Interfaces 接口层
 - mq 消息队列抽象
 - client 客户端接口（防腐层ACL接口部分）
### Domain 领域
 - aggregate 聚合
 - entity 实体
 - interactor 交互器
 - repository 仓储（定义实现方法）
 - throwable 异常类型
 - valobj 值对象
### Infrastructure 基础设施
 - persistence 持久化
    - postgresql 
    - redis
 - security
 - auth
 


> interactor 接收处理来自Presenter(路由)传递过来的参数
> repository 定义实现持久化方法persistence来实现它

> 一块蛋糕（View）只由一个主厨（Presenter）进行装饰摆放，但是蛋糕上所有的饰品食材，都由这位主厨指派给多个不同的厨师（Interactor）进行加工（Entity）

### Anti-corruption Layer(ACL) 反腐层