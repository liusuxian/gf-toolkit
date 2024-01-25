/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:15:17
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-25 21:12:59
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfmiddleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"github.com/liusuxian/gf-toolkit/gfresp"
	"net/http"
)

// HandlerResponse 自定义返回中间件，会默认打印错误
func HandlerResponse(req *ghttp.Request) {
	req.Middleware.Next()

	if req.Response.BufferLength() > 0 {
		return
	}

	var (
		err   = req.GetError()
		res   = req.GetHandlerResponse()
		rCode = gerror.Code(err)
	)
	if err != nil {
		if rCode == gcode.CodeNil {
			rCode = gcode.CodeInternalError
		}
	} else {
		if req.Response.Status > 0 && req.Response.Status != http.StatusOK {
			errText := http.StatusText(req.Response.Status)
			switch req.Response.Status {
			case http.StatusNotFound:
				rCode = gcode.CodeNotFound
			case http.StatusForbidden:
				rCode = gcode.CodeNotAuthorized
			default:
				rCode = gcode.CodeUnknown
			}
			err = gerror.NewCode(rCode, errText)
			req.SetError(err)
		} else {
			rCode = gcode.CodeOK
		}
	}
	// 打印错误
	if reqErr := req.GetError(); reqErr != nil {
		req.SetError(nil)
		gflogger.Errorf(req.GetCtx(), "HandlerResponse Error: %v", reqErr)
	}
	// 返回
	gfresp.Response{
		Code:    rCode.Code(),
		Message: rCode.Message(),
		Data:    res,
	}.Resp(req)
}
