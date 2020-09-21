package cdn

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

// ListRealtimeLogDeliveryDomains invokes the cdn.ListRealtimeLogDeliveryDomains API synchronously
// api document: https://help.aliyun.com/api/cdn/listrealtimelogdeliverydomains.html
func (client *Client) ListRealtimeLogDeliveryDomains(request *ListRealtimeLogDeliveryDomainsRequest) (response *ListRealtimeLogDeliveryDomainsResponse, err error) {
	response = CreateListRealtimeLogDeliveryDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRealtimeLogDeliveryDomainsWithChan invokes the cdn.ListRealtimeLogDeliveryDomains API asynchronously
// api document: https://help.aliyun.com/api/cdn/listrealtimelogdeliverydomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRealtimeLogDeliveryDomainsWithChan(request *ListRealtimeLogDeliveryDomainsRequest) (<-chan *ListRealtimeLogDeliveryDomainsResponse, <-chan error) {
	responseChan := make(chan *ListRealtimeLogDeliveryDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRealtimeLogDeliveryDomains(request)
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

// ListRealtimeLogDeliveryDomainsWithCallback invokes the cdn.ListRealtimeLogDeliveryDomains API asynchronously
// api document: https://help.aliyun.com/api/cdn/listrealtimelogdeliverydomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRealtimeLogDeliveryDomainsWithCallback(request *ListRealtimeLogDeliveryDomainsRequest, callback func(response *ListRealtimeLogDeliveryDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRealtimeLogDeliveryDomainsResponse
		var err error
		defer close(result)
		response, err = client.ListRealtimeLogDeliveryDomains(request)
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

// ListRealtimeLogDeliveryDomainsRequest is the request struct for api ListRealtimeLogDeliveryDomains
type ListRealtimeLogDeliveryDomainsRequest struct {
	*requests.RpcRequest
	Project  string           `position:"Query" name:"Project"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	Region   string           `position:"Query" name:"Region"`
	Logstore string           `position:"Query" name:"Logstore"`
}

// ListRealtimeLogDeliveryDomainsResponse is the response struct for api ListRealtimeLogDeliveryDomains
type ListRealtimeLogDeliveryDomainsResponse struct {
	*responses.BaseResponse
	RequestId string                                  `json:"RequestId" xml:"RequestId"`
	Content   ContentInListRealtimeLogDeliveryDomains `json:"Content" xml:"Content"`
}

// CreateListRealtimeLogDeliveryDomainsRequest creates a request to invoke ListRealtimeLogDeliveryDomains API
func CreateListRealtimeLogDeliveryDomainsRequest() (request *ListRealtimeLogDeliveryDomainsRequest) {
	request = &ListRealtimeLogDeliveryDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "ListRealtimeLogDeliveryDomains", "", "")
	request.Method = requests.GET
	return
}

// CreateListRealtimeLogDeliveryDomainsResponse creates a response to parse from ListRealtimeLogDeliveryDomains response
func CreateListRealtimeLogDeliveryDomainsResponse() (response *ListRealtimeLogDeliveryDomainsResponse) {
	response = &ListRealtimeLogDeliveryDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
