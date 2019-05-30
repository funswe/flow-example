package controllers

import (
	"github.com/funswe/flow"
)

func init() {
	flow.ALL("/json", json())
	flow.GET("/render", render())
	flow.GET("/parse/:name", parse())
	flow.GET("/download", download())
}

// json方式返回
func json() flow.Handler {
	return func(ctx *flow.Context) {
		ctx.Json(map[string]interface{}{
			"hello": "world",
		})
	}
}

// 渲染html返回
func render() flow.Handler {
	return func(ctx *flow.Context) {
		ctx.Render("index.html", map[string]interface{}{
			"hello": "world",
		})
	}
}

type parseReq struct {
	Name string
	Age  int32 `json:"age,string"`
}

// 参数解析
// /parse/zhang?age=29
func parse() flow.Handler {
	return func(ctx *flow.Context) {
		params := &parseReq{}
		ctx.Parse(params)
		ctx.Json(map[string]interface{}{
			"name":  params.Name,
			"age":   params.Age,
			"name1": ctx.GetParam("name"),
			"age1":  ctx.GetParam("age"),
			"name2": ctx.GetParamDefault("name2", "test"),
		})
	}
}

func download() flow.Handler {
	return func(ctx *flow.Context) {
		ctx.Download("flink-1.8.0-src.tgz")
	}
}
