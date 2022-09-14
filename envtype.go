package verihubsgo

// EnvironmentType value
type EnvironmentType int8

const (
	_ EnvironmentType = iota

	// Sandbox : represent sandbox environment
	Sandbox

	// Production : represent production environment
	Production
)

// TODO should read from config
var typeString = map[EnvironmentType]string{
	Sandbox:    "https://api.verihubs.com",
	Production: "https://api.verihubs.com",
}

// implement stringer
func (e EnvironmentType) String() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}

// CreateSmsOtpURL : Create SMSOTP URL
func (e EnvironmentType) CreateSmsOtpURL() string {
	return e.String() + "/v1/otp/send"
}

// CreateSmsOtpVerifyURL : Create VERIFY URL
func (e EnvironmentType) CreateSmsOtpVerifyURL() string {
	return e.String() + "/v1/otp/verify"
}

// CreateWhatsappOtpURL : Create SMSOTP URL
func (e EnvironmentType) CreateWhatsappOtpURL() string {
	return e.String() + "/v1/whatsapp/otp/send"
}

// CreateWhatsappOtpVerifyURL : Create VERIFY URL
func (e EnvironmentType) CreateWhatsappOtpVerifyURL() string {
	return e.String() + "/v1/whatsapp/otp/verify"
}

// CreateMisscallURL : Create Misscall URL
func (e EnvironmentType) CreateMiscalllUrl() string {
	return e.String() + "/v2/flashcall/send"
}

// CreateMisscallVerifyURL : Create Misscall verify URL
func (e EnvironmentType) CreateMiscallVerifyUrl() string {
	return e.String() + "/v2/flashcall/verify"
}
