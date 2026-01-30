package main

import (
	"fmt"

	verihubsgo "github.com/grosenia/verihubs-client/v2"
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

	MSISDN := "628372873827382"
	LangCode := "id"
	TemplateName := "notiffication_order"

	// Create the content object
	content := verihubsgo.ContentWhatsapp{
		BodyParams:   []string{"Reyvin", "209182918291089", "Rp 100.000", "20 Desember 2025", "http://xendit.com/aojsasjaiojsa"},
		HeaderParams: []string{"string"},
		ButtonParam:  verihubsgo.ButtonParam{},
	}

	var request = &verihubsgo.VerihubsWhatsappMessageRequest{
		MSISDN:   MSISDN,
		LangCode: LangCode,
		Template: TemplateName,
		Content:  content,
	}

	resp, err := smsGateway.RequestWhatsappMessage(request)

	if err != nil {
		fmt.Println("Error server")
		return
	}

	if resp.ErrorStatus {
		// Ada error
		fmt.Println("Citcall response: ")
		fmt.Println("Error: ", resp.Error())
		return
	}

	fmt.Println("Citcall response: ", resp)
	// fmt.Println(resp)
	fmt.Println("Code: ", resp.Code)
	fmt.Println("SessionId: ", resp.SessionId)

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
