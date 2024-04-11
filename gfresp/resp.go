/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-04-11 17:55:29
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/go-toolkit/gtkresp"
	"time"
)

// RespFail 返回失败
func RespFail(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	gtkresp.RespFail(req.Response.Writer, rCode.Code(), rCode.Message(), data...)
}

// RespSucc 返回成功
func RespSucc(req *ghttp.Request, data any) {
	gtkresp.RespSucc(req.Response.Writer, data)
}

// RespSSEFail 返回流式数据失败
func RespSSEFail(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	gtkresp.RespSSEFail(req.Response.Writer, rCode.Code(), rCode.Message(), data...)
}

// RespSSESucc 返回流式数据成功
func RespSSESucc(req *ghttp.Request, data any) {
	gtkresp.RespSSESucc(req.Response.Writer, data)
}

// Redirect 重定向
func Redirect(req *ghttp.Request, link string) {
	gtkresp.Redirect(req.Response.Writer, link)
}

// WriteSuccMessage 写成功响应消息
func WriteSuccMessage(ws *ghttp.WebSocket, messageType int, data any) (err error) {
	return gtkresp.WriteSuccMessage(ws.Conn, messageType, data)
}

// WriteFailMessage 写失败响应消息
func WriteFailMessage(req *ghttp.Request, ws *ghttp.WebSocket, messageType int, err error, data ...any) (e error) {
	rCode := gerror.Code(err)
	return gtkresp.WriteFailMessage(ws.Conn, messageType, rCode.Code(), rCode.Message(), data...)
}

// WriteMessage 写任意消息
func WriteMessage(ws *ghttp.WebSocket, messageType int, data any) (err error) {
	return gtkresp.WriteMessage(ws.Conn, messageType, data)
}

// WriteControl 使用给定的截止时间写入一个控制消息
//
//	允许的消息类型包括 `CloseMessage`，`PingMessage`，`PongMessage`
func WriteControl(ws *ghttp.WebSocket, messageType int, data any, deadline time.Time) (err error) {
	return gtkresp.WriteControl(ws.Conn, messageType, data, deadline)
}
