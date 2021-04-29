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

package v20200224

import (
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-24"

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

func NewManageMarketingRiskRequest() (request *ManageMarketingRiskRequest) {
	request = &ManageMarketingRiskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aa", APIVersion, "ManageMarketingRisk")
	return
}

func NewManageMarketingRiskResponse() (response *ManageMarketingRiskResponse) {
	response = &ManageMarketingRiskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 活动防刷、注册保护、登录保护等营销产品的高级版本
func (c *Client) ManageMarketingRisk(request *ManageMarketingRiskRequest) (response *ManageMarketingRiskResponse, err error) {
	if request == nil {
		request = NewManageMarketingRiskRequest()
	}
	response = NewManageMarketingRiskResponse()
	err = c.Send(request, response)
	return
}

func NewQueryActivityAntiRushRequest() (request *QueryActivityAntiRushRequest) {
	request = &QueryActivityAntiRushRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aa", APIVersion, "QueryActivityAntiRush")
	return
}

func NewQueryActivityAntiRushResponse() (response *QueryActivityAntiRushResponse) {
	response = &QueryActivityAntiRushResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 腾讯云活动防刷（ActivityAntiRush，AA）是针对电商、O2O、P2P、游戏、支付等行业在促销活动中遇到“羊毛党”恶意刷取优惠福利的行为时，通过防刷引擎，精准识别出“薅羊毛”恶意行为的活动防刷服务，避免了企业被刷带来的巨大经济损失。
func (c *Client) QueryActivityAntiRush(request *QueryActivityAntiRushRequest) (response *QueryActivityAntiRushResponse, err error) {
	if request == nil {
		request = NewQueryActivityAntiRushRequest()
	}
	response = NewQueryActivityAntiRushResponse()
	err = c.Send(request, response)
	return
}

func NewQueryActivityAntiRushAdvancedRequest() (request *QueryActivityAntiRushAdvancedRequest) {
	request = &QueryActivityAntiRushAdvancedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("aa", APIVersion, "QueryActivityAntiRushAdvanced")
	return
}

func NewQueryActivityAntiRushAdvancedResponse() (response *QueryActivityAntiRushAdvancedResponse) {
	response = &QueryActivityAntiRushAdvancedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 活动防刷高级版，支持对网赚众包、网赚防刷、引流反诈骗场景的检测识别
func (c *Client) QueryActivityAntiRushAdvanced(request *QueryActivityAntiRushAdvancedRequest) (response *QueryActivityAntiRushAdvancedResponse, err error) {
	if request == nil {
		request = NewQueryActivityAntiRushAdvancedRequest()
	}
	response = NewQueryActivityAntiRushAdvancedResponse()
	err = c.Send(request, response)
	return
}
