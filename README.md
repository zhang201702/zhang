# zhang
基于goframe的web

# Quick Start
默认80端口，默认html文件夹为web页面文件
```go
package main

import (
	"github.com/zhang201702/zhang"
)

func main(){
	server := zhang.Default()
	server.Run()
}
```
