/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 20:59:43
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-25 21:01:37
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gflogger

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
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
func defaultLog(name ...string) *glog.Logger {
	return Log(1, name...)
}
