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

// ScaleCluster invokes the cs.ScaleCluster API synchronously
func (client *Client) ScaleCluster(request *ScaleClusterRequest) (response *ScaleClusterResponse, err error) {
	response = CreateScaleClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ScaleClusterWithChan invokes the cs.ScaleCluster API asynchronously
func (client *Client) ScaleClusterWithChan(request *ScaleClusterRequest) (<-chan *ScaleClusterResponse, <-chan error) {
	responseChan := make(chan *ScaleClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScaleCluster(request)
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

// ScaleClusterWithCallback invokes the cs.ScaleCluster API asynchronously
func (client *Client) ScaleClusterWithCallback(request *ScaleClusterRequest, callback func(response *ScaleClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScaleClusterResponse
		var err error
		defer close(result)
		response, err = client.ScaleCluster(request)
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

// ScaleClusterRequest is the request struct for api ScaleCluster
type ScaleClusterRequest struct {
	*requests.RoaRequest
	KeyPair                  string           `position:"Body" name:"key_pair"`
	WorkerDataDisk           requests.Boolean `position:"Body" name:"worker_data_disk"`
	Count                    requests.Integer `position:"Body" name:"count"`
	WorkerSystemDiskCategory string           `position:"Body" name:"worker_system_disk_category"`
	CloudMonitorFlags        requests.Boolean `position:"Body" name:"cloud_monitor_flags"`
	ClusterId                string           `position:"Path" name:"ClusterId"`
	WorkerPeriodUnit         string           `position:"Body" name:"worker_period_unit"`
	WorkerAutoRenew          requests.Boolean `position:"Body" name:"worker_auto_renew"`
	WorkerAutoRenewPeriod    requests.Integer `position:"Body" name:"worker_auto_renew_period"`
	WorkerPeriod             requests.Integer `position:"Body" name:"worker_period"`
	LoginPassword            string           `position:"Body" name:"login_password"`
	WorkerSystemDiskSize     requests.Integer `position:"Body" name:"worker_system_disk_size"`
	CpuPolicy                string           `position:"Body" name:"cpu_policy"`
	DisableRollback          requests.Boolean `position:"Body" name:"disable_rollback"`
	WorkerInstanceChargeType string           `position:"Body" name:"worker_instance_charge_type"`
}

// ScaleClusterResponse is the response struct for api ScaleCluster
type ScaleClusterResponse struct {
	*responses.BaseResponse
	ClusterId string `json:"cluster_id" xml:"cluster_id"`
	TaskId    string `json:"task_id" xml:"task_id"`
	RequestId string `json:"request_id" xml:"request_id"`
}

// CreateScaleClusterRequest creates a request to invoke ScaleCluster API
func CreateScaleClusterRequest() (request *ScaleClusterRequest) {
	request = &ScaleClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ScaleCluster", "/clusters/[ClusterId]", "", "")
	request.Method = requests.PUT
	return
}

// CreateScaleClusterResponse creates a response to parse from ScaleCluster response
func CreateScaleClusterResponse() (response *ScaleClusterResponse) {
	response = &ScaleClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}