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

// CreateAccount invokes the adb.CreateAccount API synchronously
// api document: https://help.aliyun.com/api/adb/createaccount.html
func (client *Client) CreateAccount(request *CreateAccountRequest) (response *CreateAccountResponse, err error) {
	response = CreateCreateAccountResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAccountWithChan invokes the adb.CreateAccount API asynchronously
// api document: https://help.aliyun.com/api/adb/createaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccountWithChan(request *CreateAccountRequest) (<-chan *CreateAccountResponse, <-chan error) {
	responseChan := make(chan *CreateAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAccount(request)
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

// CreateAccountWithCallback invokes the adb.CreateAccount API asynchronously
// api document: https://help.aliyun.com/api/adb/createaccount.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAccountWithCallback(request *CreateAccountRequest, callback func(response *CreateAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAccountResponse
		var err error
		defer close(result)
		response, err = client.CreateAccount(request)
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

// CreateAccountRequest is the request struct for api CreateAccount
type CreateAccountRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountType          string           `position:"Query" name:"AccountType"`
	AccountDescription   string           `position:"Query" name:"AccountDescription"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
}

// CreateAccountResponse is the response struct for api CreateAccount
type CreateAccountResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	TaskId      int    `json:"TaskId" xml:"TaskId"`
}

// CreateCreateAccountRequest creates a request to invoke CreateAccount API
func CreateCreateAccountRequest() (request *CreateAccountRequest) {
	request = &CreateAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "CreateAccount", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAccountResponse creates a response to parse from CreateAccount response
func CreateCreateAccountResponse() (response *CreateAccountResponse) {
	response = &CreateAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
