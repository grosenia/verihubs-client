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

	// Example
	MSISDN := "62839102839"
	Template := "Kode OTP Anda $OTP berlaku 5 menit"
	OTP := "234123"
	TimeLimit := "300" // 5 minutes
	Challenge := "update_account"

	var request = &verihubsgo.VerihubsSmsOtpRequest{
		MSISDN:    MSISDN,
		Template:  Template,
		OTP:       OTP,
		Challenge: Challenge,
		TimeLimit: TimeLimit,
	}

	resp, err := smsGateway.SendSmsOtp(request)

	if err != nil {
		fmt.Println("Error server")
		return
	}

	if resp.ErrorStatus && resp.Code != 201 {
		// Ada error
		fmt.Println("Error: ", resp.Error())
		return
	}

	fmt.Println("Verihubs response: ")
	fmt.Println(resp)
	fmt.Println("TrxID: ", resp.SessionId)
	fmt.Println("Token: ", resp.OTP)

	fmt.Println("Test Verify Fail: ")

	var verifyRequest = &verihubsgo.VerihubsSmsOtpVerifyRequest{
		MSISDN:    MSISDN,
		OTP:       OTP,
		Challenge: Challenge,
	}
	respVerify, err := smsGateway.VerifySmsOtp(verifyRequest)

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
	verifyRequest = &verihubsgo.VerihubsSmsOtpVerifyRequest{
		MSISDN:    MSISDN,
		OTP:       OTP,
		Challenge: Challenge,
	}
	respVerifySuccess, err := smsGateway.VerifySmsOtp(verifyRequest)

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
