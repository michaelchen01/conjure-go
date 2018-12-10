// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was generated by Conjure and should not be manually edited.

package api

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/palantir/conjure-go-runtime/conjure-go-client/httpclient"
)

type TestServiceClient interface {
	Echo(ctx context.Context) error
	PathParam(ctx context.Context, paramArg string) error
	Bytes(ctx context.Context) (CustomObject, error)
	Binary(ctx context.Context) (io.ReadCloser, error)
}

type testServiceClient struct {
	client httpclient.Client
}

func NewTestServiceClient(client httpclient.Client) TestServiceClient {
	return &testServiceClient{client: client}
}

func (c *testServiceClient) Echo(ctx context.Context) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Echo"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/echo"))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *testServiceClient) PathParam(ctx context.Context, paramArg string) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("PathParam"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/path/%s", url.PathEscape(fmt.Sprint(paramArg))))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *testServiceClient) Bytes(ctx context.Context) (CustomObject, error) {
	var defaultReturnVal CustomObject
	var returnVal *CustomObject
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Bytes"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/bytes"))
	requestParams = append(requestParams, httpclient.WithJSONResponse(&returnVal))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return defaultReturnVal, err
	}
	_ = resp
	if returnVal == nil {
		return defaultReturnVal, fmt.Errorf("returnVal cannot be nil")
	}
	return *returnVal, nil
}

func (c *testServiceClient) Binary(ctx context.Context) (io.ReadCloser, error) {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Binary"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/binary"))
	requestParams = append(requestParams, httpclient.WithRawResponseBody())
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}