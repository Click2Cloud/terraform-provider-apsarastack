package polardb

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

// DescribeBackupPolicy invokes the polardb.DescribeBackupPolicy API synchronously
// api document: https://help.aliyun.com/api/polardb/describebackuppolicy.html
func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (response *DescribeBackupPolicyResponse, err error) {
	response = CreateDescribeBackupPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupPolicyWithChan invokes the polardb.DescribeBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/polardb/describebackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackupPolicyWithChan(request *DescribeBackupPolicyRequest) (<-chan *DescribeBackupPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupPolicy(request)
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

// DescribeBackupPolicyWithCallback invokes the polardb.DescribeBackupPolicy API asynchronously
// api document: https://help.aliyun.com/api/polardb/describebackuppolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeBackupPolicyWithCallback(request *DescribeBackupPolicyRequest, callback func(response *DescribeBackupPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupPolicy(request)
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

// DescribeBackupPolicyRequest is the request struct for api DescribeBackupPolicy
type DescribeBackupPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeBackupPolicyResponse is the response struct for api DescribeBackupPolicy
type DescribeBackupPolicyResponse struct {
	*responses.BaseResponse
	RequestId                              string `json:"RequestId" xml:"RequestId"`
	BackupRetentionPeriod                  int    `json:"BackupRetentionPeriod" xml:"BackupRetentionPeriod"`
	PreferredNextBackupTime                string `json:"PreferredNextBackupTime" xml:"PreferredNextBackupTime"`
	PreferredBackupTime                    string `json:"PreferredBackupTime" xml:"PreferredBackupTime"`
	PreferredBackupPeriod                  string `json:"PreferredBackupPeriod" xml:"PreferredBackupPeriod"`
	DataLevel1BackupRetentionPeriod        string `json:"DataLevel1BackupRetentionPeriod" xml:"DataLevel1BackupRetentionPeriod"`
	DataLevel2BackupRetentionPeriod        string `json:"DataLevel2BackupRetentionPeriod" xml:"DataLevel2BackupRetentionPeriod"`
	BackupRetentionPolicyOnClusterDeletion string `json:"BackupRetentionPolicyOnClusterDeletion" xml:"BackupRetentionPolicyOnClusterDeletion"`
}

// CreateDescribeBackupPolicyRequest creates a request to invoke DescribeBackupPolicy API
func CreateDescribeBackupPolicyRequest() (request *DescribeBackupPolicyRequest) {
	request = &DescribeBackupPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeBackupPolicy", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupPolicyResponse creates a response to parse from DescribeBackupPolicy response
func CreateDescribeBackupPolicyResponse() (response *DescribeBackupPolicyResponse) {
	response = &DescribeBackupPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
