package main

import (
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/pquerna/ffjson/ffjson"
)

//告诉内核此插件为用户身份验证插件
var CT = "auth"

//告诉内核本插件版本号，当内核重启首次加载多个插件时会生效Ver值最大的相同性质插件
var Ver float32 = 1.1

////jwt key
//var secretKey = "Coolpy7yeah"

//主程序首次启动或热更新后加触发一次
func Setup() {
	log.Println("Authentication plugin ok version", Ver)
}

func Loop(cid, username, password, remote string) bool {
	log.Println(cid, username, password, remote)

	////固定值判断认证登陆信息合法性
	//if cid == "system" && username == "premissid" && password == "testpremissid" {
	//	return true
	//}

	////jwt token
	//token, err := jwt.Parse(secretKey, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(secretKey), nil
	//})
	//if err != nil {
	//	return false
	//}
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	log.Println(claims)
	//	return true
	//} else {
	//	return false
	//}

	////http请求验证
	////参数: method, url, header key, header value, content type, body
	//res, err := httpFunc("GET", "https://192.168.200.200/api/v1/sys/authenticate/" + username, "Authorization", "5ba32f074966b41be9123fe1", "application/json", "")
	//if err == nil && res["ok"].(bool) == true {
	//	return true
	//}

	//不做作何身份验证直接允许登入
	return true
}

//内核逻辑处理异常触发此事件
func Err(cid, username, password, remote, err string) {
	log.Println(cid, username, password, remote, err)
}

//http请求
func httpFunc(method, url, headerKey, headerValue, contentType, body string) (map[string]interface{}, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	reqest, err := http.NewRequest(method, url, strings.NewReader(body))
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
