/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-02-20 23:45:57
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-02-21 00:05:37
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfrobot

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/liusuxian/gf-toolkit/internal/utils"
	"github.com/liusuxian/go-toolkit/gtkrobot"
)

// FeishuRobot
type FeishuRobot struct {
	*gtkrobot.FeishuRobot
}

// SendErrorMessage 发送错误消息
func (fr *FeishuRobot) SendErrorMessage(req *ghttp.Request, e error) (err error) {
	req.SetError(e)
	err = fr.SendTextMessage(utils.ErrorLogContent(req))
	req.SetError(nil)
	return
}
