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
	"sync/atomic"
	"time"
)

type FixedWindowLimiter struct {
	limit      int32
	interval  int64
	count     int32
	timestamp int64
	l sync.Mutex
}

func (f *FixedWindowLimiter) Take() bool {
	current := time.Now().UnixNano()
	f.l.Lock()
	if f.timestamp+f.interval < current {
		// 开启新的窗口
		f.timestamp = current
		f.count = 0
	}
	f.l.Unlock()
	return atomic.AddInt32(&f.count, 1) <= f.limit
}

