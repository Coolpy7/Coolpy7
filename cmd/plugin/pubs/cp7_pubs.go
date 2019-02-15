package main

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//告诉内核此插件为消息过滤插件
var CT = "pubs"

//只会在主程序首次启动或热更新时运行一次
func Setup() {
	log.Println("pubs setup event ok.")
}

//每个用户消息推送都会触发此事件
//cid：客户端身份标识clientid, topic:主题，qos: 消息质量, payload:消息内容
//返回值1：当为true时，返回值2会替换原消息内容发送
//返回值3：内核判断是否断开当前客户端
func Loop(cid, topic string, qos uint8, payload []byte) (bool, []byte, error) {
	//log.Println(cid, topic, qos, string(payload))

	////http请求验证
	////参数: method, url, header key, header value, content type, body
	//_, err := httpFunc("POST", "https://192.168.200.200/api/v1/sys/chatLog", "Authorization", "5ba32f074966b41be9123fe1", "application/json", payload)
	//if err == nil {
	//	return false, nil, nil
	//}

	////此示例演示接收到一个json消息后
	////增加两个键，然后打印出来
	////通过返回替换模式返回新的payload给内核替换原来的payload
	//var res map[string]interface{}
	//err := ffjson.Unmarshal(payload, &res)
	//if err != nil {
	//	return false, nil, nil
	//}
	////添加两个键
	//res["createat"] = time.Now().Local()
	//res["id"] = objectid.NewObjectId().Hex()
	////打印
	//log.Println(res)
	////json序列化后返回给内核
	//nPayload, _ := ffjson.Marshal(&res)
	////返回true即内核替换payload为nPayload
	//return true, nPayload, nil

	return false, nil, nil
}

//内核过程处理相关的错误提示会触发此方法
func Err(cid, topic string, qos uint8, payload []byte, err string) {
	log.Println("pubs err event")
	log.Println(cid, topic, qos, string(payload), err)
}

//http请求
func httpFunc(method, url, headerKey, headerValue, contentType string, body []byte) (map[string]interface{}, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if headerKey != "" {
		reqest.Header.Add(headerKey, headerValue)
	}
	if contentType != "" {
		reqest.Header.Set("Content-Type", contentType)
	} else {
		reqest.Header.Set("Content-Type", "application/json;charset=utf-8")
	}
	response, err := client.Do(reqest)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		resBody := make([]byte, 0)
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			defer reader.Close()
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					return nil, err
				}

				if n == 0 {
					break
				}
				resBody = append(resBody, buf...)
			}
		default:
			resBody, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
		}
		var obj map[string]interface{}
		err = json.Unmarshal(resBody, &obj)
		if err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, err
}
