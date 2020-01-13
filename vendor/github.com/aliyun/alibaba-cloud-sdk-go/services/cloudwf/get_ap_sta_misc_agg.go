package cloudwf

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

// GetApStaMiscAgg invokes the cloudwf.GetApStaMiscAgg API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getapstamiscagg.html
func (client *Client) GetApStaMiscAgg(request *GetApStaMiscAggRequest) (response *GetApStaMiscAggResponse, err error) {
	response = CreateGetApStaMiscAggResponse()
	err = client.DoAction(request, response)
	return
}

// GetApStaMiscAggWithChan invokes the cloudwf.GetApStaMiscAgg API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getapstamiscagg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApStaMiscAggWithChan(request *GetApStaMiscAggRequest) (<-chan *GetApStaMiscAggResponse, <-chan error) {
	responseChan := make(chan *GetApStaMiscAggResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApStaMiscAgg(request)
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

// GetApStaMiscAggWithCallback invokes the cloudwf.GetApStaMiscAgg API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getapstamiscagg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApStaMiscAggWithCallback(request *GetApStaMiscAggRequest, callback func(response *GetApStaMiscAggResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApStaMiscAggResponse
		var err error
		defer close(result)
		response, err = client.GetApStaMiscAgg(request)
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

// GetApStaMiscAggRequest is the request struct for api GetApStaMiscAgg
type GetApStaMiscAggRequest struct {
	*requests.RpcRequest
	ApgroupId requests.Integer `position:"Query" name:"ApgroupId"`
}

// GetApStaMiscAggResponse is the response struct for api GetApStaMiscAgg
type GetApStaMiscAggResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetApStaMiscAggRequest creates a request to invoke GetApStaMiscAgg API
func CreateGetApStaMiscAggRequest() (request *GetApStaMiscAggRequest) {
	request = &GetApStaMiscAggRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetApStaMiscAgg", "cloudwf", "openAPI")
	return
}

// CreateGetApStaMiscAggResponse creates a response to parse from GetApStaMiscAgg response
func CreateGetApStaMiscAggResponse() (response *GetApStaMiscAggResponse) {
	response = &GetApStaMiscAggResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}