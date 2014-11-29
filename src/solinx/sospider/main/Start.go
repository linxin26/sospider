package main

import (
	"fmt"
	"solinx/sospider/Request"
	"runtime"
)

/**
SoSpider入口
author:linx
date:2014-11-19 22:00:00
 */
func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("启动SoSpider")
	req := Request.HttpRequest{"request"}
	req.Requsturl()

}




