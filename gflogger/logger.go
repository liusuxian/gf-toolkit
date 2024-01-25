/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 20:59:43
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-26 01:17:16
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gflogger

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// Log 获取log对象
func Log(skip int, name ...string) *glog.Logger {
	if len(name) > 0 && name[0] != "" {
		return g.Log(name[0]).Skip(skip).Line()
	}
	return g.Log().Skip(skip).Line()
}

func Print(ctx context.Context, v ...any) {
	defaultLog("access").Print(ctx, v...)
}

func Printf(ctx context.Context, format string, v ...any) {
	defaultLog("access").Printf(ctx, format, v...)
}

func Info(ctx context.Context, v ...any) {
	defaultLog("access").Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...any) {
	defaultLog("access").Infof(ctx, format, v...)
}

func Debug(ctx context.Context, v ...any) {
	defaultLog("access").Debug(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...any) {
	defaultLog("access").Debugf(ctx, format, v...)
}

func Notice(ctx context.Context, v ...any) {
	defaultLog("access").Notice(ctx, v...)
}

func Noticef(ctx context.Context, format string, v ...any) {
	defaultLog("access").Noticef(ctx, format, v...)
}

func Warning(ctx context.Context, v ...any) {
	defaultLog("error").Warning(ctx, v...)
}

func Warningf(ctx context.Context, format string, v ...any) {
	defaultLog("error").Warningf(ctx, format, v...)
}

func Error(ctx context.Context, v ...any) {
	defaultLog("error").Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	defaultLog("error").Errorf(ctx, format, v...)
}

func Fatal(ctx context.Context, v ...any) {
	defaultLog("error").Fatal(ctx, v...)
}

func Fatalf(ctx context.Context, format string, v ...any) {
	defaultLog("error").Fatalf(ctx, format, v...)
}

func Critical(ctx context.Context, v ...any) {
	defaultLog("error").Critical(ctx, v...)
}

func Criticalf(ctx context.Context, format string, v ...any) {
	defaultLog("error").Criticalf(ctx, format, v...)
}

func Panic(ctx context.Context, v ...any) {
	defaultLog("error").Panic(ctx, v...)
}

func Panicf(ctx context.Context, format string, v ...any) {
	defaultLog("error").Panicf(ctx, format, v...)
}

// defaultLog 默认log对象
func defaultLog(name ...string) (logger *glog.Logger) {
	logger = Log(1, name...)
	if logger == nil {
		return
	}
	logger.SetStack(false)
	return
}

// HandleAccessLog
func HandleAccessLog(req *ghttp.Request, skip int) {
	logger := Log(skip, "access")
	if logger == nil {
		return
	}

	var (
		scheme = "http"
		proto  = req.Header.Get("X-Forwarded-Proto")
	)
	if req.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		req.Response.Status, req.Method, scheme, req.Host, req.URL.String(), req.Proto,
		float64(req.LeaveTime-req.EnterTime)/1000,
		req.GetClientIp(), req.Referer(), req.UserAgent(),
	)
	logger.Debug(req.Context(), content)
}

// HandleErrorLog
func HandleErrorLog(req *ghttp.Request, skip int, err error) {
	if err == nil {
		return
	}

	logger := Log(skip, "error")
	if logger == nil {
		return
	}
	logger.SetStack(false)

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
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		req.Response.Status, req.Method, scheme, req.Host, req.URL.String(), req.Proto,
		float64(req.LeaveTime-req.EnterTime)/1000,
		req.GetClientIp(), req.Referer(), req.UserAgent(),
		rCode.Code(), rCode.Message(), codeDetailStr,
	)
	if stack := gerror.Stack(err); stack != "" {
		content += "\nStack:\n" + stack
	} else {
		content += ", " + err.Error()
	}
	req.SetError(nil)
	logger.Error(req.Context(), content)
}
