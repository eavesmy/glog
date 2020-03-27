# GLog
只是个输出，log 实现，可重定向输出及添加prefix。配合journalctl做日志管理。
符合 Journalctl 输出标准，配合 journalctl 使用。日志等级使用 journalctl 等级。

# Notice
如果需要一个更多功能的 logger 组件，请使用其他库，这个项目只是一个简单的输出，没有其他的附加功能。

# Install
```go get github.com/eavesmy/glog```

# Usage
```golang
package main

import "github.com/eavesmy/glog"
import "os"

func main(){

	// add prefix. default: "".
	log := glog.New("[Test]",os.Getenv("PRODUCTION_MODE")).

	// Open log by hand and set whitch log with level could be shown.
	// 'emerg' 'alert' 'crit' 'warn' 'notice' 'info' 'debug'
	// You can write condition to env then you can use it like:
	// log.Unable(os.GetEnv("PRODUCT_UNLOG")) // PRODUCT_UNLOG=alert,crit
	log.Unable('alert,crit')

	log.Info("Is info log")
	log.Err("Error!")
	log.Debug("Is a debug msg.")
	
	log.Time("Hi")
	
	// sleep
	// unit: unixnano

	log.TimeEnd("Hi")

	log.SetLevel(glog.ERR).Printf("%d",123)
}

```
