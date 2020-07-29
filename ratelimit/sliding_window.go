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
	"container/list"
	"sync"
	"time"
)

type SlidingWindowLimiter struct {
	rate     int
	interval int64
	mutex    *sync.Mutex
	queue    *list.List
}

func (s *SlidingWindowLimiter) Take() bool {

	s.mutex.Lock()
	defer s.mutex.Unlock()
	// 快速路径
	size := s.queue.Len()
	current := time.Now().UnixNano()
	if size < s.rate {
		s.queue.PushBack(current)
		return true
	}

	// 慢路径
	boundary := current - s.interval

	timestamp := s.queue.Front()

	for timestamp != nil && timestamp.Value.(int64) < boundary {
		s.queue.Remove(timestamp)
		timestamp = s.queue.Front()
	}
	if s.queue.Len() < s.rate {
		s.queue.PushBack(current)
		return true
	}
	return false
}
