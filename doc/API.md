# quicgoandroid API文档

# 概述

- 本项目将quic-go的部分**quic层API**封装为可以被gomobile编译的API，以供Android端使用
- API分为两部分：
    - 第一部分为Direct API，是来自Quic-go一些QUIC层函数的直接封装，以供Java通过aar包进行调用；
    - 第二部分为Java API，通过一些简单封装，使Direct API更加易用

# DirectAPI

## Server

### Listen

- 获取quic服务器的listener进行listen
- 入参：
    - addr：地址字符串

    ```
    "host:port"
    ```

- 返回值：
    - 有错误返回JSON格式的error，无错误返回空字符串

        ```
        {
            error： string
        }
        ```


### Accept

- 返回一个新的quic连接，如果想一直监听，应该将此函数放在一个loop中
- 返回值：
    - connectID从1开始，若创建连接失败则为0

        ```
        {
        	error：string,
        	connect_id: int,
        }
        ```


### **AcceptStream**

- 应答Client发出的建立双向Stream请求
- 输入：
    - connectID：标识stream的连接
- 返回值：
    - streamID与connectID类似，从1开始，失败为0

        ```
        {
        	error：string,
        	stream_id: int,
        }
        ```


## Clinet

### Dial

- 向指定服务器发送请求，返回一个quic连接
- 输入
    - addr
- 返回值

    ```
    {
    	error：string,
    	connect_id: int,
    }
    ```


### OpenStreamSync

- 打开一个双向Stream，如果不能打开，将阻塞
- 输入
    - connectionID
- 返回值

    ```
    {
    	error：string,
    	stream_id: int,
    }
    ```


## common

- Client和Server的共用函数

### ReadStream

- 从Stream中读数据
- 输入：streamID
- 返回值

    ```
    {
    	error: string
    	data: string
    }
    ```


### WriteStream

- 从Stream中读数据
- 输入：
    - streamID
    - message
- 返回值

    ```
    {
    	error: string
    }
    ```


### ReceiveMessage

- 从Connect中读数据报
- 输入：connectID
- 返回值

    ```
    {
    	error: string
    	data: string
    }
    ```


### SendMessage

- 从Connect中发送数据报
- 输入
    - streamID
    - message
- 返回值

    ```
    {
    	error: string
    }
    ```


# JavaAPI

- 将Go API做进一步封装，提高在Android端的易用性

## QuicGoServerEngine

- 需提供一个address字符串以构造，address格式为`host:port`
    - eg：`logclhost:41420`

### listen

- 服务器开始监听

### accept

- 获取一个quic connection
- 阻塞
- 建议在循环中使用以建立多个连接

## QuicGoClientEngine

### dial

- 向服务器请求建立连接

## Connect

- connection的抽象类，拓展为ServerConnnect和ClientConnect两个子类

### send

- 发送udp数据

### get

- 接受udp数据

### ServerConnect.acceptStream

- 在此Connection下建立一个Stream
- 注：此函数需要Client那边通过Stream发送数据才会accept，否则一直阻塞

### ClientConnect.openStreamSync

- 创建一个Stream以发送数据

## Stream

### send

- 通过此流发送数据

### get

- 通过此流接收sendData数据
