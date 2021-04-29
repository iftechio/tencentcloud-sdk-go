// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20201207

import (
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-07"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewDescribeSREInstanceAccessAddressRequest() (request *DescribeSREInstanceAccessAddressRequest) {
    request = &DescribeSREInstanceAccessAddressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstanceAccessAddress")
    return
}

func NewDescribeSREInstanceAccessAddressResponse() (response *DescribeSREInstanceAccessAddressResponse) {
    response = &DescribeSREInstanceAccessAddressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询微服务注册引擎实例访问地址
func (c *Client) DescribeSREInstanceAccessAddress(request *DescribeSREInstanceAccessAddressRequest) (response *DescribeSREInstanceAccessAddressResponse, err error) {
    if request == nil {
        request = NewDescribeSREInstanceAccessAddressRequest()
    }
    response = NewDescribeSREInstanceAccessAddressResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSREInstancesRequest() (request *DescribeSREInstancesRequest) {
    request = &DescribeSREInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tse", APIVersion, "DescribeSREInstances")
    return
}

func NewDescribeSREInstancesResponse() (response *DescribeSREInstancesResponse) {
    response = &DescribeSREInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 用于查询微服务注册中心实例列表
func (c *Client) DescribeSREInstances(request *DescribeSREInstancesRequest) (response *DescribeSREInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeSREInstancesRequest()
    }
    response = NewDescribeSREInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewManageConfigRequest() (request *ManageConfigRequest) {
    request = &ManageConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tse", APIVersion, "ManageConfig")
    return
}

func NewManageConfigResponse() (response *ManageConfigResponse) {
    response = &ManageConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理配置
func (c *Client) ManageConfig(request *ManageConfigRequest) (response *ManageConfigResponse, err error) {
    if request == nil {
        request = NewManageConfigRequest()
    }
    response = NewManageConfigResponse()
    err = c.Send(request, response)
    return
}
