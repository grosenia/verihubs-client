package verihubsgo

import (
	"bytes"
	"encoding/json"

	"github.com/nbs-go/clog"
	_ "github.com/nbs-go/clogrus"
)

// SmsGateway struct
type SmsGateway struct {
	Client Client
}

// SendSmsOtp send sms for OTP
func (gateway *SmsGateway) SendSmsOtp(req *VerihubsSmsOtpRequest) (response *VerihubsSmsOtpResponse, err error) {
	log := clog.Get()
	resp := VerihubsSmsOtpResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateSmsOtpURL()
	httpRequest, err := gateway.Client.NewRequest("POST", path, bytes.NewBuffer(jsonReq))

	if err != nil {
		return nil, err
	}

	httpStatus, err := gateway.Client.ExecuteRequest(httpRequest, &resp)
	if err != nil {
		log.Error("Error charging ", err)
		return nil, err
	}
	log.Debugf("STATUSS", httpStatus)
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
func (gateway *SmsGateway) VerifySmsOtp(req *VerihubsOtpVerifyRequest) (response *VerihubsOtpVerifyResponse, err error) {
	log := clog.Get()
	resp := VerihubsOtpVerifyResponse{}
	jsonReq, _ := json.Marshal(req)

	path := gateway.Client.APIEnvType.CreateSmsOtpVerifyURL()
	httpRequest, err := gateway.Client.NewRequest("POST", path, bytes.NewBuffer(jsonReq))

	if err != nil {
		return nil, err
	}

	httpStatus, err := gateway.Client.ExecuteRequest(httpRequest, &resp)
	if err != nil {
		log.Error("Error charging ", err)
		return nil, err
	}
	if httpStatus != VerihubsSuccess {
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

	if httpStatus != VerihubsSuccess {
		resp.Code = httpStatus
		resp.ErrorStatus = true
	} else {
		resp.Code = httpStatus
		resp.ErrorStatus = false
	}

	return &resp, nil
}
