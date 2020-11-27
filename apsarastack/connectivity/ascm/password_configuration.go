package ascm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

//////// Create
type SetPasswordPolicyRequest struct {
	*requests.RpcRequest
	HardExpiry      bool `position:"Query" name:"HardExpiry"`
	RequireLowercaseCharacters      bool `position:"Query" name:"RequireLowercaseCharacters"`
	RequireNumbers      bool `position:"Query" name:"RequireNumbers"`
	RequireSymbols      bool `position:"Query" name:"RequireSymbols"`
	RequireUppercaseCharacters      bool `position:"Query" name:"RequireUppercaseCharacters"`
	MaxLoginAttemps      int `position:"Query" name:"MaxLoginAttemps"`
	MaxPasswordAge      int `position:"Query" name:"MaxPasswordAge"`
	MinimumPasswordLength      int `position:"Query" name:"MinimumPasswordLength"`
	PasswordErrorLockPeriod      int `position:"Query" name:"PasswordErrorLockPeriod"`
	PasswordErrorTolerancePeriod      int `position:"Query" name:"PasswordErrorTolerancePeriod"`
	PasswordReusePrevention      int `position:"Query" name:"PasswordReusePrevention"`
}

type SetPasswordPolicyResponse struct {
	*responses.BaseResponse
	Code            string `json:"code" xml:"code"`
	Cost            int    `json:"string" xml:"string"`
	Message         string `json:"message" xml:"message"`
	Redirect        bool   `json:"redirect" xml:"redirect"`
	Success         string `json:"success" xml:"success"`
	AsapiRequestId  string `json:"asapiRequestId" xml:"asapiRequestId"`
	Data            PasswordPolicy `json:"data" xml:"data"`
}

func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (response *SetPasswordPolicyResponse, err error) {
	response = CreateSetPasswordPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func CreateSetPasswordPolicyRequest() (request *SetPasswordPolicyRequest) {
	request = &SetPasswordPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ascm", "2019-05-10", "SetPasswordPolicy", "ascm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAscmResponse creates a response to parse from CreateConfiguration response
func CreateSetPasswordPolicyResponse() (response *SetPasswordPolicyResponse) {
	response = &SetPasswordPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

///// Delete
type RemovePasswordPolicyByIDRequest struct {
	*requests.RpcRequest
	Id      bool `position:"Query" name:"Id"`
}

type RemovePasswordPolicyByIDResponse struct {
	*responses.BaseResponse
	Code        string `json:"code" xml:"code"`
	Cost             int `json:"string" xml:"string"`
	Message         string `json:"message" xml:"message"`
	Redirect        bool `json:"redirect" xml:"redirect"`
	Success string `json:"success" xml:"success"`
	AsapiRequestId  string `json:"asapiRequestId" xml:"asapiRequestId"`
}

func (client *Client) RemovePasswordPolicyByID(request *RemovePasswordPolicyByIDRequest) (response *RemovePasswordPolicyByIDResponse, err error) {
	response = CreateRemovePasswordPolicyByIDResponse()
	err = client.DoAction(request, response)
	return
}

func CreateRemovePasswordPolicyByIDRequest() (request *RemovePasswordPolicyByIDRequest) {
	request = &RemovePasswordPolicyByIDRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ascm", "2019-05-10", "RemoveLoginPolicyById", "ascm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAscmResponse creates a response to parse from CreateConfiguration response
func CreateRemovePasswordPolicyByIDResponse() (response *RemovePasswordPolicyByIDResponse) {
	response = &RemovePasswordPolicyByIDResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
///// Describe

type GetPasswordPolicyRequest struct {
	*requests.RpcRequest
}

type GetPasswordPolicyResponse struct {
	*responses.BaseResponse
	Code            string `json:"code" xml:"code"`
	Cost            int    `json:"string" xml:"string"`
	Message         string `json:"message" xml:"message"`
	Redirect        bool   `json:"redirect" xml:"redirect"`
	Success         string `json:"success" xml:"success"`
	AsapiRequestId  string `json:"asapiRequestId" xml:"asapiRequestId"`
	Data            PasswordPolicy `json:"data" xml:"data"`
}

func (client *Client) GetPasswordPolicy(request *GetPasswordPolicyRequest) (response *GetPasswordPolicyResponse, err error) {
	response = CreateGetPasswordPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func CreateGetPasswordPolicyRequest() (request *GetPasswordPolicyRequest) {
	request = &GetPasswordPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ascm", "2019-05-10", "GetPasswordPolicyRequest", "ascm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAscmResponse creates a response to parse from CreateConfiguration response
func CreateGetPasswordPolicyResponse() (response *GetPasswordPolicyResponse) {
	response = &GetPasswordPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}