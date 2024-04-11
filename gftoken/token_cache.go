/*
 * @Author: liusuxian 382185882@qq.com
 * @Date: 2024-01-20 15:38:07
 * @LastEditors: liusuxian 382185882@qq.com
 * @LastEditTime: 2024-04-11 17:55:07
 * @Description:
 *
 * Copyright (c) 2024 by liusuxian email: 382185882@qq.com, All Rights Reserved.
 */
package gftoken

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/liusuxian/gf-toolkit/gflogger"
	"github.com/liusuxian/go-toolkit/gtkresp"
	"time"
)

// setCache 设置缓存
func (m *Token) setCache(ctx context.Context, cacheKey string, userCache g.Map) gtkresp.Response {
	switch m.CacheMode {
	case CacheModeCache, CacheModeFile:
		gcache.Set(ctx, cacheKey, userCache, gconv.Duration(m.Timeout)*time.Millisecond)
		if m.CacheMode == CacheModeFile {
			m.writeFileCache(ctx)
		}
	case CacheModeRedis:
		var (
			cacheValueJson []byte
			err            error
		)
		if cacheValueJson, err = gjson.Encode(userCache); err != nil {
			gflogger.Error(ctx, "[gftoken]cache json encode error", err)
			return gtkresp.Fail(ERROR, "cache json encode error")
		}
		if err := m.RedisCache.Set(ctx, cacheKey, cacheValueJson, time.Duration(m.Timeout/1000)*time.Second); err != nil {
			gflogger.Error(ctx, "[gftoken]cache set error", err)
			return gtkresp.Fail(ERROR, "cache set error")
		}
	default:
		return gtkresp.Fail(ERROR, "cache model error")
	}

	return gtkresp.Succ(userCache)
}

// getCache 获取缓存
func (m *Token) getCache(ctx context.Context, cacheKey string) gtkresp.Response {
	var userCache g.Map
	switch m.CacheMode {
	case CacheModeCache, CacheModeFile:
		var (
			userCacheValue *gvar.Var
			err            error
		)
		if userCacheValue, err = gcache.Get(ctx, cacheKey); err != nil {
			gflogger.Error(ctx, "[gftoken]cache get error", err)
			return gtkresp.Fail(ERROR, "cache get error")
		}
		if userCacheValue.IsNil() {
			return gtkresp.Unauthorized("login timeout or not login", "")
		}
		userCache = gconv.Map(userCacheValue)
	case CacheModeRedis:
		var (
			userCacheJson any
			err           error
		)
		if userCacheJson, err = m.RedisCache.Get(ctx, cacheKey); err != nil {
			gflogger.Error(ctx, "[gftoken]cache get error", err)
			return gtkresp.Fail(ERROR, "cache get error")
		}
		if userCacheJson == nil {
			return gtkresp.Unauthorized("login timeout or not login", "")
		}
		if err = gjson.DecodeTo(userCacheJson, &userCache); err != nil {
			gflogger.Error(ctx, "[gftoken]cache get json error", err)
			return gtkresp.Fail(ERROR, "cache get json error")
		}
	default:
		return gtkresp.Fail(ERROR, "cache model error")
	}

	return gtkresp.Succ(userCache)
}

// removeCache 删除缓存
func (m *Token) removeCache(ctx context.Context, cacheKey string) gtkresp.Response {
	switch m.CacheMode {
	case CacheModeCache, CacheModeFile:
		if _, err := gcache.Remove(ctx, cacheKey); err != nil {
			gflogger.Error(ctx, err)
		}
		if m.CacheMode == CacheModeFile {
			m.writeFileCache(ctx)
		}
	case CacheModeRedis:
		if err := m.RedisCache.Delete(ctx, cacheKey); err != nil {
			gflogger.Error(ctx, "[gftoken]cache remove error", err)
			return gtkresp.Fail(ERROR, "cache remove error")
		}
	default:
		return gtkresp.Fail(ERROR, "cache model error")
	}

	return gtkresp.Succ("")
}

func (m *Token) writeFileCache(ctx context.Context) {
	file := gfile.Temp(CacheModeFileDat)
	data, e := gcache.Data(ctx)
	if e != nil {
		gflogger.Error(ctx, "[gftoken]cache writeFileCache error", e)
	}
	gfile.PutContents(file, gjson.New(data).MustToJsonString())
}

func (m *Token) initFileCache(ctx context.Context) {
	file := gfile.Temp(CacheModeFileDat)
	if !gfile.Exists(file) {
		return
	}
	data := gfile.GetContents(file)
	maps := gconv.Map(data)
	if maps == nil || len(maps) <= 0 {
		return
	}
	for k, v := range maps {
		gcache.Set(ctx, k, v, gconv.Duration(m.Timeout)*time.Millisecond)
	}
}
