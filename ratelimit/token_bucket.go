// Copyright 2020 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ratelimit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	// 生产令牌的时间间隔
	interval time.Duration
	// 桶的容量
	capacity int

	// 上次生产令牌的时间
	lastTime time.Time
	// 上次生产了到如今剩下的令牌数量
	availableTokens int

	l sync.Mutex
}

func (t *TokenBucket) Allow(request Request, ctx ApplicationContext, cfg Config) bool {
	return t.Take()
}

func (t *TokenBucket) Take() bool {
	t.l.Lock()
	defer t.l.Unlock()

	// 上一次的生产的令牌还有
	if t.availableTokens > 0 {
		t.availableTokens --
		return true
	}

	// 上一次的生产的令牌全没了，要开始生产了
	now := time.Now()
	dur := now.Sub(t.lastTime)

	// 间隔时间/令牌间隔 = 令牌数量
	// 这里向下取整了
	tick := int(dur/t.interval)

	// 间隔时间太短，一个令牌都没生产到
	if tick <= 0 {
		return false

	}

	// 更新时间
	t.lastTime = now

	// 间隔太长，生产了很多令牌，但是只能放capacity个
	if tick > t.capacity {
		tick = t.capacity
	}
	t.availableTokens = tick

	// 取令牌，返回
	t.availableTokens --
	return true
}

func NewTokenBucket(interval time.Duration, capacity int) *TokenBucket {
	return &TokenBucket{
		interval:interval,
		capacity:capacity,

		lastTime:time.Now(),
		availableTokens:capacity,
	}
}

