/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-02-20 23:40:38
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-03-01 23:17:38
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

// AccessLogContent
func AccessLogContent(req *ghttp.Request) (content string) {
	var (
		scheme = "http"
		proto  = req.Header.Get("X-Forwarded-Proto")
	)
	if req.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	content = fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		req.Response.Status, req.Method, scheme, req.Host, req.URL.String(), req.Proto,
		float64(gtime.TimestampMilli()-req.EnterTime)/1000,
		req.GetClientIp(), req.Referer(), req.UserAgent(),
	)
	return
}

// ErrorLogContent
func ErrorLogContent(req *ghttp.Request, err error) (content string) {
	if err == nil {
		return
	}

	var (
		rCode         = gerror.Code(err)
		scheme        = "http"
		codeDetail    = rCode.Detail()
		proto         = req.Header.Get("X-Forwarded-Proto")
		codeDetailStr string
	)
	if req.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content = fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		req.Response.Status, req.Method, scheme, req.Host, req.URL.String(), req.Proto,
		float64(gtime.TimestampMilli()-req.EnterTime)/1000,
		req.GetClientIp(), req.Referer(), req.UserAgent(),
		rCode.Code(), rCode.Message(), codeDetailStr,
	)
	if stack := gerror.Stack(err); stack != "" {
		content += "\nStack:\n" + stack
	} else {
		content += ", " + err.Error()
	}
	return
}
