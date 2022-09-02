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
	Sandbox:    "https://api.verihubs.com/v1",
	Production: "https://api.verihubs.com/v1",
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
	return e.String() + "/otp/send"
}

// CreateSmsOtpVerifyURL : Create VERIFY URL
func (e EnvironmentType) CreateSmsOtpVerifyURL() string {
	return e.String() + "/otp/verify"
}

// CreateWhatsappOtpURL : Create SMSOTP URL
func (e EnvironmentType) CreateWhatsappOtpURL() string {
	return e.String() + "/whatsapp/otp/send"
}

// CreateWhatsappOtpVerifyURL : Create VERIFY URL
func (e EnvironmentType) CreateWhatsappOtpVerifyURL() string {
	return e.String() + "/whatsapp/otp/verify"
}
