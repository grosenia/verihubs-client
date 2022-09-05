package verihubsgo

import (
	"bytes"
	"encoding/json"

	"github.com/nbs-go/clog"
	_ "github.com/nbs-go/clogrus"
)

// SmsGateway struct
type WhatsappGateway struct {
	Client Client
}

// SendSmsOtp send sms for OTP
func (gateway *WhatsappGateway) SendWhatsAppOtp(req *VerihubsWhatsappOtpRequest) (response *VerihubsWhatsappOtpResponse, err error) {
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

	if httpStatus != 200 {
		resp.ErrorStatus = true
	} else {
		if resp.Code != VerihubsDelivered {
			resp.ErrorStatus = true
		} else {
			resp.ErrorStatus = false
		}
	}

	return &resp, nil
}

// VerifySmsOtp Verify SMS OTP
func (gateway *WhatsappGateway) VerifyOtp(req *VerihubsOtpVerifyRequest) (response *VerihubsOtpVerifyResponse, err error) {
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

	if httpStatus != 200 {
		resp.ErrorStatus = true
	} else {
		if resp.Code != VerihubsDelivered {
			resp.ErrorStatus = true
		} else {
			resp.ErrorStatus = false
		}
	}

	return &resp, nil
}
