/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-02-22 13:46:33
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/liusuxian/gf-toolkit/gfrobot"
	"github.com/liusuxian/go-toolkit/gtkjson"
	"net/http"
	"sync"
)

var (
	once        sync.Once
	feishuRobot *gfrobot.FeishuRobot // 飞书机器人
)

// Response 通用响应数据结构
type Response struct {
	Code    int    `json:"code"    dc:"错误码(0:成功, 非0:错误)"`   // 错误码(0:成功, 非0:错误)
	Message string `json:"message" dc:"错误消息"`               // 错误消息
	Data    any    `json:"data"    dc:"根据API定义，对特定请求的结果数据"` // 根据`API`定义，对特定请求的结果数据
}

// Success 判断是否成功
func (resp Response) Success() (ok bool) {
	return resp.Code == gcode.CodeOK.Code()
}

// DataString 获取`Data`转字符串
func (resp Response) DataString() (data string) {
	return gconv.String(resp.Data)
}

// DataInt 获取`Data`转`Int`
func (resp Response) DataInt() (data int) {
	return gconv.Int(resp.Data)
}

// GetString 获取`Data`值转字符串
func (resp Response) GetString(key string) (data string) {
	return gconv.String(resp.Get(key))
}

// GetInt 获取`Data`值转`Int`
func (resp Response) GetInt(key string) (data int) {
	return gconv.Int(resp.Get(key))
}

// Get 获取`Data`值
func (resp Response) Get(key string) (data *gvar.Var) {
	m := gconv.Map(resp.Data)
	if m == nil {
		return nil
	}
	return gvar.New(m[key])
}

// Json
func (resp Response) Json() (str string) {
	b, _ := json.Marshal(resp)
	return string(b)
}

// Succ 成功
func Succ(data any) (resp Response) {
	return Response{Code: gcode.CodeOK.Code(), Message: gcode.CodeOK.Message(), Data: data}
}

// Fail 失败
func Fail(code int, msg string, data ...any) (resp Response) {
	var rData any
	if len(data) > 0 {
		rData = data[0]
	}
	return Response{Code: code, Message: msg, Data: rData}
}

// Unauthorized 认证失败
func Unauthorized(msg string, data any) (resp Response) {
	return Response{Code: http.StatusUnauthorized, Message: msg, Data: data}
}

// NewFeishuRobot 新建飞书机器人
func NewFeishuRobot(webHookURL string) {
	once.Do(func() {
		feishuRobot = gfrobot.NewFeishuRobot(webHookURL)
	})
}

// ResFail 返回失败
func ResFail(req *ghttp.Request, err error, isExit bool, data ...any) {
	if feishuRobot != nil {
		go func() {
			feishuRobot.SendErrMessage(req, err)
		}()
	}
	rCode := gerror.Code(err)
	if isExit {
		req.Response.WriteJsonExit(Fail(rCode.Code(), rCode.Message(), data...))
	} else {
		req.Response.WriteJson(Fail(rCode.Code(), rCode.Message(), data...))
	}
}

// ResFailByCtx 返回失败
func ResFailByCtx(ctx context.Context, err error, isExit bool, data ...any) {
	ResFail(ghttp.RequestFromCtx(ctx), err, isExit, data...)
}

// ResFailPrintErr 返回失败，默认打印错误日志
func ResFailPrintErr(req *ghttp.Request, err error, isExit bool, data ...any) {
	resFailPrintErr(req, err, isExit, data...)
}

// ResFailPrintErrByCtx 返回失败，默认打印错误日志
func ResFailPrintErrByCtx(ctx context.Context, err error, isExit bool, data ...any) {
	resFailPrintErr(ghttp.RequestFromCtx(ctx), err, isExit, data...)
}

// ResSucc 返回成功
func ResSucc(req *ghttp.Request, data any, isExit bool) {
	if isExit {
		req.Response.WriteJsonExit(Succ(data))
	} else {
		req.Response.WriteJson(Succ(data))
	}
}

// ResSuccByCtx 返回成功
func ResSuccByCtx(ctx context.Context, data any, isExit bool) {
	ResSucc(ghttp.RequestFromCtx(ctx), data, isExit)
}

// ResFailStream 返回流式数据失败
func ResFailStream(req *ghttp.Request, err error, isExit bool, data ...any) {
	if feishuRobot != nil {
		go func() {
			feishuRobot.SendErrMessage(req, err)
		}()
	}
	// 设置`SSE`的`Content-Type`
	req.Response.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	// 序列化数据为`JSON`字符串
	rCode := gerror.Code(err)
	jsonData := gtkjson.MustJsonMarshal(Fail(rCode.Code(), rCode.Message(), data...))
	// 发送数据：'<jsonData>\n'
	fmt.Fprintf(req.Response.ResponseWriter, "%s\n", jsonData)
	// 确保即时发送数据
	req.Response.Flush()
	// 结束请求处理
	if isExit {
		req.Exit()
	}
}

// ResFailStreamByCtx 返回流式数据失败
func ResFailStreamByCtx(ctx context.Context, err error, isExit bool, data ...any) {
	ResFailStream(ghttp.RequestFromCtx(ctx), err, isExit, data...)
}

// ResFailStreamPrintErr 返回流式数据失败，默认打印错误日志
func ResFailStreamPrintErr(req *ghttp.Request, err error, isExit bool, data ...any) {
	resFailStreamPrintErr(req, err, isExit, data...)
}

// ResFailStreamPrintErrByCtx 返回流式数据失败，默认打印错误日志
func ResFailStreamPrintErrByCtx(ctx context.Context, err error, isExit bool, data ...any) {
	resFailStreamPrintErr(ghttp.RequestFromCtx(ctx), err, isExit, data...)
}

// ResSuccStream 返回流式数据成功
func ResSuccStream(req *ghttp.Request, data any, isExit bool) {
	// 设置`SSE`的`Content-Type`
	req.Response.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	// 序列化数据为`JSON`字符串
	jsonData := gtkjson.MustJsonMarshal(Succ(data))
	// 发送数据：'<jsonData>\n'
	fmt.Fprintf(req.Response.ResponseWriter, "%s\n", jsonData)
	// 确保即时发送数据
	req.Response.Flush()
	// 结束请求处理
	if isExit {
		req.Exit()
	}
}

// ResSuccStreamByCtx 返回流式数据成功
func ResSuccStreamByCtx(ctx context.Context, data any, isExit bool) {
	ResSuccStream(ghttp.RequestFromCtx(ctx), data, isExit)
}

// Redirect 重定向
func Redirect(req *ghttp.Request, link string) {
	req.Response.Header().Set("Location", link)
	req.Response.WriteHeader(302)
}
