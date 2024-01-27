/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-19 22:29:06
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-01-27 15:06:58
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gfweixin_test

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/liusuxian/gf-toolkit/gfredis"
	"github.com/liusuxian/gf-toolkit/gfweixin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthCode2Session(t *testing.T) {
	r := miniredis.RunT(t)
	client := gfredis.NewClient(func(cc *gfredis.ClientConfig) {
		cc.Addr = r.Addr()
		cc.Password = ""
		cc.DB = 1
	})
	defer client.Close()

	weChatService := gfweixin.NewWeChatService("wx65064684d6c0f73f", "e20bf5f51062ab55ed2b1cec8e540502", client)
	var (
		ctx    = context.Background()
		assert = assert.New(t)
		resMap map[string]any
		err    error
	)
	resMap, err = weChatService.AuthCode2Session(ctx, "0d3AYx000aRfrR17go400dpCcR2AYx0c")
	if assert.Error(err) {
		assert.Equal(40029, gconv.Int(resMap["errcode"]))
	}
}

func TestGetStableAccessToken(t *testing.T) {
	r := miniredis.RunT(t)
	client := gfredis.NewClient(func(cc *gfredis.ClientConfig) {
		cc.Addr = r.Addr()
		cc.Password = ""
		cc.DB = 1
	})
	defer client.Close()

	weChatService := gfweixin.NewWeChatService("wx65064684d6c0f73f", "e20bf5f51062ab55ed2b1cec8e540502", client)
	var (
		ctx         = context.Background()
		assert      = assert.New(t)
		accessToken string
		err         error
	)
	accessToken, err = weChatService.GetStableAccessToken(ctx)
	if assert.NoError(err) {
		assert.NotEmpty(accessToken)
	}
	accessToken, err = weChatService.GetStableAccessToken(ctx, true)
	if assert.NoError(err) {
		assert.NotEmpty(accessToken)
	}
}

func TestGetPhoneNumber(t *testing.T) {
	r := miniredis.RunT(t)
	client := gfredis.NewClient(func(cc *gfredis.ClientConfig) {
		cc.Addr = r.Addr()
		cc.Password = ""
		cc.DB = 1
	})
	defer client.Close()

	weChatService := gfweixin.NewWeChatService("wx65064684d6c0f73f", "e20bf5f51062ab55ed2b1cec8e540502", client)
	var (
		ctx         = context.Background()
		assert      = assert.New(t)
		accessToken string
		err         error
	)
	accessToken, err = weChatService.GetStableAccessToken(ctx)
	if assert.NoError(err) {
		assert.NotEmpty(accessToken)
	}
	var (
		phoneNumber        string
		invalidAccessToken bool
	)
	phoneNumber, invalidAccessToken, err = weChatService.GetPhoneNumber(ctx, "", accessToken)
	if assert.Error(err) {
		assert.False(invalidAccessToken)
		assert.Empty(phoneNumber)
	}
}
