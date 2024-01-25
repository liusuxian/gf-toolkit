/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-26 01:06:29
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/liusuxian/gf-toolkit/gferror"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"net/http"
)

// Response 通用响应数据结构
type Response struct {
	Code    int         `json:"code"    dc:"错误码(0:成功, 非0:错误)"`     // 错误码(0:成功, 非0:错误)
	Message string      `json:"message" dc:"错误消息"`                 // 错误消息
	Data    interface{} `json:"data"    dc:"根据 API 定义，对特定请求的结果数据"` // 根据 API 定义，对特定请求的结果数据
}

// Success 判断是否成功
func (resp Response) Success() (ok bool) {
	return resp.Code == gcode.CodeOK.Code()
}

// DataString 获取Data转字符串
func (resp Response) DataString() (data string) {
	return gconv.String(resp.Data)
}

// DataInt 获取Data转Int
func (resp Response) DataInt() (data int) {
	return gconv.Int(resp.Data)
}

// GetString 获取Data值转字符串
func (resp Response) GetString(key string) (data string) {
	return gconv.String(resp.Get(key))
}

// GetInt 获取Data值转Int
func (resp Response) GetInt(key string) (data int) {
	return gconv.Int(resp.Get(key))
}

// Get 获取Data值
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

// Resp 响应数据返回
func (resp Response) Resp(req *ghttp.Request) {
	req.Response.WriteJson(resp)
}

// RespCtx 响应数据返回
func (resp Response) RespCtx(ctx context.Context) {
	req := g.RequestFromCtx(ctx)
	req.Response.WriteJson(resp)
}

// RespExit 响应数据返回并退出
func (resp Response) RespExit(req *ghttp.Request) {
	req.Response.WriteJsonExit(resp)
}

// RespCtxExit 响应数据返回并退出
func (resp Response) RespCtxExit(ctx context.Context) {
	req := g.RequestFromCtx(ctx)
	req.Response.WriteJsonExit(resp)
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

// RespFail 返回失败
func RespFail(req *ghttp.Request, err error, data ...any) {
	gflogger.HandleErrorLog(req, "RespFail Error: ", err)
	rCode, _ := gferror.HandleError(err)
	Fail(rCode.Code(), rCode.Message(), data...).Resp(req)
}

// RespFailCtx 返回失败
func RespFailCtx(ctx context.Context, err error, data ...any) {
	gflogger.HandleErrorLog(g.RequestFromCtx(ctx), "RespFailCtx Error: ", err)
	rCode, _ := gferror.HandleError(err)
	Fail(rCode.Code(), rCode.Message(), data...).RespCtx(ctx)
}

// RespFailExit 返回失败并退出
func RespFailExit(req *ghttp.Request, err error, data ...any) {
	gflogger.HandleErrorLog(req, "RespFailExit Error: ", err)
	rCode, _ := gferror.HandleError(err)
	Fail(rCode.Code(), rCode.Message(), data...).RespExit(req)
}

// RespFailCtxExit 返回失败并退出
func RespFailCtxExit(ctx context.Context, err error, data ...any) {
	gflogger.HandleErrorLog(g.RequestFromCtx(ctx), "RespFailCtxExit Error: ", err)
	rCode, _ := gferror.HandleError(err)
	Fail(rCode.Code(), rCode.Message(), data...).RespCtxExit(ctx)
}

// RespSucc 返回成功
func RespSucc(req *ghttp.Request, data any) {
	Succ(data).Resp(req)
}

// RespSuccCtx 返回成功
func RespSuccCtx(ctx context.Context, data any) {
	Succ(data).RespCtx(ctx)
}

// RespSuccExit 返回成功并退出
func RespSuccExit(req *ghttp.Request, data any) {
	Succ(data).RespExit(req)
}

// RespSuccCtxExit 返回成功并退出
func RespSuccCtxExit(ctx context.Context, data any) {
	Succ(data).RespCtxExit(ctx)
}

// Redirect 重定向
func Redirect(req *ghttp.Request, link string) {
	req.Response.Header().Set("Location", link)
	req.Response.WriteHeader(302)
}
