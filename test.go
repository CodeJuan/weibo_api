package main

import (
    "flag"
    "fmt"
    "github.com/huichen/gobo"
)

var (
    weibo = gobo.Weibo{}
    access_token = flag.String("access_token", "", "用户的访问令牌")
)

func main() {
    // 解析命令行参数
    flag.Parse()

    // 调用API
    var statuses gobo.Statuses
    params := gobo.Params{"screen_name": "人民日报", "count": 10}
    err := weibo.Call("statuses/user_timeline", "get", *access_token, params, &statuses)

    // 处理返回结果
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, status := range statuses.Statuses {
        fmt.Println(status.Text)
    }
}
