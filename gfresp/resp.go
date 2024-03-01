/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-03-01 23:41:57
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
