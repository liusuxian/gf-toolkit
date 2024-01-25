/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-25 20:21:25
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-26 01:19:31
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gferror

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// HandlerError 处理错误
func HandlerError(err error) (rCode gcode.Code, unwrapErr error) {
	rCode = gerror.Code(err)
	if rCode == gcode.CodeNil {
		rCode = gcode.CodeInternalError
	}
	unwrapErr = gerror.Unwrap(err)
	return
}
