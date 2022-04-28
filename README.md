# quicgo-for-android
# 简介
quicgo-for-android是基于quic-go进行二次封装的Android 库，选取了quic-go的QUIC层的API进行JAVA封装以方便在Android端进行QUIC通信。
# 环境
1. Android studio
2. ndk 24
3. [quic-go](https://github.com/lucas-clemente/quic-go)
4. gomobile
# 项目结构
## go端
quicgoandroid文件夹下，使用go对quic-go的API进行封装以符合gomobile的交叉编译条件
## Java端
### quicgo包
此包内包含了封装的quic-go API
### util包
此包内包含了两个工具类，分别为JSON解析工具类以及Log工具类
### MainActivity
实现了一个echo demo
# API的使用
详见[文档]()
