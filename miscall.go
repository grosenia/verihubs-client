package verihubsgo

import (
	"bytes"
	"encoding/json"

	"github.com/nbs-go/clog"
	_ "github.com/nbs-go/clogrus"
)

// RequestMiscallOtp request miscall for OTP
func (gateway *SmsGateway) RequestMiscallOtp(req *VerihubsMisscallOtpRequest) (response *VerihubsMisscallOtpResponse, err error) {
	log := clog.Get()
	resp := VerihubsMisscallOtpResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateMiscalllUrl()
	httpRequest, err := gateway.Client.NewRequest("POST", path, bytes.NewBuffer(jsonReq))

	if err != nil {
		return nil, err
	}

	httpStatus, err := gateway.Client.ExecuteRequest(httpRequest, &resp)
	if err != nil {
		log.Error("Error charging ", err)
		return nil, err
	}

	log.Debugf("HTTP STATUS Response [%d]", httpStatus)
	if httpStatus != VerihubsSuccessRequest {
		resp.Code = httpStatus
		resp.ErrorStatus = true
	} else {
		resp.Code = httpStatus
		resp.ErrorStatus = false
	}

	return &resp, nil
}

// RequestMiscallOtp request miscall for OTP
func (gateway *SmsGateway) RequestMiscallVerifyOtp(req *VerihubsOtpVerifyRequest) (response *VerihubsOtpVerifyResponse, err error) {
	log := clog.Get()
	resp := VerihubsOtpVerifyResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateMiscallVerifyUrl()
	httpRequest, err := gateway.Client.NewRequest("POST", path, bytes.NewBuffer(jsonReq))

	if err != nil {
		return nil, err
	}

	httpStatus, err := gateway.Client.ExecuteRequest(httpRequest, &resp)
	if err != nil {
		log.Error("Error charging ", err)
		return nil, err
	}

	log.Debugf("HTTP STATUS Response [%d]", httpStatus)
	if httpStatus != VerihubsSuccess {
		resp.Code = httpStatus
		resp.ErrorStatus = true
	} else {
		resp.Code = httpStatus
		resp.ErrorStatus = false
	}

	return &resp, nil
}
