package service

import (
	"log"

	"github.com/caioeverest/ingresso-watcher/client"

	"github.com/caioeverest/ingresso-watcher/repository"
	"github.com/caioeverest/ingresso-watcher/service/errors"
)

type ContactBody struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func AddNewContact(r repository.Interface, contact ContactBody, wpp client.WhatsAppInterface) error {
	log.Printf("Salvando contato de %s", contact.Name)
	r.Set(contact.Phone, contact.Name)
	return SendGreetingsMessage(wpp, contact.Phone, contact.Name)
}

func ChangeContactName(r repository.Interface, phone, newName string) error {
	oldName, ok := r.GetById(phone)
	if !ok {
		log.Print("Contato não encontrado")
		return errors.ContactNotFound
	}
	log.Printf("Atualizando nome do contato de dumero telefone %s de %s para %s", phone, oldName, newName)
	r.Set(phone, newName)
	return nil
}

func GetContactByNumber(r repository.Interface, phone string) (string, error) {
	name, ok := r.GetById(phone)
	if !ok {
		log.Print("Contato não encontrado")
		return "", errors.ContactNotFound
	}
	log.Printf("Contato de %s encontrado", name)
	return name, nil
}

func GetContactList(r repository.Interface) []ContactBody {
	rawContacts := r.GetAll()
	var contacts []ContactBody
	for phone, name := range rawContacts {
		contact := ContactBody{phone, name}
		contacts = append(contacts, contact)
	}
	return contacts
}

func DeleteContact(r repository.Interface, phone string) error {
	if err := r.Delete(phone); err != nil {
		log.Printf("Não foi possivel remover o telefone %s pois %s", phone, err)
		return err
	}
	log.Printf("Numero %s removido da lista de notificação com sucesso!", phone)
	return nil
}
