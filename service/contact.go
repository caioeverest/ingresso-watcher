package service

import (
	"log"

	"github.com/caioeverest/ingressoWatcher/repository"
	"github.com/caioeverest/ingressoWatcher/service/errors"
)

type PostContactBody struct {
	Phone string `json:"phone"`
	Name  string `json:"name"`
}

func AddNewContact(r repository.Interface, contact PostContactBody) {
	log.Printf("Salvando contato de %s", contact.Name)
	r.Set(contact.Phone, contact.Name)
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

func GetContactList(r repository.Interface) map[string]string {
	contacts := r.GetAll()
	log.Printf("Foram encontrados %d contatos na lista", len(contacts))
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
