--jwt key
secretKey = "Coolpy7yeah"

--只会在主程序首次启动或热更新本脚本加载时运行一次
function setup()
    print("auth setup event ok.")
end

--每个用户登陆都会触发此事件 return true即允许登入
function loop(cid, username, password, remote)
    --print(cid, username, password, remote)

    ----固定值判断认证登陆信息合法性
    --if (cid == "system" and username == "premissid" and password == "testpremissid") then
    --    return true
    --end

    ----jwt 验证 参数: 1 jwt token ,2 密钥
    --ok, res = cp7.parseJwt(password, secretKey)
    --if ok then
    --    print(ok, res["sub"], res["name"], res["iat"])
    --    return true;
    --else
    --    print("auth jwt token err", res)
    --end

    ----http请求验证
    ----参数: method, url, header key, header value, content type, body
    --ok, res = cp7.http("GET", "https://192.168.200.200/api/v1/sys/authenticate/" + username, "Authorization", "5ba32f074966b41be9123fe1", "application/json", "")
    --if ok then
    --    print(ok, res["data"])
    --    return true;
    --else
    --    print("auth jwt token err", res)
    --end

    --不做作何身份验证直接允许登入
    return true
end

--内核身份验证过程处理相关的错误提示会触发此方法
function err(cid, username, password, remote, err)
    print("auth err event")
    print(cid, username, password, remote, err)
end