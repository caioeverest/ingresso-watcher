package service

import (
	"fmt"
	"log"

	"github.com/caioeverest/ingressoWatcher/client"
	"github.com/caioeverest/ingressoWatcher/config"
	"github.com/caioeverest/ingressoWatcher/repository"
	"github.com/jasonlvhit/gocron"
)

func Monitor(contactList, eventList repository.Interface,
	whatsapp client.WhatsAppInterface, conf *config.Config) {

	gocron.Every(conf.Interval).Seconds().Do(monitorDeEventos, contactList, eventList, whatsapp, conf)

	<-gocron.Start()
}

func monitorDeEventos(contactList, eventList repository.Interface, whatsapp client.WhatsAppInterface, conf *config.Config) {
	events := GetAllRegistredEvents(eventList)

	for eventId, eventName := range events {
		_, err := CheckIfHaveTickets(conf, eventId)
		if err == nil {
			contacts := GetContactList(contactList)
			url := fmt.Sprintf("https://www.ingressorapido.com.br/event/%s-1", eventId)
			for phone, name := range contacts {
				if err := SendFoundTicketMessage(whatsapp, phone, name, eventName, url); err != nil {
					log.Printf("Ocorreu um erro %s ao enviar a notificação para %s - %s", err, phone, name)
				}
			}
		}
	}

}
