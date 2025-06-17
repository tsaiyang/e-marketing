package ioc

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/spf13/viper"
)

func InitSendgrid() *sendgrid.Client {
	client := sendgrid.NewSendClient(viper.GetString("sendgrid"))
	if client == nil {
		panic("new sendgrid client failed")
	}

	return client
}
