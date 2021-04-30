package cns

import (
	"encoding/json"

	tchttp "github.com/iftechio/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DomainListRequest struct {
	*tchttp.BaseRequest

	Offset     *int    `json:"offset,omitempty" name:"offset"`
	Length     *int    `json:"length,omitempty" name:"length"`
	Keyword    *string `json:"keyword,omitempty" name:"keyword"`
	QProjectId *int    `json:"qProjectId,omitempty" name:"qProjectId"`
}

func (r *DomainListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DomainListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainListResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
	Data     *struct {
		Info *struct {
			DomainTotal *int `json:"domain_total"`
		}
		Domains *[]struct {
			ID               *int        `json:"id" name:"ID"`
			Status           *string     `json:"status"`
			GroupID          interface{} `json:"group_id"`
			SearchEnginePush *string     `json:"searchengine_push"`
			IsMark           *string     `json:"is_mark"`
			TTL              *string     `json:"ttl"`
			CnameSpeedup     *string     `json:"cname_speedup"`
			Remark           *string     `json:"remark"`
			CreatedOn        *string     `json:"created_on"`
			UpdatedOn        *string     `json:"updated_on"`
			QProjectID       *int        `json:"q_project_id"`
			Punycode         *string     `json:"punycode"`
			ExtStatus        *string     `json:"ext_status"`
			SrcFlag          *string     `json:"src_flag"`
			Name             *string     `json:"name"`
			Grade            *string     `json:"grade"`
			GradeTitle       *string     `json:"grade_title"`
			IsVip            *string     `json:"is_vip"`
			Owner            *string     `json:"owner"`
			Records          *string     `json:"records"`
			MinTTL           *int        `json:"min_ttl"`
			VipStartAt       *string     `json:"vip_start_at,omitempty"`
			VipEndAt         *string     `json:"vip_end_at,omitempty"`
			VipAutoRenew     *string     `json:"vip_auto_renew,omitempty"`
		} `json:"domains"`
	} `json:"data"`
}

