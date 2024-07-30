package verihubsgo

const (
	VerihubsSuccess        = 200
	VerihubsSuccessRequest = 201
	VerihubsBadRequest     = 400
	VerihubsForbidden      = 403
	VerihubsConflict       = 409
	VerihubsTooManyReq     = 429
	VerihubsInternalError  = 500
)

// VerihubsSmsOtpRequest Verihubs SMSOTP Request
type VerihubsSmsOtpRequest struct {
	MSISDN    string `json:"msisdn"`
	Template  string `json:"template"`
	OTP       string `json:"otp"`
	TimeLimit string `json:"time_limit"`
	Challenge string `json:"challenge"`
}

// VerifySmsOtpVerihubs
type VerihubsOtpVerifyRequest struct {
	OTP       string `json:"otp"`
	MSISDN    string `json:"msisdn"`
	Challenge string `json:"challenge"`
}

// VerihubsWhatsappOtpRequest Verihubs SMSOTP Request
type VerihubsWhatsappOtpRequest struct {
	MSISDN string `json:"msisdn"`
	//Content      []string `json:"content"`
	OTP          string `json:"otp"`
	TimeLimit    string `json:"time_limit"`
	Challenge    string `json:"challenge"`
	LangCode     string `json:"lang_code"`
	TemplateName string `json:"template_name"`
	//OtpLength    string   `json:"otp_length"`
	//CallbackUrl string `json:"callback_url"`
}

// VerihubsMiscallOtpRequest Verihubs misscallOTP Request
type VerihubsMisscallOtpRequest struct {
	MSISDN    string `json:"msisdn"`
	OTP       string `json:"otp"`
	TimeLimit string `json:"time_limit"`
	Challenge string `json:"challenge"`
}
