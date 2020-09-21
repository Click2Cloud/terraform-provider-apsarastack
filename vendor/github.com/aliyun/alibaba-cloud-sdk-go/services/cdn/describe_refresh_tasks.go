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

// DescribeRefreshTasks invokes the cdn.DescribeRefreshTasks API synchronously
// api document: https://help.aliyun.com/api/cdn/describerefreshtasks.html
func (client *Client) DescribeRefreshTasks(request *DescribeRefreshTasksRequest) (response *DescribeRefreshTasksResponse, err error) {
	response = CreateDescribeRefreshTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRefreshTasksWithChan invokes the cdn.DescribeRefreshTasks API asynchronously
// api document: https://help.aliyun.com/api/cdn/describerefreshtasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRefreshTasksWithChan(request *DescribeRefreshTasksRequest) (<-chan *DescribeRefreshTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeRefreshTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRefreshTasks(request)
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

// DescribeRefreshTasksWithCallback invokes the cdn.DescribeRefreshTasks API asynchronously
// api document: https://help.aliyun.com/api/cdn/describerefreshtasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRefreshTasksWithCallback(request *DescribeRefreshTasksRequest, callback func(response *DescribeRefreshTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRefreshTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeRefreshTasks(request)
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

// DescribeRefreshTasksRequest is the request struct for api DescribeRefreshTasks
type DescribeRefreshTasksRequest struct {
	*requests.RpcRequest
	ObjectPath      string           `position:"Query" name:"ObjectPath"`
	StartTime       string           `position:"Query" name:"StartTime"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ObjectType      string           `position:"Query" name:"ObjectType"`
	TaskId          string           `position:"Query" name:"TaskId"`
	DomainName      string           `position:"Query" name:"DomainName"`
	EndTime         string           `position:"Query" name:"EndTime"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	Status          string           `position:"Query" name:"Status"`
}

// DescribeRefreshTasksResponse is the response struct for api DescribeRefreshTasks
type DescribeRefreshTasksResponse struct {
	*responses.BaseResponse
	RequestId  string                      `json:"RequestId" xml:"RequestId"`
	PageNumber int64                       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64                       `json:"PageSize" xml:"PageSize"`
	TotalCount int64                       `json:"TotalCount" xml:"TotalCount"`
	Tasks      TasksInDescribeRefreshTasks `json:"Tasks" xml:"Tasks"`
}

// CreateDescribeRefreshTasksRequest creates a request to invoke DescribeRefreshTasks API
func CreateDescribeRefreshTasksRequest() (request *DescribeRefreshTasksRequest) {
	request = &DescribeRefreshTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeRefreshTasks", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRefreshTasksResponse creates a response to parse from DescribeRefreshTasks response
func CreateDescribeRefreshTasksResponse() (response *DescribeRefreshTasksResponse) {
	response = &DescribeRefreshTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
