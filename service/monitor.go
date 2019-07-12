package service

import (
	"fmt"
	"log"

	"github.com/caioeverest/ingresso-watcher/client"
	"github.com/caioeverest/ingresso-watcher/config"
	"github.com/caioeverest/ingresso-watcher/repository"
	"github.com/jasonlvhit/gocron"
)

func Monitor(contactList, eventList repository.Interface,
	whatsapp client.WhatsAppInterface, conf *config.Config) {

	gocron.Every(conf.Interval).Seconds().Do(monitorDeEventos, contactList, eventList, whatsapp, conf)

	<-gocron.Start()
}

func monitorDeEventos(contactList, eventList repository.Interface, whatsapp client.WhatsAppInterface, conf *config.Config) {
	events := GetAllRegistredEvents(eventList)

	for _, event := range events {
		_, err := CheckIfHaveTickets(conf, event.Id)
		if err == nil {
			contacts := GetContactList(contactList)
			url := fmt.Sprintf("https://www.ingressorapido.com.br/event/%s-1", event.Id)
			for _, contact := range contacts {
				if err := SendFoundTicketMessage(whatsapp, contact.Phone, contact.Name, event.Name, url); err != nil {
					log.Printf("Ocorreu um erro %s ao enviar a notificação para %s - %s", err, contact.Phone, contact.Name)
				}
			}
		}
	}

}
