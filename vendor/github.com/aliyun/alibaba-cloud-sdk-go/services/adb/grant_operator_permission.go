package adb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GrantOperatorPermission invokes the adb.GrantOperatorPermission API synchronously
// api document: https://help.aliyun.com/api/adb/grantoperatorpermission.html
func (client *Client) GrantOperatorPermission(request *GrantOperatorPermissionRequest) (response *GrantOperatorPermissionResponse, err error) {
	response = CreateGrantOperatorPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// GrantOperatorPermissionWithChan invokes the adb.GrantOperatorPermission API asynchronously
// api document: https://help.aliyun.com/api/adb/grantoperatorpermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantOperatorPermissionWithChan(request *GrantOperatorPermissionRequest) (<-chan *GrantOperatorPermissionResponse, <-chan error) {
	responseChan := make(chan *GrantOperatorPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GrantOperatorPermission(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GrantOperatorPermissionWithCallback invokes the adb.GrantOperatorPermission API asynchronously
// api document: https://help.aliyun.com/api/adb/grantoperatorpermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GrantOperatorPermissionWithCallback(request *GrantOperatorPermissionRequest, callback func(response *GrantOperatorPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GrantOperatorPermissionResponse
		var err error
		defer close(result)
		response, err = client.GrantOperatorPermission(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GrantOperatorPermissionRequest is the request struct for api GrantOperatorPermission
type GrantOperatorPermissionRequest struct {
	*requests.RpcRequest
	Privileges           string           `position:"Query" name:"Privileges"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ExpiredTime          string           `position:"Query" name:"ExpiredTime"`
}

// GrantOperatorPermissionResponse is the response struct for api GrantOperatorPermission
type GrantOperatorPermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGrantOperatorPermissionRequest creates a request to invoke GrantOperatorPermission API
func CreateGrantOperatorPermissionRequest() (request *GrantOperatorPermissionRequest) {
	request = &GrantOperatorPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "GrantOperatorPermission", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGrantOperatorPermissionResponse creates a response to parse from GrantOperatorPermission response
func CreateGrantOperatorPermissionResponse() (response *GrantOperatorPermissionResponse) {
	response = &GrantOperatorPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
