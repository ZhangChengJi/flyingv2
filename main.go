package main

import (
	"flyingv2/cmd/rootcmd"
)

// @title flyingv2
// @version 2.0
// @description etcd 版分布式配置中心

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /
func main() {
	rootcmd.Execute()

}
