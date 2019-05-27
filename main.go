package main

import (
	"flag"
	"fmt"
	"github.com/funswe/flow"
	_ "github.com/funswe/flow-example/router"
	"log"
	"os"
	"path/filepath"
)

var (
	h           bool
	appName     string
	proxy       bool
	address     string
	viewPath    string
	logPath     string
	loggerLevel string
)

func init() {
	path, _ := filepath.Abs(".")
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&appName, "app-name", "flow-example", "set app name")
	flag.BoolVar(&proxy, "proxy", false, "set proxy mode")
	flag.StringVar(&address, "address", "localhost:9505", "set listen address")
	flag.StringVar(&viewPath, "view-path", filepath.Join(path, "views"), "set view path")
	flag.StringVar(&logPath, "log-path", filepath.Join(path, "logs"), "set log path")
	flag.StringVar(&loggerLevel, "logger-level", "debug", "set logger level")
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	// 设置app名称
	flow.SetAppName(appName)
	// 设置代理模式
	flow.SetProxy(proxy)
	// 设置监听地址
	flow.SetAddress(address)
	// 设置view地址
	flow.SetViewPath(viewPath)
	// 设置日志路径
	flow.SetLogPath(logPath)
	// 设置日志级别
	flow.SetLoggerLevel(loggerLevel)
	// 添加一个中间件
	flow.Use(func(ctx *flow.Context, next flow.Next) {
		fmt.Println("start...")
		next()
		fmt.Println("end...")
	})
	// 添加另一个中间件
	flow.Use(func(ctx *flow.Context, next flow.Next) {
		fmt.Println("this is a middleware")
		next() // 不能丢，否则不能正常执行handle
	})
	log.Fatal(flow.Run())
}
