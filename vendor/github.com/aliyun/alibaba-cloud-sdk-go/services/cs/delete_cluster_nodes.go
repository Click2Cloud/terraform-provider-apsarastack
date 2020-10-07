package cs

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

// DeleteClusterNodes invokes the cs.DeleteClusterNodes API synchronously
// api document: https://help.aliyun.com/api/cs/deleteclusternodes.html
func (client *Client) DeleteClusterNodes(request *DeleteClusterNodesRequest) (response *DeleteClusterNodesResponse, err error) {
	response = CreateDeleteClusterNodesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteClusterNodesWithChan invokes the cs.DeleteClusterNodes API asynchronously
// api document: https://help.aliyun.com/api/cs/deleteclusternodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteClusterNodesWithChan(request *DeleteClusterNodesRequest) (<-chan *DeleteClusterNodesResponse, <-chan error) {
	responseChan := make(chan *DeleteClusterNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteClusterNodes(request)
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

// DeleteClusterNodesWithCallback invokes the cs.DeleteClusterNodes API asynchronously
// api document: https://help.aliyun.com/api/cs/deleteclusternodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteClusterNodesWithCallback(request *DeleteClusterNodesRequest, callback func(response *DeleteClusterNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteClusterNodesResponse
		var err error
		defer close(result)
		response, err = client.DeleteClusterNodes(request)
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

// DeleteClusterNodesRequest is the request struct for api DeleteClusterNodes
type DeleteClusterNodesRequest struct {
	*requests.RoaRequest
	ReleaseNode string `position:"Body" name:"release_node"`
	ClusterId   string `position:"Path" name:"ClusterId"`
}

// DeleteClusterNodesResponse is the response struct for api DeleteClusterNodes
type DeleteClusterNodesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteClusterNodesRequest creates a request to invoke DeleteClusterNodes API
func CreateDeleteClusterNodesRequest() (request *DeleteClusterNodesRequest) {
	request = &DeleteClusterNodesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DeleteClusterNodes", "/clusters/[ClusterId]/nodes", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteClusterNodesResponse creates a response to parse from DeleteClusterNodes response
func CreateDeleteClusterNodesResponse() (response *DeleteClusterNodesResponse) {
	response = &DeleteClusterNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
