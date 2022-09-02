package verihubsgo

const (
	VerihubsRequested    = 0
	VerihubsDelivered    = 1
	VerihubsFailed       = 4
	VerihubsRequestError = 5
	VerihubsUndelivered  = 7
	VerihubsROTDR        = 8
	VErihubsBlocked      = 9
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
type VerihubsSmsOtpVerifyRequest struct {
	OTP       string `json:"otp"`
	MSISDN    string `json:"msisdn"`
	Challenge string `json:"challenge"`
}

// VerihubsWhatsappOtpRequest Verihubs SMSOTP Request
type VerihubsWhatsappOtpRequest struct {
	MSISDN       string   `json:"msisdn"`
	Content      []string `json:"content"`
	OTP          string   `json:"otp"`
	TimeLimit    string   `json:"time_limit"`
	Challenge    string   `json:"challenge"`
	LangCode     string   `json:"lang_code"`
	TemplateName string   `json:"template_name"`
	//OtpLength    string   `json:"otp_length"`
	//CallbackUrl string `json:"callback_url"`
}

// VerifyWhatsappOtpVerihubs
type VerihubsWhatsappOtpVerifyRequest struct {
	OTP       string `json:"otp"`
	MSISDN    string `json:"msisdn"`
	Challenge string `json:"challenge"`
}
