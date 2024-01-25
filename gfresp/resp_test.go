/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:04:44
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-25 16:50:17
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfresp_test

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRespError(t *testing.T) {
	var (
		err    error
		assert = assert.New(t)
	)
	err = gerror.NewCode(gcode.CodeInternalError)
	assert.True(gerror.HasCode(err, gcode.CodeInternalError))
	err = gerror.NewCode(gcode.CodeInternalError, "i am error")
	assert.True(gerror.HasCode(err, gcode.CodeInternalError))
	rCode := gerror.Code(nil)
	assert.Equal(gcode.CodeNil, rCode)
}
