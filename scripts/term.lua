--只会在主程序首次启动或热更新本脚本加载时运行一次
function setup()
    print("terminate setup event ok.")

    return 1
end

--每个用户断开连接都会触发此事件
--cid：客户端身份标识clientid, error:退出原因
function loop(cid, error)
    print(cid, error)

    return nil
end