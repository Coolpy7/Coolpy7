[![Author](https://img.shields.io/badge/author-@jacoblai-blue.svg?style=flat)](http://coolpy.net/) [![Platform](https://img.shields.io/badge/platform-Linux,%20OpenWrt,%20Arm,%20Mac,%20Windows-green.svg?style=flat)](https://github.com/coolpy7/coolpy7) [![MQTT](https://img.shields.io/badge/MQTT3.1.1-pink.svg?tyle=flat)](https://github.com/coolpy7/coolpy7)

# Coolpy7

**一个高性能、高稳定性的跨平台MQTT服务端**

一个高性能、高稳定性的跨平台MQTT服务端，基于EPOLL之上开发，可以在嵌入式设备（OpenWrt/Arm）、Linux、Windows、Mac上使用，拥有完善Qos功能和配套开源周边功能库，以极少的资源实现优质的单机百万千万级MQTT服务，并且无缝衔接主流数据库。

## 优势：

- **具有极高的稳定性**：无论是**掉线重连**，**丢包重发**，都是**严格遵循MQTT协议标准**执行，除此之外对**大数据量**的测试无论是收是发，都是非常稳定，**高频**测试也是非常稳定。

- **完善的技术文档**：https://coolpy7.gitbook.io/coolpy7book/

- **多种技术客户端示例** [APP(Android,Flutter);后端(Golang,Java,Nodejs,Python3,C#);前端(Electron,React,Vue,微信小程序,WebSocket);单片机(ESP8266)](https://github.com/Coolpy7/mqtt-client-examples)

- **性能测试项目**，https://github.com/Coolpy7/coolpy7_benchmark

- **内核功能扩展服务**，https://github.com/Coolpy7/coolpy7_extend_service

- **MQTT-SN项目**，[MQTT-SN网关](https://github.com/Coolpy7/mqttsn-gateway)，[MQTT-SN下位机开发包](https://github.com/Coolpy7/mqttsn-sdk)，[MQTT-SN网关](https://github.com/Coolpy7/mqttsn-gateway)

- **ESP8266下位机(单片机客户端)开发包**，https://github.com/Coolpy7/EspSdk

- **数据库代理服务**，https://github.com/Coolpy7/DbPoxy

- **TLS加密连接代理服务**，https://github.com/Coolpy7/Coolpy7-TLS-Poxy

- **WebSocket代理服务器**，https://github.com/Coolpy7/Coolpy7/blob/master/go_build_Coolpy7_ws_go_linux.zip

- **WebSocket-TLS代理服务**，https://github.com/Coolpy7/Coolpy7-WS-TLS-Poxy

- **OATUTH2.0用户中心**，[Coolpy7用户中心服务器端](https://github.com/Coolpy7/oauth2_server)，[用户中心Web前端UI](https://github.com/Coolpy7/oauth2_ui)

## 应用示例
- **多种技术客户端示例** https://github.com/Coolpy7/mqtt-client-examples

- **上传图片或文件到DbPoxy代理服务器**，https://github.com/Coolpy7/dbpoxyclient

- **Web浏览器应用示例**，https://github.com/Coolpy7/mqtt_web_browser_client

- **微信小程序聊天室示例**，https://github.com/Coolpy7/wxsmallapp

- **Web浏览器聊天室示例（可与微信小程序示例连同一Coolpy7后互相聊天）**：https://github.com/Coolpy7/Cp7Chat

- **Web浏览器mqtt.js客户端示例**， https://github.com/Coolpy7/mqttjs_browser_client_demo

## 整体框架

拥有非常明确的分层框架。

![整体框架](https://gblobscdn.gitbook.com/assets%2F-LPz4APWKTuf0FRQG1lh%2F-M0a3JzqBKD0TtVU3f7a%2F-M0a3UKQ3w-nqiZAybZf%2FCoolpy7%E6%96%B0%E6%9E%B6%E6%9E%84.png?alt=media&token=79671425-a1bc-42a5-8184-1af1a51e8b36)

## 支持的平台

**目前已实现了Linux、Mac、Arm64、Mips平台，除此之外更多框架需要使用，请提交Issues需求！**

| 平台           | 代码位置 |
| -------------- | -------- |
| Linux          | [go_build_Coolpy7_go_linux.zip](https://github.com/Coolpy7/Coolpy7/blob/master/go_build_Coolpy7_go_linux.zip) |
| Mac            | [go_build_Coolpy7_go_mac](https://github.com/Coolpy7/Coolpy7/blob/master/go_build_Coolpy7_go_mac.zip) |
| Windows        | [docker_windows_coolpy7](https://coolpy7.gitbook.io/coolpy7book/kai-shi-shi-yong/docker-bu-shu-windows-yun-hang-cp7) |
| Arm64          | [NanoPi,JetsonNano,树莓派等等...](https://github.com/Coolpy7/Coolpy7/blob/master/go_build_Coolpy7_go_arm64_linux.zip) |
| Mips(le)      | [OpenWrt(MT7688)等等...](https://github.com/Coolpy7/Coolpy7/blob/master/go_build_Coolpy7_go_mipsle_linux.zip) |


## 版本

| 发布版本 | 描述 | 
| --- | --- |
| [v7.3.2] | 使用extend扩展功能服务代替之前以插件方式开发扩展功能 |
| [最新版本] | [release下载](https://github.com/Coolpy7/Coolpy7/releases) |

## 问题

欢迎以 [GitHub Issues](https://github.com/Coolpy7/Coolpy7/issues) 的形式提交问题和bug报告

## 版权和许可

Coolpy7 遵循 MIT 开源协议。可以自由的使用、修改源代码，也可以将修改后的代码作为开源或闭源软件发布。也可直接用于商业项目。

## QQ群

<img src="https://github.com/Coolpy7/Coolpy7/blob/master/qq.jpg" width="400" height="688">
