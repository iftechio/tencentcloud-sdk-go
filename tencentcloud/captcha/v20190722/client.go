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

package v20190722

import (
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-22"

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


func NewDescribeCaptchaAppIdInfoRequest() (request *DescribeCaptchaAppIdInfoRequest) {
    request = &DescribeCaptchaAppIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaAppIdInfo")
    return
}

func NewDescribeCaptchaAppIdInfoResponse() (response *DescribeCaptchaAppIdInfoResponse) {
    response = &DescribeCaptchaAppIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询安全验证码应用APPId信息
func (c *Client) DescribeCaptchaAppIdInfo(request *DescribeCaptchaAppIdInfoRequest) (response *DescribeCaptchaAppIdInfoResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaAppIdInfoRequest()
    }
    response = NewDescribeCaptchaAppIdInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaDataRequest() (request *DescribeCaptchaDataRequest) {
    request = &DescribeCaptchaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaData")
    return
}

func NewDescribeCaptchaDataResponse() (response *DescribeCaptchaDataResponse) {
    response = &DescribeCaptchaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3  分钟级查询
func (c *Client) DescribeCaptchaData(request *DescribeCaptchaDataRequest) (response *DescribeCaptchaDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaDataRequest()
    }
    response = NewDescribeCaptchaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaDataSumRequest() (request *DescribeCaptchaDataSumRequest) {
    request = &DescribeCaptchaDataSumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaDataSum")
    return
}

func NewDescribeCaptchaDataSumResponse() (response *DescribeCaptchaDataSumResponse) {
    response = &DescribeCaptchaDataSumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码查询请求数据概况，例如：按照时间段查询数据  昨日请求量、昨日恶意比例、昨日验证量、昨日通过量、昨日恶意拦截量……
func (c *Client) DescribeCaptchaDataSum(request *DescribeCaptchaDataSumRequest) (response *DescribeCaptchaDataSumResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaDataSumRequest()
    }
    response = NewDescribeCaptchaDataSumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniDataRequest() (request *DescribeCaptchaMiniDataRequest) {
    request = &DescribeCaptchaMiniDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniData")
    return
}

func NewDescribeCaptchaMiniDataResponse() (response *DescribeCaptchaMiniDataResponse) {
    response = &DescribeCaptchaMiniDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码小程序插件分类查询数据接口，请求量type=0、通过量type=1、验证量type=2、拦截量type=3 小时级查询（五小时左右延迟）
func (c *Client) DescribeCaptchaMiniData(request *DescribeCaptchaMiniDataRequest) (response *DescribeCaptchaMiniDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniDataRequest()
    }
    response = NewDescribeCaptchaMiniDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniDataSumRequest() (request *DescribeCaptchaMiniDataSumRequest) {
    request = &DescribeCaptchaMiniDataSumRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniDataSum")
    return
}

func NewDescribeCaptchaMiniDataSumResponse() (response *DescribeCaptchaMiniDataSumResponse) {
    response = &DescribeCaptchaMiniDataSumResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码小程序插件查询请求数据概况
func (c *Client) DescribeCaptchaMiniDataSum(request *DescribeCaptchaMiniDataSumRequest) (response *DescribeCaptchaMiniDataSumResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniDataSumRequest()
    }
    response = NewDescribeCaptchaMiniDataSumResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniOperDataRequest() (request *DescribeCaptchaMiniOperDataRequest) {
    request = &DescribeCaptchaMiniOperDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniOperData")
    return
}

func NewDescribeCaptchaMiniOperDataResponse() (response *DescribeCaptchaMiniOperDataResponse) {
    response = &DescribeCaptchaMiniOperDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码小程序插件用户操作数据查询
func (c *Client) DescribeCaptchaMiniOperData(request *DescribeCaptchaMiniOperDataRequest) (response *DescribeCaptchaMiniOperDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniOperDataRequest()
    }
    response = NewDescribeCaptchaMiniOperDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniResultRequest() (request *DescribeCaptchaMiniResultRequest) {
    request = &DescribeCaptchaMiniResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniResult")
    return
}

func NewDescribeCaptchaMiniResultResponse() (response *DescribeCaptchaMiniResultResponse) {
    response = &DescribeCaptchaMiniResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 核查验证码小程序插件票据结果
func (c *Client) DescribeCaptchaMiniResult(request *DescribeCaptchaMiniResultRequest) (response *DescribeCaptchaMiniResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniResultRequest()
    }
    response = NewDescribeCaptchaMiniResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaMiniRiskResultRequest() (request *DescribeCaptchaMiniRiskResultRequest) {
    request = &DescribeCaptchaMiniRiskResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaMiniRiskResult")
    return
}

func NewDescribeCaptchaMiniRiskResultResponse() (response *DescribeCaptchaMiniRiskResultResponse) {
    response = &DescribeCaptchaMiniRiskResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 核查验证码小程序插件票据接入风控结果(Beta)
func (c *Client) DescribeCaptchaMiniRiskResult(request *DescribeCaptchaMiniRiskResultRequest) (response *DescribeCaptchaMiniRiskResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaMiniRiskResultRequest()
    }
    response = NewDescribeCaptchaMiniRiskResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaOperDataRequest() (request *DescribeCaptchaOperDataRequest) {
    request = &DescribeCaptchaOperDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaOperData")
    return
}

func NewDescribeCaptchaOperDataResponse() (response *DescribeCaptchaOperDataResponse) {
    response = &DescribeCaptchaOperDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码用户操作数据查询，验证码加载耗时type = 1 、拦截情况type = 2、 一周通过平均尝试次数 type = 3、尝试次数分布 type = 4
func (c *Client) DescribeCaptchaOperData(request *DescribeCaptchaOperDataRequest) (response *DescribeCaptchaOperDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaOperDataRequest()
    }
    response = NewDescribeCaptchaOperDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaResultRequest() (request *DescribeCaptchaResultRequest) {
    request = &DescribeCaptchaResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaResult")
    return
}

func NewDescribeCaptchaResultResponse() (response *DescribeCaptchaResultResponse) {
    response = &DescribeCaptchaResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 核查验证码票据结果
func (c *Client) DescribeCaptchaResult(request *DescribeCaptchaResultRequest) (response *DescribeCaptchaResultResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaResultRequest()
    }
    response = NewDescribeCaptchaResultResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaTicketDataRequest() (request *DescribeCaptchaTicketDataRequest) {
    request = &DescribeCaptchaTicketDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaTicketData")
    return
}

func NewDescribeCaptchaTicketDataResponse() (response *DescribeCaptchaTicketDataResponse) {
    response = &DescribeCaptchaTicketDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码用户操作票据数据查询
func (c *Client) DescribeCaptchaTicketData(request *DescribeCaptchaTicketDataRequest) (response *DescribeCaptchaTicketDataResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaTicketDataRequest()
    }
    response = NewDescribeCaptchaTicketDataResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCaptchaUserAllAppIdRequest() (request *DescribeCaptchaUserAllAppIdRequest) {
    request = &DescribeCaptchaUserAllAppIdRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "DescribeCaptchaUserAllAppId")
    return
}

func NewDescribeCaptchaUserAllAppIdResponse() (response *DescribeCaptchaUserAllAppIdResponse) {
    response = &DescribeCaptchaUserAllAppIdResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 安全验证码获取用户注册所有APPId和应用名称
func (c *Client) DescribeCaptchaUserAllAppId(request *DescribeCaptchaUserAllAppIdRequest) (response *DescribeCaptchaUserAllAppIdResponse, err error) {
    if request == nil {
        request = NewDescribeCaptchaUserAllAppIdRequest()
    }
    response = NewDescribeCaptchaUserAllAppIdResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateCaptchaAppIdInfoRequest() (request *UpdateCaptchaAppIdInfoRequest) {
    request = &UpdateCaptchaAppIdInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("captcha", APIVersion, "UpdateCaptchaAppIdInfo")
    return
}

func NewUpdateCaptchaAppIdInfoResponse() (response *UpdateCaptchaAppIdInfoResponse) {
    response = &UpdateCaptchaAppIdInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新验证码应用APPId信息
func (c *Client) UpdateCaptchaAppIdInfo(request *UpdateCaptchaAppIdInfoRequest) (response *UpdateCaptchaAppIdInfoResponse, err error) {
    if request == nil {
        request = NewUpdateCaptchaAppIdInfoRequest()
    }
    response = NewUpdateCaptchaAppIdInfoResponse()
    err = c.Send(request, response)
    return
}
