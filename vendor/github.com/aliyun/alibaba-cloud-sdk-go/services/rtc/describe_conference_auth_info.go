package rtc

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

// DescribeConferenceAuthInfo invokes the rtc.DescribeConferenceAuthInfo API synchronously
// api document: https://help.aliyun.com/api/rtc/describeconferenceauthinfo.html
func (client *Client) DescribeConferenceAuthInfo(request *DescribeConferenceAuthInfoRequest) (response *DescribeConferenceAuthInfoResponse, err error) {
	response = CreateDescribeConferenceAuthInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConferenceAuthInfoWithChan invokes the rtc.DescribeConferenceAuthInfo API asynchronously
// api document: https://help.aliyun.com/api/rtc/describeconferenceauthinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConferenceAuthInfoWithChan(request *DescribeConferenceAuthInfoRequest) (<-chan *DescribeConferenceAuthInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeConferenceAuthInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConferenceAuthInfo(request)
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

// DescribeConferenceAuthInfoWithCallback invokes the rtc.DescribeConferenceAuthInfo API asynchronously
// api document: https://help.aliyun.com/api/rtc/describeconferenceauthinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeConferenceAuthInfoWithCallback(request *DescribeConferenceAuthInfoRequest, callback func(response *DescribeConferenceAuthInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConferenceAuthInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeConferenceAuthInfo(request)
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

// DescribeConferenceAuthInfoRequest is the request struct for api DescribeConferenceAuthInfo
type DescribeConferenceAuthInfoRequest struct {
	*requests.RpcRequest
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	ConferenceId string           `position:"Query" name:"ConferenceId"`
	AppId        string           `position:"Query" name:"AppId"`
}

// DescribeConferenceAuthInfoResponse is the response struct for api DescribeConferenceAuthInfo
type DescribeConferenceAuthInfoResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	ConferenceId string   `json:"ConferenceId" xml:"ConferenceId"`
	AuthInfo     AuthInfo `json:"AuthInfo" xml:"AuthInfo"`
}

// CreateDescribeConferenceAuthInfoRequest creates a request to invoke DescribeConferenceAuthInfo API
func CreateDescribeConferenceAuthInfoRequest() (request *DescribeConferenceAuthInfoRequest) {
	request = &DescribeConferenceAuthInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "DescribeConferenceAuthInfo", "", "")
	return
}

// CreateDescribeConferenceAuthInfoResponse creates a response to parse from DescribeConferenceAuthInfo response
func CreateDescribeConferenceAuthInfoResponse() (response *DescribeConferenceAuthInfoResponse) {
	response = &DescribeConferenceAuthInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}