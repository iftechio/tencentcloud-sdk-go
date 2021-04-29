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

package v20190422

import (
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-04-22"

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


func NewCreateJobRequest() (request *CreateJobRequest) {
    request = &CreateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateJob")
    return
}

func NewCreateJobResponse() (response *CreateJobResponse) {
    response = &CreateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 新建作业接口，一个 AppId 最多允许创建1000个作业
func (c *Client) CreateJob(request *CreateJobRequest) (response *CreateJobResponse, err error) {
    if request == nil {
        request = NewCreateJobRequest()
    }
    response = NewCreateJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateJobConfigRequest() (request *CreateJobConfigRequest) {
    request = &CreateJobConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateJobConfig")
    return
}

func NewCreateJobConfigResponse() (response *CreateJobConfigResponse) {
    response = &CreateJobConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建作业配置，一个作业最多有100个配置版本
func (c *Client) CreateJobConfig(request *CreateJobConfigRequest) (response *CreateJobConfigResponse, err error) {
    if request == nil {
        request = NewCreateJobConfigRequest()
    }
    response = NewCreateJobConfigResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
    request = &CreateResourceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateResource")
    return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
    response = &CreateResourceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建资源接口
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
    if request == nil {
        request = NewCreateResourceRequest()
    }
    response = NewCreateResourceResponse()
    err = c.Send(request, response)
    return
}

func NewCreateResourceConfigRequest() (request *CreateResourceConfigRequest) {
    request = &CreateResourceConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "CreateResourceConfig")
    return
}

func NewCreateResourceConfigResponse() (response *CreateResourceConfigResponse) {
    response = &CreateResourceConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建资源配置接口
func (c *Client) CreateResourceConfig(request *CreateResourceConfigRequest) (response *CreateResourceConfigResponse, err error) {
    if request == nil {
        request = NewCreateResourceConfigRequest()
    }
    response = NewCreateResourceConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourceConfigsRequest() (request *DeleteResourceConfigsRequest) {
    request = &DeleteResourceConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteResourceConfigs")
    return
}

func NewDeleteResourceConfigsResponse() (response *DeleteResourceConfigsResponse) {
    response = &DeleteResourceConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除资源版本
func (c *Client) DeleteResourceConfigs(request *DeleteResourceConfigsRequest) (response *DeleteResourceConfigsResponse, err error) {
    if request == nil {
        request = NewDeleteResourceConfigsRequest()
    }
    response = NewDeleteResourceConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteResourcesRequest() (request *DeleteResourcesRequest) {
    request = &DeleteResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteResources")
    return
}

func NewDeleteResourcesResponse() (response *DeleteResourcesResponse) {
    response = &DeleteResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除资源接口
func (c *Client) DeleteResources(request *DeleteResourcesRequest) (response *DeleteResourcesResponse, err error) {
    if request == nil {
        request = NewDeleteResourcesRequest()
    }
    response = NewDeleteResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteTableConfigRequest() (request *DeleteTableConfigRequest) {
    request = &DeleteTableConfigRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DeleteTableConfig")
    return
}

func NewDeleteTableConfigResponse() (response *DeleteTableConfigResponse) {
    response = &DeleteTableConfigResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除作业表配置
func (c *Client) DeleteTableConfig(request *DeleteTableConfigRequest) (response *DeleteTableConfigResponse, err error) {
    if request == nil {
        request = NewDeleteTableConfigRequest()
    }
    response = NewDeleteTableConfigResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobConfigsRequest() (request *DescribeJobConfigsRequest) {
    request = &DescribeJobConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeJobConfigs")
    return
}

func NewDescribeJobConfigsResponse() (response *DescribeJobConfigsResponse) {
    response = &DescribeJobConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询作业配置列表，一次最多查询100个
func (c *Client) DescribeJobConfigs(request *DescribeJobConfigsRequest) (response *DescribeJobConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeJobConfigsRequest()
    }
    response = NewDescribeJobConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeJobsRequest() (request *DescribeJobsRequest) {
    request = &DescribeJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeJobs")
    return
}

func NewDescribeJobsResponse() (response *DescribeJobsResponse) {
    response = &DescribeJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询作业
func (c *Client) DescribeJobs(request *DescribeJobsRequest) (response *DescribeJobsResponse, err error) {
    if request == nil {
        request = NewDescribeJobsRequest()
    }
    response = NewDescribeJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceConfigsRequest() (request *DescribeResourceConfigsRequest) {
    request = &DescribeResourceConfigsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeResourceConfigs")
    return
}

func NewDescribeResourceConfigsResponse() (response *DescribeResourceConfigsResponse) {
    response = &DescribeResourceConfigsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述资源配置接口
func (c *Client) DescribeResourceConfigs(request *DescribeResourceConfigsRequest) (response *DescribeResourceConfigsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceConfigsRequest()
    }
    response = NewDescribeResourceConfigsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourceRelatedJobsRequest() (request *DescribeResourceRelatedJobsRequest) {
    request = &DescribeResourceRelatedJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeResourceRelatedJobs")
    return
}

func NewDescribeResourceRelatedJobsResponse() (response *DescribeResourceRelatedJobsResponse) {
    response = &DescribeResourceRelatedJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取资源关联作业信息
func (c *Client) DescribeResourceRelatedJobs(request *DescribeResourceRelatedJobsRequest) (response *DescribeResourceRelatedJobsResponse, err error) {
    if request == nil {
        request = NewDescribeResourceRelatedJobsRequest()
    }
    response = NewDescribeResourceRelatedJobsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
    request = &DescribeResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeResources")
    return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
    response = &DescribeResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述资源接口
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeResourcesRequest()
    }
    response = NewDescribeResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSystemResourcesRequest() (request *DescribeSystemResourcesRequest) {
    request = &DescribeSystemResourcesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "DescribeSystemResources")
    return
}

func NewDescribeSystemResourcesResponse() (response *DescribeSystemResourcesResponse) {
    response = &DescribeSystemResourcesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 描述系统资源接口
func (c *Client) DescribeSystemResources(request *DescribeSystemResourcesRequest) (response *DescribeSystemResourcesResponse, err error) {
    if request == nil {
        request = NewDescribeSystemResourcesRequest()
    }
    response = NewDescribeSystemResourcesResponse()
    err = c.Send(request, response)
    return
}

func NewRunJobsRequest() (request *RunJobsRequest) {
    request = &RunJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "RunJobs")
    return
}

func NewRunJobsResponse() (response *RunJobsResponse) {
    response = &RunJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量启动或者恢复作业，批量操作数量上限20
func (c *Client) RunJobs(request *RunJobsRequest) (response *RunJobsResponse, err error) {
    if request == nil {
        request = NewRunJobsRequest()
    }
    response = NewRunJobsResponse()
    err = c.Send(request, response)
    return
}

func NewStopJobsRequest() (request *StopJobsRequest) {
    request = &StopJobsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("oceanus", APIVersion, "StopJobs")
    return
}

func NewStopJobsResponse() (response *StopJobsResponse) {
    response = &StopJobsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量停止作业，批量操作数量上限为20
func (c *Client) StopJobs(request *StopJobsRequest) (response *StopJobsResponse, err error) {
    if request == nil {
        request = NewStopJobsRequest()
    }
    response = NewStopJobsResponse()
    err = c.Send(request, response)
    return
}
