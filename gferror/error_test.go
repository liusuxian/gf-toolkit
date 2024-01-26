/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-25 20:21:25
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-26 22:11:54
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
		err       error
		rCode     gcode.Code
		unwrapErr error
		assert    = assert.New(t)
	)
	
	err = gerror.NewCode(gcode.CodeInternalError)
	rCode, unwrapErr = gferror.HandlerError(err)
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Equal("Internal Error", err.Error())
	assert.Nil(unwrapErr)

	err = gerror.NewCode(gcode.CodeInternalError, "i am error")
	rCode, unwrapErr = gferror.HandlerError(err)
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Equal("i am error", err.Error())
	assert.Nil(unwrapErr)

	err = gerror.WrapCode(gcode.CodeInternalError, errors.New("i am error"))
	rCode, unwrapErr = gferror.HandlerError(err)
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Equal("Internal Error: i am error", err.Error())
	assert.Equal(errors.New("i am error"), unwrapErr)

	err = gerror.WrapCode(gcode.CodeInternalError, nil)
	rCode, unwrapErr = gferror.HandlerError(err)
	assert.Equal(gcode.CodeInternalError, rCode)
	assert.Nil(err)
	assert.Nil(unwrapErr)
}
