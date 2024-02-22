/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-02-22 13:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-02-22 13:46:16
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"github.com/liusuxian/go-toolkit/gtkjson"
)

// resFailPrintErr 返回失败，默认打印错误日志
func resFailPrintErr(req *ghttp.Request, err error, isExit bool, data ...any) {
	if feishuRobot != nil {
		go func() {
			feishuRobot.SendErrMessage(req, err)
		}()
	}
	rCode := gerror.Code(err)
	req.Response.WriteJson(Fail(rCode.Code(), rCode.Message(), data...))
	req.SetError(err)
	gflogger.HandlerErrorLog(req, 3)
	req.SetError(nil)
	if isExit {
		req.Exit()
	}
}

// resFailStreamPrintErr 返回流式数据失败，默认打印错误日志
func resFailStreamPrintErr(req *ghttp.Request, err error, isExit bool, data ...any) {
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
	// 打印错误日志
	req.SetError(err)
	gflogger.HandlerErrorLog(req, 3)
	req.SetError(nil)
	// 结束请求处理
	if isExit {
		req.Exit()
	}
}
