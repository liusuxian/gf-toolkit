/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-02-16 00:37:45
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"github.com/liusuxian/go-toolkit/gtkjson"
	"net/http"
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

// ResFail 返回失败
func ResFail(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	req.Response.WriteJson(Fail(rCode.Code(), rCode.Message(), data...))
}

// ResFailExit 返回失败并退出
func ResFailExit(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	req.Response.WriteJsonExit(Fail(rCode.Code(), rCode.Message(), data...))
}

// ResFailPrintErr 返回失败，默认打印错误日志
func ResFailPrintErr(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	req.Response.WriteJson(Fail(rCode.Code(), rCode.Message(), data...))
	req.SetError(err)
	gflogger.HandlerErrorLog(req, 2)
	req.SetError(nil)
}

// ResFailPrintErrExit 返回失败并退出，默认打印错误日志
func ResFailPrintErrExit(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	req.Response.WriteJson(Fail(rCode.Code(), rCode.Message(), data...))
	req.SetError(err)
	gflogger.HandlerErrorLog(req, 2)
	req.SetError(nil)
	req.Exit()
}

// ResSucc 返回成功
func ResSucc(req *ghttp.Request, data any) {
	req.Response.WriteJson(Succ(data))
}

// RespSuccExit 返回成功并退出
func RespSuccExit(req *ghttp.Request, data any) {
	req.Response.WriteJsonExit(Succ(data))
}

// ResJsonViaSSE 返回`JSON`流式数据
func ResJsonViaSSE(req *ghttp.Request, data ...any) {
	// 设置`SSE`的`Content-Type`
	req.Response.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	// 准备发送的数据
	var jsonData string
	if len(data) > 0 {
		// 序列化数据为`JSON`字符串
		jsonData = gtkjson.MustJsonMarshal(Succ(data[0]))
	} else {
		// 如果没有数据，则发送表示结束的消息
		jsonData = gtkjson.MustJsonMarshal(Succ("EOF"))
	}
	// 按`SSE`格式发送数据：'data: <jsonData>\n\n'
	fmt.Fprintf(req.Response.ResponseWriter, "data: %s\n\n", jsonData)
	// 确保即时发送数据
	req.Response.Flush()
	// 如果没有数据要发送，结束请求处理
	if len(data) == 0 {
		req.Exit()
	}
}

// Redirect 重定向
func Redirect(req *ghttp.Request, link string) {
	req.Response.Header().Set("Location", link)
	req.Response.WriteHeader(302)
}
