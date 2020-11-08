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

package main

import (
	"context"
	"fmt"

	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"google.golang.org/grpc"
)

type GlobalRateLimiterBuilder struct {
	RejectAction func(interface{}, grpc.UnaryHandler) (interface{}, error)
}

func NewGlobalRateLimiterBuilder(rejectAction func(interface{}, grpc.UnaryHandler) (interface{}, error)) (*GlobalRateLimiterBuilder, error) {
	err := api.InitDefault()
	// _ = api.InitWithConfigFile("/path/to/your/config/file")
	// _ = api.InitWithConfig(&config.Entity{})
	if err != nil {
		return nil, err
	}

	// 将远程配置load下来
	_, err = flow.LoadRules([]*flow.FlowRule{
		{
			Resource:        "created-from-dashboard",
		},
	})
	if err != nil {
		return nil, err
	}
	return &GlobalRateLimiterBuilder{RejectAction: rejectAction}, nil
}

func (gb *GlobalRateLimiterBuilder) GlobalRateLimit(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	en, err := api.Entry("created-from-dashboard")
	fmt.Println(en)
	if err != nil {
		// 请求被拒绝，返回error
		return gb.RejectAction(req, handler)
	} else {
		// 请求允许通过，进行下一步
		return handler(ctx, req)
	}
}
