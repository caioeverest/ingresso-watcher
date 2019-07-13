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

func TestEvent(conf *config.Config, id string) ([]BaseDataFromEventResponse, error) {
	events, err := client.GetEventById(conf, id)
	if err != nil {
		return nil, err
	}

	var res []BaseDataFromEventResponse
	for _, i := range events {
		item := i.(map[string]interface{})
		localDate, err := time.Parse(time.RFC3339, item["presentation_local_date_time"].(string))
		if err != nil {
			return nil, err
		}
		irItem := BaseDataFromEventResponse{
			Id:          item["id"].(float64),
			Name:        item["name"].(string),
			QtdAvaiable: item["total_available"].(float64),
			Date:        localDate,
		}
		if irItem.Date.After(time.Now()) {
			res = append(res, irItem)
		}
	}

	if len(res) <= 0 {
		return nil, errors.LateEvent
	}

	return res, nil
}

func CheckIfHaveTickets(conf *config.Config, id string) (map[string]string, error) {
	events, err := client.GetEventById(conf, id)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	res := make(map[string]string)

	for _, i := range events {
		item := i.(map[string]interface{})
		qtdAvaiable := item["total_available"].(float64)

		if qtdAvaiable > 0 {
			name := item["name"].(string)
			url := fmt.Sprintf("https://www.ingressorapido.com.br/event/%s-1", id)
			res[name] = url
		}
	}

	if len(res) <= 0 {
		return nil, errors.NoTickets
	}

	return res, nil
}
