# Hawos
一个可供快速开发业务逻辑的脚手架


## 环境要求
+ Linux/Darwin
+ golang 1.16.5
+ Redis v6.2.2
+ Nsq v1.2.0 / Kafka v2.8.0
+ etcd v3.5.0


## 服务简介
### frontend
``提供websocket服务，管理客户端长链``
+ [CPU占用分析](https://sp.hawtech.cn/hawpic/images/svg/cpu_frontend.svg)
+ [内存占用分析](https://sp.hawtech.cn/hawpic/images/svg/mem_frontend.svg)

### config
``提供http服务, 通过api获取配置信息等等``
+ [CPU占用分析](https://sp.hawtech.cn/hawpic/images/svg/cpu_config.svg)
+ [内存占用分析](https://sp.hawtech.cn/hawpic/images/svg/mem_config.svg)

### chat
``逻辑服（聊天室），处理聊天信息``
+ [CPU占用分析](https://sp.hawtech.cn/hawpic/images/svg/cpu_chat.svg)
+ [内存占用分析](https://sp.hawtech.cn/hawpic/images/svg/mem_chat.svg)