func (r *DomainListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DomainListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCreateRequest struct {
	*tchttp.BaseRequest

	Domain    *string `json:"domain" name:"domain"`
	ProjectId *int    `json:"projectId,omitempty" name:"projectId"`
}

func (r *DomainCreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DomainCreateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainCreateResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
	Data     *struct {
		Domain *struct {
			ID       *int    `json:"id"`
			Punycode *string `json:"punycode"`
			Domain   *string `json:"domain"`
		} `json:"domain"`
	} `json:"data"`
}

func (r *DomainCreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DomainCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDomainStatusRequest struct {
	*tchttp.BaseRequest

	Domain    *string `json:"domain" name:"domain"`
	ProjectId *int    `json:"projectId,omitempty" name:"projectId"`
}

func (r *SetDomainStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDomainStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDomainStatusResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
}

func (r *SetDomainStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDomainStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainDeleteRequest struct {
	*tchttp.BaseRequest

	Domain *string `json:"domain" name:"domain"`
}

func (r *DomainDeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DomainDeleteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DomainDeleteResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
	Data     *struct {
		Record *struct {
			ID     *string `json:"id"`
			Name   *string `json:"name"`
			Status *string `json:"status"`
			Weight *int    `json:"weight"`
		} `json:"record"`
	} `json:"data"`
}

func (r *DomainDeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DomainDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordCreateRequest struct {
	*tchttp.BaseRequest

	Domain     *string `json:"domain" name:"domain"`
	SubDomain  *string `json:"subDomain" name:"subDomain"`
	RecordType *string `json:"recordType" name:"recordType"`
	RecordLine *string `json:"recordLine" name:"recordLine"`
	Value      *string `json:"value" name:"value"`
	TTL        *int    `json:"ttl,omitempty" name:"ttl" `
	Mx         *int    `json:"mx,omitempty" name:"mx"`
}

func (r *RecordCreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordCreateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordCreateResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
	Data     *struct {
		Record *struct {
			ID     *string `json:"id"`
			Name   *string `json:"name"`
			Status *string `json:"status"`
			Weight *int    `json:"weight"`
		} `json:"record"`
	} `json:"data"`
}

func (r *RecordCreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordStatusRequest struct {
	*tchttp.BaseRequest

	Domain   *string `json:"domain" name:"domain"`
	RecordID *int    `json:"recordId" name:"recordId"`
	Status   *string `json:"status,omitempty" name:"status"`
}

func (r *RecordStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordStatusResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
	Data     *struct {
		Record *struct {
			ID     *string `json:"id"`
			Name   *string `json:"name"`
			Status *string `json:"status"`
			Weight *int    `json:"weight"`
		} `json:"record"`
	} `json:"data"`
}

func (r *RecordStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordModifyRequest struct {
	*tchttp.BaseRequest

	Domain     *string `json:"domain" name:"domain"`
	RecordID   *int    `json:"recordID" name:"recordId"`
	SubDomain  *string `json:"subDomain" name:"subDomain"`
	RecordType *string `json:"recordType" name:"recordType"`
	RecordLine *string `json:"recordLine" name:"recordLine"`
	Value      *string `json:"value" name:"value"`
	TTL        *int    `json:"ttl,omitempty" name:"ttl"`
	Mx         *int    `json:"mx,omitempty" name:"mx"`
}

func (r *RecordModifyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordModifyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordModifyResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code,omitempty" name:"code"`
	Message  *string `json:"message,omitempty" name:"message"`
	CodeDesc *string `json:"codeDesc,omitempty" name:"codeDesc"`
	Data     *struct {
		Record *struct {
			ID     *int    `json:"id"`
			Name   *string `json:"name"`
			Status *string `json:"status"`
			Value  *string `json:"value"`
			Weight *int    `json:"weight"`
		} `json:"record"`
	} `json:"data"`
}

func (r *RecordModifyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordModifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordListRequest struct {
	*tchttp.BaseRequest

	Domain     *string `json:"domain" name:"domain"`
	Offset     *int    `json:"offset,omitempty" name:"offset"`
	Length     *int    `json:"length,omitempty" name:"length"`
	SubDomain  *string `json:"subDomain,omitempty" name:"subDomain"`
	RecordType *string `json:"recordType,omitempty" name:"recordType"`
	QProjectID *int    `json:"qProjectId,omitempty" name:"qProjectId"`
}

func (r *RecordListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Domain struct {
	ID         *string   `json:"id" `
	Name       *string   `json:"name"`
	Punycode   *string   `json:"punycode"`
	Grade      *string   `json:"grade"`
	Owner      *string   `json:"owner"`
	ExtStatus  *string   `json:"ext_status"`
	TTL        *int      `json:"ttl"`
	MinTTL     *int      `json:"min_ttl"`
	DnspodNs   *[]string `json:"dnspod_ns"`
	Status     *string   `json:"status"`
	QProjectID *int      `json:"q_project_id"`
}

type Info struct {
	SubDomains  *string `json:"sub_domains"`
	RecordTotal *string `json:"record_total"`
}
type Record struct {
	ID         *int        `json:"id"`
	TTL        *int        `json:"ttl"`
	Value      *string     `json:"value"`
	Enabled    *int        `json:"enabled"`
	Status     *string     `json:"status"`
	UpdatedOn  *string     `json:"updated_on"`
	QProjectID interface{} `json:"q_project_id"`
	Name       *string     `json:"name"`
	Line       *string     `json:"line"`
	LineID     *string     `json:"line_id"`
	Type       *string     `json:"type"`
	Remark     *string     `json:"remark"`
	Mx         *int        `json:"mx"`
	Hold       *string     `json:"hold,omitempty"`
}

type Data struct {
	Domain  *Domain   `json:"domain"`
	Info    *Info     `json:"info"`
	Records []*Record `json:"records"`
}

type RecordListResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
	Data     *Data   `json:"data"`
}

func (r *RecordListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordDeleteRequest struct {
	*tchttp.BaseRequest

	Domain   *string `json:"domain" name:"domain"`
	RecordID *int    `json:"recordId,omitempty" name:"recordId"`
}

func (r *RecordDeleteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordDeleteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordDeleteResponse struct {
	*tchttp.BaseResponse

	Code     *int    `json:"code"`
	Message  *string `json:"message"`
	CodeDesc *string `json:"codeDesc"`
}

func (r *RecordDeleteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RecordDeleteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
