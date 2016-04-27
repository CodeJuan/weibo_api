package main

import (
    "flag"
    "fmt"
    "github.com/huichen/gobo"
    "strconv"
)

var (
    weibo = gobo.Weibo{}
    access_token = flag.String("access_token", "", "用户的访问令牌")
)

func main() {
    // 解析命令行参数
    flag.Parse()

    //post
    postWeibo()

    // 调用API
    var statuses gobo.Statuses
    params := gobo.Params{"sort": 0, "lat":34.19394, "long":108.84218, "range":11132}
    err := weibo.Call("place/nearby_timeline", "get", *access_token, params, &statuses)

//    params := gobo.Params{"poiid": "B2094656D269A6FA419F"}
//    err := weibo.Call("place/poi_timeline", "get", *access_token, params, &statuses)

    // 处理返回结果
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(statuses.Statuses)
    for _, status := range statuses.Statuses {
        fmt.Println(status)
        id, _ := strconv.ParseInt(status.Id, 10, 64)
        comments := gobo.Params{"id": id, "attitude": "smile"}
        err := weibo.Call("attitudes/create", "post", *access_token, comments, nil)
        if err != nil {
            fmt.Println(err)
        }
    }
}


func postWeibo(){
//    var statuses gobo.Statuses
//    params := gobo.Params{"status": "test2", "lat":34.19394, "long":108.84218}
//    err := weibo.Call("statuses/update", "post", *access_token, params, &statuses)
//
//    // 处理返回结果
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
}