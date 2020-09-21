package hbase

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

// PurgeInstance invokes the hbase.PurgeInstance API synchronously
// api document: https://help.aliyun.com/api/hbase/purgeinstance.html
func (client *Client) PurgeInstance(request *PurgeInstanceRequest) (response *PurgeInstanceResponse, err error) {
	response = CreatePurgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// PurgeInstanceWithChan invokes the hbase.PurgeInstance API asynchronously
// api document: https://help.aliyun.com/api/hbase/purgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PurgeInstanceWithChan(request *PurgeInstanceRequest) (<-chan *PurgeInstanceResponse, <-chan error) {
	responseChan := make(chan *PurgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PurgeInstance(request)
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

// PurgeInstanceWithCallback invokes the hbase.PurgeInstance API asynchronously
// api document: https://help.aliyun.com/api/hbase/purgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PurgeInstanceWithCallback(request *PurgeInstanceRequest, callback func(response *PurgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PurgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.PurgeInstance(request)
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

// PurgeInstanceRequest is the request struct for api PurgeInstance
type PurgeInstanceRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// PurgeInstanceResponse is the response struct for api PurgeInstance
type PurgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePurgeInstanceRequest creates a request to invoke PurgeInstance API
func CreatePurgeInstanceRequest() (request *PurgeInstanceRequest) {
	request = &PurgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "PurgeInstance", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePurgeInstanceResponse creates a response to parse from PurgeInstance response
func CreatePurgeInstanceResponse() (response *PurgeInstanceResponse) {
	response = &PurgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
