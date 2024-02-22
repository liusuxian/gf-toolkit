/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-02-20 23:45:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-02-22 12:54:33
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfrobot

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/gf-toolkit/internal/utils"
	"github.com/liusuxian/go-toolkit/gtkrobot"
)

// FeishuRobot
type FeishuRobot struct {
	*gtkrobot.FeishuRobot
}

// NewFeishuRobot 新建飞书机器人
func NewFeishuRobot(webHookURL string) (fr *FeishuRobot) {
	return &FeishuRobot{
		gtkrobot.NewFeishuRobot(webHookURL),
	}
}

// SendErrMessage 发送错误消息
func (fr *FeishuRobot) SendErrMessage(req *ghttp.Request, e error) (err error) {
	req.SetError(e)
	err = fr.SendTextMessage(utils.ErrorLogContent(req))
	req.SetError(nil)
	return
}

// SendErrMessageByCtx 发送错误消息
func (fr *FeishuRobot) SendErrMessageByCtx(ctx context.Context, e error) (err error) {
	return fr.SendErrMessage(ghttp.RequestFromCtx(ctx), e)
}
