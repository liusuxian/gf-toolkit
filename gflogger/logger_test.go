/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 21:56:00
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-24 23:14:16
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gflogger_test

import (
	"context"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"testing"
)

func TestLogger(t *testing.T) {
	ctx := context.Background()
	gflogger.Print(ctx, "hello")
	gflogger.Printf(ctx, "hello %s", "world")
}
