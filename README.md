# flow-example
this is a example of flow

# 安装
go get -u github.com/funswe/flow-example

# 运行（前提是GOPATH的bin目录在PATH里）
查看帮助
```
flow-example -h
Usage of flow-example:
  -address string
    	set listen address (default "localhost:9505")
  -app-name string
    	set app name (default "flow-example")
  -h	this help
  -log-path string
    	set log path (default "当前路径下的logs目录")
  -proxy
    	set proxy mode
  -view-path string
    	set view path (default "当前路径下的views目录")
```

执行命令
```
flow-example --app-name <appName> --address <address> --view-path $GOPATH/src/github.com/funswe/flow-example --log-path <logPath>
```

# 查看效果
- 浏览器访问http://<设置的address>/json
- 浏览器访问http://<设置的address>/render
- 浏览器访问http://<设置的address>/parse/flow-example?age=30
