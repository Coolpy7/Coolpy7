package main

import "log"

//告诉内核此插件为用户连接断开插件
var CT = "term"

//只会在主程序首次启动或热更新本脚本加载时运行一次
func Setup() {
	log.Println("terminate setup event ok.")
}

//每个用户断开连接都会触发此事件
//cid：客户端身份标识clientid, error:退出原因
func Loop(cid, err string) {
	log.Println(cid, err)
}
