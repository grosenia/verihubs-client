package verihubsgo

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// A FlexInt is an int that can be unmarshalled from a JSON field
// that has either a number or a string value.
// E.g. if the json field contains an string "42", the
// FlexInt value will be "42".
type FlexInt int

// UnmarshalJSON implements the json.Unmarshaler interface, which
// allows us to ingest values of any json type as an int and run our custom conversion
func (fi *FlexInt) UnmarshalJSON(b []byte) error {
	if b[0] != '"' {
		return json.Unmarshal(b, (*int)(fi))
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*fi = FlexInt(i)
	return nil
}

// VerihubsSmsOtpResponse Verihubs SMSOTP Response
type VerihubsSmsOtpResponse struct {
	Code         FlexInt `json:"code"`
	Message      string  `json:"message"`
	OTP          string  `json:"otp"`
	MSISDN       string  `json:"msisdn"`
	SessionId    string  `json:"session_id"`
	TryCount     int     `json:"try_count"`
	SegmentCount int     `json:"segment_count"`
	ErrorStatus  bool    `json:"-"`
}

func (e VerihubsSmsOtpResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// VerihubsSmsOtpVerifyResponse Verihubs SMSOTP Verify Response
type VerihubsOtpVerifyResponse struct {
	Code        FlexInt `json:"code"`
	Message     string  `json:"message"`
	ErrorStatus bool    `json:"-"`
}

func (e VerihubsOtpVerifyResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// VerihubsWhatsappOtpResponse Verihubs WhatsappOTP Response
type VerihubsWhatsappOtpResponse struct {
	Code        FlexInt `json:"code"`
	Message     string  `json:"message"`
	OTP         string  `json:"otp"`
	MSISDN      string  `json:"msisdn"`
	SessionId   string  `json:"session_id"`
	TryCount    int     `json:"try_count"`
	ErrorStatus bool    `json:"-"`
}

func (e VerihubsWhatsappOtpResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// VerihubsSmsOtpResponse Verihubs SMSOTP Response
type VerihubsMisscallOtpResponse struct {
	Code        FlexInt `json:"code"`
	Message     string  `json:"message"`
	OTP         string  `json:"otp"`
	MSISDN      string  `json:"msisdn"`
	SessionId   string  `json:"session_id"`
	TryCount    int     `json:"try_count"`
	ErrorStatus bool    `json:"-"`
}

func (e VerihubsMisscallOtpResponse) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}
