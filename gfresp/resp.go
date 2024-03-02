/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-03-02 18:18:13
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"github.com/liusuxian/gf-toolkit/internal/utils"
	"github.com/liusuxian/go-toolkit/gtkresp"
	"github.com/liusuxian/go-toolkit/gtkrobot"
	"sync"
	"time"
)

var (
	once        sync.Once
	feishuRobot *gtkrobot.FeishuRobot // 飞书机器人
)

// NewFeishuRobot 新建飞书机器人
func NewFeishuRobot(webHookURL string) {
	once.Do(func() {
		feishuRobot = gtkrobot.NewFeishuRobot(webHookURL)
	})
}

// sendFeishuRobot 发送飞书机器人
func sendFeishuRobot(req *ghttp.Request, err error) {
	if feishuRobot != nil && err != nil {
		go func() {
			feishuRobot.SendTextMessage(req.GetCtx(), utils.ErrorLogContent(req, err))
		}()
	}
}

// RespFail 返回失败
func RespFail(req *ghttp.Request, err error, data ...any) {
	sendFeishuRobot(req, err) // 发送飞书机器人
	rCode := gerror.Code(err)
	gtkresp.RespFail(req.Response.Writer, rCode.Code(), rCode.Message(), data...)
}

// RespFailPrintErr 返回失败，默认打印错误日志
func RespFailPrintErr(req *ghttp.Request, err error, data ...any) {
	RespFail(req, err, data...)
	gflogger.HandlerErrorLog(req, err, 2)
}

// RespSucc 返回成功
func RespSucc(req *ghttp.Request, data any) {
	gtkresp.RespSucc(req.Response.Writer, data)
}

// RespSSEFail 返回流式数据失败
func RespSSEFail(req *ghttp.Request, err error, data ...any) {
	sendFeishuRobot(req, err) // 发送飞书机器人
	rCode := gerror.Code(err)
	gtkresp.RespSSEFail(req.Response.Writer, rCode.Code(), rCode.Message(), data...)
}

// RespSSEFailPrintErr 返回流式数据失败，默认打印错误日志
func RespSSEFailPrintErr(req *ghttp.Request, err error, data ...any) {
	RespSSEFail(req, err, data...)
	gflogger.HandlerErrorLog(req, err, 2)
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
func WriteFailMessage(req *ghttp.Request, ws *ghttp.WebSocket, err error, messageType int, data ...any) (e error) {
	sendFeishuRobot(req, err) // 发送飞书机器人
	rCode := gerror.Code(err)
	return gtkresp.WriteFailMessage(ws.Conn, messageType, rCode.Code(), rCode.Message(), data...)
}

// WriteFailMessagePrintErr 写失败响应消息，默认打印错误日志
func WriteFailMessagePrintErr(req *ghttp.Request, ws *ghttp.WebSocket, err error, messageType int, data ...any) (e error) {
	e = WriteFailMessage(req, ws, err, messageType, data...)
	gflogger.HandlerErrorLog(req, err, 2)
	return
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
