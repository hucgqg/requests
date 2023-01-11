# requests

#### 介绍

封装 golang get post 请求

#### 安装教程

```bash
go get gitee.com/hcqcode/requests
```

#### 使用说明

```go
package main

import (
	"gitee.com/hcqcode/requests"
)

func main() {
    url := "https://test.com/api"
    r := requests.Request{Url: &url, Method: "POST", Data: &map[string]string{}, Headers: &map[string]string{},BasicAuth: &map[string]string{}}
    r.Body()
    fmt.Println(r.RepInfo)
}
```
