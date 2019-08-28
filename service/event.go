package service

import (
	"fmt"
	"log"
	"time"

	"github.com/caioeverest/ingresso-watcher/client"
	"github.com/caioeverest/ingresso-watcher/config"
	"github.com/caioeverest/ingresso-watcher/repository"
	"github.com/caioeverest/ingresso-watcher/service/errors"
)

type EventBody struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BaseDataFromEventResponse struct {
	Id          float64   `json:"id"`
	Name        string    `json:"name"`
	QtdAvaiable float64   `json:"qtd_avaiable"`
	Date        time.Time `json:"date"`
}

type EventFoundPayload struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Status string `json:"status"`
}

func SaveNewEvent(r repository.Interface, event EventBody) {
	log.Printf("Salvando evento %s", event.Name)
	r.Set(event.Id, event.Name)
}

func UpdateEvent(r repository.Interface, id, newName string) error {
	oldEventName, ok := r.GetById(id)
	if !ok {
		log.Print("Evento não encontrado")
		return errors.EventNotFound
	}
	log.Printf("Atualizando evento id %s de nome %s para %s", id, oldEventName, newName)
	r.Set(id, newName)
	return nil
}

func GetEventById(r repository.Interface, id string) (string, error) {
	event, ok := r.GetById(id)
	if !ok {
		log.Print("Evento não encontrado")
		return "", errors.EventNotFound
	}

	log.Printf("Evento %s encontrado", event)
	return event, nil
}

func GetAllRegistredEvents(r repository.Interface) []EventBody {
	rawEvents := r.GetAll()
	var events []EventBody
	for id, name := range rawEvents {
		event := EventBody{id, name}
		events = append(events, event)
	}
	return events
}

func StopWatchEvent(r repository.Interface, id string) error {
	if err := r.Delete(id); err != nil {
		log.Printf("Não foi possivel remover evento %s pois %s", id, err)
		return err
	}
	log.Printf("Evento %s deletado com sucesso!", id)
	return nil
}

func CheckIfHaveTickets(conf *config.Config, id string) (*EventFoundPayload, error) {
	payload, err := client.GetEventById(conf, id)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	event := payload.(map[string]interface{})
	status, ok := event["status"].(string)
	if ok && status == "SOLD_OUT" {
		return nil, errors.NoTickets
	}

	name, ok := event["name"].(string)
	url := fmt.Sprintf("https://www.ingressorapido.com.br/event/%s", id)

	return &EventFoundPayload{
		Name:   name,
		Url:    url,
		Status: status,
	}, nil
}
