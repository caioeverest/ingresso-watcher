package client

import "github.com/caioeverest/ingresso-watcher/config"

type IrInterface interface {
	GetEventById(conf *config.Config, id string) ([]map[string]interface{}, error)
}

type WhatsAppInterface interface {
	Send(phoneNumber, text string) (string, error)
}
