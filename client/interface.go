package client

import "github.com/caioeverest/ingressoWatcher/config"

type IrInterface interface {
	GetEventById(conf *config.Config, id string) ([]map[string]interface{}, error)
}

type WhatsAppInterface interface {
	Send(phoneNumber, text string) (string, error)
}
