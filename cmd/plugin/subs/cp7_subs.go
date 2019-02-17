package main

import (
	"log"
)

//告诉内核此插件为订阅过滤插件
var CT = "subs"

//告诉内核本插件版本号，当内核重启首次加载多个插件时会生效Ver值最大的相同性质插件
var Ver float32 = 1.1

//主程序首次启动或热更新后加触发一次
func Setup() {
	log.Println("Subscribe plugin ok version", Ver)
}

//每个用户消息推送都会触发此事件
//cid：客户端身份标识clientid, topic:主题，qos: 消息质量
//返回值1：当为true时，返回值2会替换原订阅主题
//返回值3：内核判断是否断开当前客户端
func Loop(cid, topic string, qos uint8) (bool, string, error) {
	log.Println(cid, topic, qos)

	return false, "", nil
}

//内核逻辑处理异常触发此事件
func Err(cid, topic string, qos uint8, err string) {
	log.Println("subs err event")
	log.Println(cid, topic, qos, err)
}
