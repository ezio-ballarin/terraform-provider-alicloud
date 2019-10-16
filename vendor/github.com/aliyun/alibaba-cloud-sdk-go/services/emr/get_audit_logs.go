package emr

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

// GetAuditLogs invokes the emr.GetAuditLogs API synchronously
// api document: https://help.aliyun.com/api/emr/getauditlogs.html
func (client *Client) GetAuditLogs(request *GetAuditLogsRequest) (response *GetAuditLogsResponse, err error) {
	response = CreateGetAuditLogsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAuditLogsWithChan invokes the emr.GetAuditLogs API asynchronously
// api document: https://help.aliyun.com/api/emr/getauditlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuditLogsWithChan(request *GetAuditLogsRequest) (<-chan *GetAuditLogsResponse, <-chan error) {
	responseChan := make(chan *GetAuditLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAuditLogs(request)
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

// GetAuditLogsWithCallback invokes the emr.GetAuditLogs API asynchronously
// api document: https://help.aliyun.com/api/emr/getauditlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAuditLogsWithCallback(request *GetAuditLogsRequest, callback func(response *GetAuditLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAuditLogsResponse
		var err error
		defer close(result)
		response, err = client.GetAuditLogs(request)
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

// GetAuditLogsRequest is the request struct for api GetAuditLogs
type GetAuditLogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageCount       requests.Integer `position:"Query" name:"PageCount"`
	OrderMode       string           `position:"Query" name:"OrderMode"`
	EntityId        string           `position:"Query" name:"EntityId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	Limit           requests.Integer `position:"Query" name:"Limit"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentSize     requests.Integer `position:"Query" name:"CurrentSize"`
	OrderField      string           `position:"Query" name:"OrderField"`
	Operation       string           `position:"Query" name:"Operation"`
}

// GetAuditLogsResponse is the response struct for api GetAuditLogs
type GetAuditLogsResponse struct {
	*responses.BaseResponse
	RequestId  string              `json:"RequestId" xml:"RequestId"`
	PageNumber int                 `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                 `json:"PageSize" xml:"PageSize"`
	TotalCount int                 `json:"TotalCount" xml:"TotalCount"`
	Items      ItemsInGetAuditLogs `json:"Items" xml:"Items"`
}

// CreateGetAuditLogsRequest creates a request to invoke GetAuditLogs API
func CreateGetAuditLogsRequest() (request *GetAuditLogsRequest) {
	request = &GetAuditLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "GetAuditLogs", "emr", "openAPI")
	return
}

// CreateGetAuditLogsResponse creates a response to parse from GetAuditLogs response
func CreateGetAuditLogsResponse() (response *GetAuditLogsResponse) {
	response = &GetAuditLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}