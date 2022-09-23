package main

import (
	"fmt"

	verihubsgo "github.com/grosenia/verihubs-client"
	viper "github.com/spf13/viper"
)

var verihubsClient verihubsgo.Client
var smsGateway verihubsgo.SmsGateway

func main() {
	fmt.Println("Load Config...")

	viper.SetConfigType("props")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	fmt.Println("Load Config success")
	fmt.Println("Setup client")

	setupClient()

	MSISDN := "6287771311133"
	OTP := "9021"
	TimeLimit := "300" // 5 minutes
	Challenge := "update_account"

	var request = &verihubsgo.VerihubsMisscallOtpRequest{
		MSISDN:    MSISDN,
		OTP:       OTP,
		Challenge: Challenge,
		TimeLimit: TimeLimit,
	}

	resp, err := smsGateway.RequestMiscallOtp(request)

	if err != nil {
		fmt.Println("Error server")
		return
	}

	if resp.ErrorStatus {
		// Ada error
		fmt.Println("Error: ", resp.Error())
		return
	}

	fmt.Println("Citcall response: ", resp)
	// fmt.Println(resp)
	//fmt.Println("Code: ", resp.Code)
	fmt.Println("SessionId: ", resp.Data.SessionId)

	var verifyRequest = &verihubsgo.VerihubsOtpVerifyRequest{
		MSISDN:    MSISDN,
		OTP:       OTP,
		Challenge: Challenge,
	}
	respVerify, err := smsGateway.RequestMiscallVerifyOtp(verifyRequest)

	if err != nil {
		fmt.Println("Error server")
	}

	if respVerify.ErrorStatus {
		// Ada error
		fmt.Println("Error: ", respVerify.Error())
	}

	fmt.Println("Verihubs response: ")
	fmt.Println(respVerify)
	fmt.Println("Error Status: ", respVerify.ErrorStatus)
	fmt.Println("Error Code: ", respVerify.Code)
	fmt.Println("Error Info: ", respVerify.Message)

	fmt.Println("Test Verify Success: ")
	verifyRequest = &verihubsgo.VerihubsOtpVerifyRequest{
		MSISDN:    MSISDN,
		OTP:       "9021",
		Challenge: Challenge,
	}
	respVerifySuccess, err := smsGateway.RequestMiscallVerifyOtp(verifyRequest)

	if err != nil {
		fmt.Println("Error server")
	}

	if respVerifySuccess.ErrorStatus {
		// Ada error
		fmt.Println("Error: ", respVerifySuccess.Error())
	}

	fmt.Println("Verihubs response: ")
	fmt.Println(respVerifySuccess)
	fmt.Println("Error Status: ", respVerifySuccess.ErrorStatus)
	fmt.Println("Error Code: ", respVerifySuccess.Code)
	fmt.Println("Error Info: ", respVerifySuccess.Message)

	//
}

func setupClient() {
	verihubsClient = verihubsgo.NewClient()

	// TODO: should put in config file
	verihubsClient.AppId = viper.GetString("App-ID")
	verihubsClient.AppKey = viper.GetString("API-Key")
	verihubsClient.APIEnvType = verihubsgo.Sandbox
	verihubsClient.LogLevel = 3

	smsGateway = verihubsgo.SmsGateway{
		Client: verihubsClient,
	}
}
