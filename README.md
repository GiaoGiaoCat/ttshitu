# ttshitu

方便在 golang 项目中使用图鉴的图片识别功能

## HowTo

```golang
package main

import "github.com/GiaoGiaoCat/ttshitu"

func main() {
  req := &ttshitu.RequestBody{Username: "Rocky", Password: "5454", Image: "xxxxx"}

  ttshitu.Print(req)
}
```
