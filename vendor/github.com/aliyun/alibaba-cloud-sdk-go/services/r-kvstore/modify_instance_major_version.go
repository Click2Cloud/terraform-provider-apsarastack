package r_kvstore

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

// ModifyInstanceMajorVersion invokes the r_kvstore.ModifyInstanceMajorVersion API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyinstancemajorversion.html
func (client *Client) ModifyInstanceMajorVersion(request *ModifyInstanceMajorVersionRequest) (response *ModifyInstanceMajorVersionResponse, err error) {
	response = CreateModifyInstanceMajorVersionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceMajorVersionWithChan invokes the r_kvstore.ModifyInstanceMajorVersion API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyinstancemajorversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceMajorVersionWithChan(request *ModifyInstanceMajorVersionRequest) (<-chan *ModifyInstanceMajorVersionResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceMajorVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceMajorVersion(request)
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

// ModifyInstanceMajorVersionWithCallback invokes the r_kvstore.ModifyInstanceMajorVersion API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyinstancemajorversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyInstanceMajorVersionWithCallback(request *ModifyInstanceMajorVersionRequest, callback func(response *ModifyInstanceMajorVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceMajorVersionResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceMajorVersion(request)
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

// ModifyInstanceMajorVersionRequest is the request struct for api ModifyInstanceMajorVersion
type ModifyInstanceMajorVersionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EffectTime           string           `position:"Query" name:"EffectTime"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	MajorVersion         string           `position:"Query" name:"MajorVersion"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ModifyInstanceMajorVersionResponse is the response struct for api ModifyInstanceMajorVersion
type ModifyInstanceMajorVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceMajorVersionRequest creates a request to invoke ModifyInstanceMajorVersion API
func CreateModifyInstanceMajorVersionRequest() (request *ModifyInstanceMajorVersionRequest) {
	request = &ModifyInstanceMajorVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceMajorVersion", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceMajorVersionResponse creates a response to parse from ModifyInstanceMajorVersion response
func CreateModifyInstanceMajorVersionResponse() (response *ModifyInstanceMajorVersionResponse) {
	response = &ModifyInstanceMajorVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
