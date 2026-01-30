package verihubsgo

import (
	"bytes"
	"encoding/json"

	"github.com/nbs-go/clog"
	_ "github.com/nbs-go/clogrus"
)

// RequestMiscallOtp request miscall for OTP
func (gateway *SmsGateway) RequestWhatsappMessage(req *VerihubsWhatsappMessageRequest) (response *VerihubsMessageWhatsappResponse, err error) {
	log := clog.Get()
	resp := VerihubsMessageWhatsappResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateMessageWhatsapp()
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

func (gateway *SmsGateway) SendWhatsAppOtp(req *VerihubsWhatsappOtpRequest) (response *VerihubsWhatsappOtpResponse, err error) {
	log := clog.Get()
	resp := VerihubsWhatsappOtpResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateWhatsappOtpURL()
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

// VerifySmsOtp Verify SMS OTP
func (gateway *SmsGateway) VerifyOtp(req *VerihubsOtpVerifyRequest) (response *VerihubsOtpVerifyResponse, err error) {
	log := clog.Get()
	resp := VerihubsOtpVerifyResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateWhatsappOtpVerifyURL()
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
