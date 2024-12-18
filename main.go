package main

import (
	"fmt"
	"github.com/shadow1ng/fscan/Common"
	"github.com/shadow1ng/fscan/Config"
	"github.com/shadow1ng/fscan/Plugins"
	"time"
)

func main() {
	start := time.Now()
	var Info Config.HostInfo
	Common.Flag(&Info)
	Common.Parse(&Info)
	Plugins.Scan(Info)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", time.Since(start))
}
