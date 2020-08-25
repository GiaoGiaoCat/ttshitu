<h1 align="center">ttshitu - 图鉴验证码识别</h1>

[![Dependabot](https://api.dependabot.com/badges/status?host=github&repo=GiaoGiaoCat/ttshitu&identifier=290180072)](https://app.dependabot.com/accounts/GiaoGiaoCat/repos/290180072)
[![Go Report Card](https://goreportcard.com/badge/github.com/GiaoGiaoCat/ttshitu)](https://goreportcard.com/report/github.com/GiaoGiaoCat/ttshitu)
[![Release](https://img.shields.io/github/release/GiaoGiaoCat/ttshitu.svg?style=flat-square)](https://github.com/GiaoGiaoCat/ttshitu/releases)

<img align="right" width="140px" src="https://raw.githubusercontent.com/ashleymcnamara/gophers/master/GOPHERCON_ICELAND.png">

方便在 golang 项目中使用[图鉴](http://www.ttshitu.com)的图片验证码识别功能。

**先去图鉴官网注册账号，获得用户名和密码。免费用户可以试用 10 次。如果发现无法返回结果，有可能是免费次数已经用完，需要充值。**

目前只支持 base64 编码的形式提交图片信息。

调用 IdentifyingCode 方法，获取识别后的验证码，如果无法识别则返回空字符串。

## HowTo

```bash
go get github.com/GiaoGiaoCat/ttshitu
```

```golang
package main

import "github.com/GiaoGiaoCat/ttshitu"

func main() {
  req := &ttshitu.RequestBody{Username: "Rocky", Password: "5454", Image: "xxxxx"}
  // 调用 IdentifyingCode 方法，获取识别后的验证码，如果无法识别则返回空字符串。
  code := ttshitu.IdentifyingCode(req)
  fmt.Println(code)
}
```

* Username 和 Password 是在图鉴申请的账号密码。
* Image 是图片的 Base64 编码。

## 参考

* [Making external HTTP requests in Go](https://medium.com/rungo/making-external-http-requests-in-go-eb4c015f8839)
* [Making HTTP Requests in Golang](https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7)
* [How do I send a JSON string in a POST request in Go](https://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go)
* [How to Parse a JSON Request Body in Go](https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body)
* [Encoding and Decoding JSON, with Go’s net/http package](https://kevin.burke.dev/kevin/golang-json-http/)
* [xml.NewDecoder(resp.Body).Decode Giving EOF Error _GOLang](https://stackoverflow.com/questions/24879587/xml-newdecoderresp-body-decode-giving-eof-error-golang)
* [写一个自己的Golang Module(Golang1.11以后版本支持)](https://www.jtianling.com/private-module-in-golang.html)
