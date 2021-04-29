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

package v20201210

import (
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-10"

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

func NewCreateStructureTaskRequest() (request *CreateStructureTaskRequest) {
	request = &CreateStructureTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cii", APIVersion, "CreateStructureTask")
	return
}

func NewCreateStructureTaskResponse() (response *CreateStructureTaskResponse) {
	response = &CreateStructureTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 基于提供的客户及保单信息，启动结构化识别任务。
func (c *Client) CreateStructureTask(request *CreateStructureTaskRequest) (response *CreateStructureTaskResponse, err error) {
	if request == nil {
		request = NewCreateStructureTaskRequest()
	}
	response = NewCreateStructureTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStructCompareDataRequest() (request *DescribeStructCompareDataRequest) {
	request = &DescribeStructCompareDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cii", APIVersion, "DescribeStructCompareData")
	return
}

func NewDescribeStructCompareDataResponse() (response *DescribeStructCompareDataResponse) {
	response = &DescribeStructCompareDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 结构化对比查询接口，对比结构化复核前后数据差异，查询识别正确率，召回率。
func (c *Client) DescribeStructCompareData(request *DescribeStructCompareDataRequest) (response *DescribeStructCompareDataResponse, err error) {
	if request == nil {
		request = NewDescribeStructCompareDataRequest()
	}
	response = NewDescribeStructCompareDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStructureTaskResultRequest() (request *DescribeStructureTaskResultRequest) {
	request = &DescribeStructureTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cii", APIVersion, "DescribeStructureTaskResult")
	return
}

func NewDescribeStructureTaskResultResponse() (response *DescribeStructureTaskResultResponse) {
	response = &DescribeStructureTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 依据任务ID获取结构化结果接口。
func (c *Client) DescribeStructureTaskResult(request *DescribeStructureTaskResultRequest) (response *DescribeStructureTaskResultResponse, err error) {
	if request == nil {
		request = NewDescribeStructureTaskResultRequest()
	}
	response = NewDescribeStructureTaskResultResponse()
	err = c.Send(request, response)
	return
}
