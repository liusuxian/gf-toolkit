/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-25 20:21:25
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-25 20:33:48
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gferror_test

import (
	"errors"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/liusuxian/gf-toolkit/gferror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleError(t *testing.T) {
	var (
		rCode     gcode.Code
		unwrapErr error
		assert    = assert.New(t)
	)
	rCode, unwrapErr = gferror.HandleError(gerror.NewCode(gcode.CodeInternalError))
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Nil(unwrapErr)
	rCode, unwrapErr = gferror.HandleError(gerror.NewCode(gcode.CodeInternalError, "i am error"))
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Nil(unwrapErr)
	rCode, unwrapErr = gferror.HandleError(gerror.WrapCode(gcode.CodeInternalError, errors.New("i am error")))
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Equal(errors.New("i am error"), unwrapErr)
	rCode, unwrapErr = gferror.HandleError(gerror.WrapCode(gcode.CodeInternalError, nil))
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Nil(unwrapErr)
}
