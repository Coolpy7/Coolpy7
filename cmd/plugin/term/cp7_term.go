package main

import "log"

//告诉内核此插件为用户连接断开插件
var CT = "term"

//告诉内核本插件版本号，当内核重启首次加载多个插件时会生效Ver值最大的相同性质插件
var Ver float32 = 1.1

//主程序首次启动或热更新后加触发一次
func Setup() {
	log.Println("Terminate plugin ok version", Ver)
}

//每个用户断开连接都会触发此事件
//cid：客户端身份标识clientid, error:退出原因
func Loop(cid, err string) {
	log.Println(cid, err)
}
