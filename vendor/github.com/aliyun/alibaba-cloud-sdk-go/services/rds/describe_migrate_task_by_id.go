package rds

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

// DescribeMigrateTaskById invokes the rds.DescribeMigrateTaskById API synchronously
// api document: https://help.aliyun.com/api/rds/describemigratetaskbyid.html
func (client *Client) DescribeMigrateTaskById(request *DescribeMigrateTaskByIdRequest) (response *DescribeMigrateTaskByIdResponse, err error) {
	response = CreateDescribeMigrateTaskByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMigrateTaskByIdWithChan invokes the rds.DescribeMigrateTaskById API asynchronously
// api document: https://help.aliyun.com/api/rds/describemigratetaskbyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMigrateTaskByIdWithChan(request *DescribeMigrateTaskByIdRequest) (<-chan *DescribeMigrateTaskByIdResponse, <-chan error) {
	responseChan := make(chan *DescribeMigrateTaskByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMigrateTaskById(request)
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

// DescribeMigrateTaskByIdWithCallback invokes the rds.DescribeMigrateTaskById API asynchronously
// api document: https://help.aliyun.com/api/rds/describemigratetaskbyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeMigrateTaskByIdWithCallback(request *DescribeMigrateTaskByIdRequest, callback func(response *DescribeMigrateTaskByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMigrateTaskByIdResponse
		var err error
		defer close(result)
		response, err = client.DescribeMigrateTaskById(request)
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

// DescribeMigrateTaskByIdRequest is the request struct for api DescribeMigrateTaskById
type DescribeMigrateTaskByIdRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string           `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeMigrateTaskByIdResponse is the response struct for api DescribeMigrateTaskById
type DescribeMigrateTaskByIdResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	DBName         string `json:"DBName" xml:"DBName"`
	MigrateTaskId  string `json:"MigrateTaskId" xml:"MigrateTaskId"`
	CreateTime     string `json:"CreateTime" xml:"CreateTime"`
	EndTime        string `json:"EndTime" xml:"EndTime"`
	BackupMode     string `json:"BackupMode" xml:"BackupMode"`
	Status         string `json:"Status" xml:"Status"`
	IsDBReplaced   string `json:"IsDBReplaced" xml:"IsDBReplaced"`
	Description    string `json:"Description" xml:"Description"`
}

// CreateDescribeMigrateTaskByIdRequest creates a request to invoke DescribeMigrateTaskById API
func CreateDescribeMigrateTaskByIdRequest() (request *DescribeMigrateTaskByIdRequest) {
	request = &DescribeMigrateTaskByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeMigrateTaskById", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeMigrateTaskByIdResponse creates a response to parse from DescribeMigrateTaskById response
func CreateDescribeMigrateTaskByIdResponse() (response *DescribeMigrateTaskByIdResponse) {
	response = &DescribeMigrateTaskByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
