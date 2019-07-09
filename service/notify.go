package service

import (
	"log"

	"github.com/caioeverest/ingressoWatcher/client"
	"github.com/caioeverest/ingressoWatcher/service/templates"
)

func SendTestMessage(wpp client.WhatsAppInterface, phone string) error {
	data := map[string]string{
		"phone": phone,
	}
	message, err := templates.FormatMessage("test", data)
	if err != nil {
		return err
	}

	res, err := wpp.Send(phone, message)
	log.Print(res)
	return err
}

func SendGreetingsMessage(wpp client.WhatsAppInterface, phone, name string) error {
	data := map[string]string{
		"phone": phone,
		"name":  name,
	}
	message, err := templates.FormatMessage("greetings", data)
	if err != nil {
		return err
	}

	res, err := wpp.Send(phone, message)
	log.Print(res)
	return err
}

func SendFoundTicketMessage(wpp client.WhatsAppInterface, phone, name, eventName, url string) error {
	data := map[string]string{
		"phone":     phone,
		"name":      name,
		"eventName": eventName,
		"url":       url,
	}
	message, err := templates.FormatMessage("found_tickets", data)
	if err != nil {
		return err
	}

	res, err := wpp.Send(phone, message)
	log.Print(res)
	return err
}
