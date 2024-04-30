/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-04-30 13:34:43
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/go-toolkit/gtkresp"
)

// RespFail 返回失败
func RespFail(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	gtkresp.RespFail(req.Response.BufferWriter, rCode.Code(), rCode.Message(), data...)
}

// RespSucc 返回成功
func RespSucc(req *ghttp.Request, data any) {
	gtkresp.RespSucc(req.Response.BufferWriter, data)
}

// RespSSEFail 返回流式数据失败
func RespSSEFail(req *ghttp.Request, err error, data ...any) {
	rCode := gerror.Code(err)
	gtkresp.RespSSEFail(req.Response.BufferWriter, rCode.Code(), rCode.Message(), data...)
}

// RespSSESucc 返回流式数据成功
func RespSSESucc(req *ghttp.Request, data any) {
	gtkresp.RespSSESucc(req.Response.BufferWriter, data)
}

// Redirect 重定向
func Redirect(req *ghttp.Request, link string) {
	gtkresp.Redirect(req.Response.BufferWriter, link)
}
