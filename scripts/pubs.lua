--注入模式
--false为并行模式(一般用于只作消息日志记录功能使用此模式)
--true为同步模式（返回bool类型生效，true为正常发送，false为消息强制取消发送, 用于消息逻辑功能使用)
sync = false

--只会在主程序首次启动或热更新本脚本加载时运行一次
function setup()
    print("pubs setup event ok.")

    return 1
end

--每个用户消息推送都会触发此事件
--cid：客户端身份标识clientid, topic:主题，qos: 消息质量, payload:消息内容
--当sync设置为true时，返回值如果是字符串类型即返回值内容会替换原消息内容发送
function loop(cid, topic, qos, payload)
    --print(cid, topic, qos, payload)

    ----通过http记录消息
    --if cid == "system" then
    --    --参数: method, url, header key, header value, content type, body
    --    ok, res = cp7.http("POST", "https://192.168.200.200/api/v1/sys/chatLog", "Authorization", "5ba32f074966b41be9123fe1", "application/json", payload)
    --    if ok then
    --        print("send http message ok.")
    --    end
    --end

    ----此示例演示接收到一个json消息后
    ----增加两个键，然后打印出来
    ----通过sync模式返回新的payload给内核替换原来的payload
    --ok, res = cp7.jsonUnmarshal(payload)
    --if ok then
    --    res["createat"] = cp7.now()
    --    res["id"] = cp7.noid()
    --    --打印
    --    --tprint(res, 1)
    --
    --    --当sync设置为true，return内容将更新payload值
    --    ok1, result = cp7.jsonMarshal(res)
    --    if ok1 then
    --        return result
    --    end
    --else
    --    print(res)
    --end

    --不处理直接返回原payload
    return nil
end

--内核消息推送过程处理相关的错误提示会触发此方法
function err(cid, topic, qos, payload, err)
    print("pubs err event")
    print(cid, topic, qos, payload, err)

    return 1
end

--打印table以debug使用
function tprint (tbl, indent)
    if not indent then
        indent = 0
    end
    for k, v in pairs(tbl) do
        formatting = string.rep("", indent) .. k .. ": "
        if type(v) == "table" then
            print(formatting)
            tprint(v, indent + 1)
        elseif type(v) == 'boolean' then
            print(formatting .. tostring(v))
        else
            print(formatting .. v)
        end
    end
end