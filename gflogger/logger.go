/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 20:59:43
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-04-11 17:55:45
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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

// Log 获取 log 对象
func Log(name ...string) (logger *glog.Logger) {
	if len(name) > 0 && name[0] != "" {
		return g.Log(name[0]).Line()
	}
	return g.Log().Line()
}

// AccessLog 获取 access log 对象
func AccessLog() (logger *glog.Logger) {
	return Log("access")
}

// ErrorLog 获取 error log 对象
func ErrorLog() (logger *glog.Logger) {
	return Log("error")
}

func Print(ctx context.Context, v ...any) {
	AccessLog().Stack(false, 1).Print(ctx, v...)
}

func Printf(ctx context.Context, format string, v ...any) {
	AccessLog().Stack(false, 1).Printf(ctx, format, v...)
}

func Info(ctx context.Context, v ...any) {
	AccessLog().Stack(false, 1).Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...any) {
	AccessLog().Stack(false, 1).Infof(ctx, format, v...)
}

func Debug(ctx context.Context, v ...any) {
	AccessLog().Stack(false, 1).Debug(ctx, v...)
}

func Debugf(ctx context.Context, format string, v ...any) {
	AccessLog().Stack(false, 1).Debugf(ctx, format, v...)
}

func Notice(ctx context.Context, v ...any) {
	AccessLog().Stack(false, 1).Notice(ctx, v...)
}

func Noticef(ctx context.Context, format string, v ...any) {
	AccessLog().Stack(false, 1).Noticef(ctx, format, v...)
}

func Warning(ctx context.Context, v ...any) {
	ErrorLog().Stack(false, 1).Warning(ctx, v...)
}

func Warningf(ctx context.Context, format string, v ...any) {
	ErrorLog().Stack(false, 1).Warningf(ctx, format, v...)
}

func Error(ctx context.Context, v ...any) {
	ErrorLog().Stack(false, 1).Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	ErrorLog().Stack(false, 1).Errorf(ctx, format, v...)
}

func Fatal(ctx context.Context, v ...any) {
	ErrorLog().Stack(false, 1).Fatal(ctx, v...)
}

func Fatalf(ctx context.Context, format string, v ...any) {
	ErrorLog().Stack(false, 1).Fatalf(ctx, format, v...)
}

func Critical(ctx context.Context, v ...any) {
	ErrorLog().Stack(false, 1).Critical(ctx, v...)
}

func Criticalf(ctx context.Context, format string, v ...any) {
	ErrorLog().Stack(false, 1).Criticalf(ctx, format, v...)
}

func Panic(ctx context.Context, v ...any) {
	ErrorLog().Stack(false, 1).Panic(ctx, v...)
}

func Panicf(ctx context.Context, format string, v ...any) {
	ErrorLog().Stack(false, 1).Panicf(ctx, format, v...)
}

// AccessContent
func AccessContent(req *ghttp.Request) (content string) {
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
		float64(gtime.TimestampMilli()-req.EnterTime.TimestampMilli())/1000,
		req.GetClientIp(), req.Referer(), req.UserAgent(),
	)
	return
}

// ErrorContent
func ErrorContent(req *ghttp.Request, err error) (content string) {
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
		float64(gtime.TimestampMilli()-req.EnterTime.TimestampMilli())/1000,
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

// PrintAccess
func PrintAccess(req *ghttp.Request, skip ...int) {
	AccessLog().Stack(false, skip...).Debug(req.Context(), AccessContent(req))
}

// PrintError
func PrintError(req *ghttp.Request, err error, skip ...int) {
	ErrorLog().Stack(false, skip...).Error(req.Context(), ErrorContent(req, err))
}
