# requests

#### 介绍
封装golang get post请求

#### 安装教程

```bash
go get gitee.com/hcqcode/requests
```

#### 使用说明
```bash
url := "https://test.com/api"
r := requests.Request{Url: &url, Method: "POST", Data: &map[string]string{}, Headers: &map[string]string{}}
r.Body()
fmt.Println(r.RepInfo)
```

