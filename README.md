# ttshitu

方便在 golang 项目中使用图鉴的图片识别功能。

[图鉴开发文档](http://www.ttshitu.com/docs/index.html)

## HowTo

```golang
package main

import "github.com/GiaoGiaoCat/ttshitu"

func main() {
  req := &ttshitu.RequestBody{Username: "Rocky", Password: "5454", Image: "xxxxx"}
  // 调用 GetCode 方法，获取识别后的验证码，如果无法识别则返回空字符串。
  code := ttshitu.GetCode(req)
  fmt.Println(code)
}
```

* Username 和 Password 是在图鉴申请的账号密码。
* Image 是图片的 Base64 编码。

## 参考

* https://medium.com/rungo/making-external-http-requests-in-go-eb4c015f8839
* https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
* https://stackoverflow.com/questions/24879587/xml-newdecoderresp-body-decode-giving-eof-error-golang
* [How to Parse a JSON Request Body in Go](https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body)
