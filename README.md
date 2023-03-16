# requests

#### 介绍

封装 golang get post 请求

#### 安装教程

```bash
go get github.com/hucgqg/requests
```

#### 使用说明

```go
package main

import (
	"github.com/hucgqg/requests"
)

func main() {
    url := "https://test.com/api"
    r := requests.Request{Url: &url, Method: "POST", Data: &map[string]interface{}{}, Headers: &map[string]string{},BasicAuth: &map[string]string{}}
    r.Body()
    fmt.Println(r.RepInfo)
}
```
