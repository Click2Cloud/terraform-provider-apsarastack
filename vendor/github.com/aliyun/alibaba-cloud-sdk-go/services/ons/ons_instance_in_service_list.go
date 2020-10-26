package ons

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

// OnsInstanceInServiceList invokes the ons.OnsInstanceInServiceList API synchronously
func (client *Client) OnsInstanceInServiceList(request *OnsInstanceInServiceListRequest) (response *OnsInstanceInServiceListResponse, err error) {
	response = CreateOnsInstanceInServiceListResponse()
	err = client.DoAction(request, response)
	return
}

// OnsInstanceInServiceListWithChan invokes the ons.OnsInstanceInServiceList API asynchronously
func (client *Client) OnsInstanceInServiceListWithChan(request *OnsInstanceInServiceListRequest) (<-chan *OnsInstanceInServiceListResponse, <-chan error) {
	responseChan := make(chan *OnsInstanceInServiceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsInstanceInServiceList(request)
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

// OnsInstanceInServiceListWithCallback invokes the ons.OnsInstanceInServiceList API asynchronously
func (client *Client) OnsInstanceInServiceListWithCallback(request *OnsInstanceInServiceListRequest, callback func(response *OnsInstanceInServiceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsInstanceInServiceListResponse
		var err error
		defer close(result)
		response, err = client.OnsInstanceInServiceList(request)
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

// OnsInstanceInServiceListRequest is the request struct for api OnsInstanceInServiceList
type OnsInstanceInServiceListRequest struct {
	*requests.RpcRequest
	Tag *[]OnsInstanceInServiceListTag `position:"Query" name:"Tag"  type:"Repeated"`
}

// OnsInstanceInServiceListTag is a repeated param struct in OnsInstanceInServiceListRequest
type OnsInstanceInServiceListTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// OnsInstanceInServiceListResponse is the response struct for api OnsInstanceInServiceList
type OnsInstanceInServiceListResponse struct {
	*responses.BaseResponse
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	HelpUrl   string                         `json:"HelpUrl" xml:"HelpUrl"`
	Data      DataInOnsInstanceInServiceList `json:"Data" xml:"Data"`
}

// CreateOnsInstanceInServiceListRequest creates a request to invoke OnsInstanceInServiceList API
func CreateOnsInstanceInServiceListRequest() (request *OnsInstanceInServiceListRequest) {
	request = &OnsInstanceInServiceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsInstanceInServiceList", "", "")
	request.Method = requests.POST
	return
}

// CreateOnsInstanceInServiceListResponse creates a response to parse from OnsInstanceInServiceList response
func CreateOnsInstanceInServiceListResponse() (response *OnsInstanceInServiceListResponse) {
	response = &OnsInstanceInServiceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}