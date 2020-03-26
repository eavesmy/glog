# GLog
符合 Journalctl 输出标准，配合 journalctl 使用。日志等级使用 journalctl 等级。

# Install
```go get github.com/eavesmy/glog```

# Usage
```golang
package main

import "github.com/eavesmy/glog"
import "os"

func main(){

	// add prefix. default: "". Add debug mode: true/false.
	// If mode = true,then stop log.
	log := glog.New("[Test]",os.Getenv("PRODUCTION_MODE")).

	// Open log by hand and set whitch level log could be shown.
	glog.Able(glog.Err,glog.Info)

	log.Info("Is info log")
	Info.Err("Error!")
	Info.Debug("Is debug msg.")
	
	Info.Time("Hi")
	
	// setInterval
	// unit: unixnano

	Info.TimeEnd("Hi")
}

```
