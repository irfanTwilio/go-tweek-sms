package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/piquette/finance-go/quote"
	"github.com/sirupsen/logrus"
	twilio "github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func sms(message string) {
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody(message)

	_, err := client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		logrus.Fatalf("Input one stock symbol", os.Args[0])
	}
	symbol := flag.Args()[0]
	a, _ := quote.Get(symbol)
	fmt.Printf(" --%v--\n", a.ShortName)
	fmt.Printf("Current Price: $%v\n", a.Ask)
	s := fmt.Sprintf("Current Price: $%v\n", a.Ask)
	sms(s)
}
