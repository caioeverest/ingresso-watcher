package service

import (
	"testing"

	"github.com/caioeverest/ingressoWatcher/repository"
	"github.com/caioeverest/ingressoWatcher/service/errors"
)

func TestContactService(t *testing.T) {
	mockDB := repository.NewMemory()

	for _, tc := range []struct {
		Scenario string
		Create	PostContactBody
		Patch	PostContactBody
		Delete	bool
		Expect	PostContactBody
	} {
		{
			Scenario: 	"Shoud create and recover a simple contact",
			Create: 	PostContactBody{"123456789", "Testing"},
			Expect: 	PostContactBody{"123456789", "Testing"},
		},
		{
			Scenario: 	"Shoud create and delete a simple contact",
			Create: 	PostContactBody{"123456789", "Testing"},
			Delete: 	true,
			Expect: 	PostContactBody{},
		},
		{
			Scenario: 	"Shoud create, patch and recover a simple contact",
			Create: 	PostContactBody{"123456789", "Testing"},
			Patch: 		PostContactBody{"123456789", "Fulano"},
			Expect: 	PostContactBody{"123456789", "Fulano"},
		},
		{
			Scenario: 	"Shoud create, patch, delete a simple contact",
			Create: 	PostContactBody{"123456789", "Testing"},
			Patch: 		PostContactBody{"123456789", "Fulano"},
			Delete: 	true,
			Expect: 	PostContactBody{},
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			AddNewContact(mockDB, tc.Create)
			if tc.Patch != (PostContactBody{}) {
				if err := ChangeContactName(mockDB, tc.Patch.Phone, tc.Patch.Name); err != nil {
					t.Errorf(err.Error())
				}
			}
			if tc.Delete {
				err := DeleteContact(mockDB, tc.Create.Phone)
				if err != nil {
					t.Errorf(err.Error())
				}
			}
			name, err := GetContactByNumber(mockDB, tc.Create.Phone)
			if err != nil && (tc.Expect == (PostContactBody{}) && err != errors.ContactNotFound) {
				t.Errorf(err.Error())
			}
			if name != tc.Expect.Name {
				t.Errorf("Got: %s\nExpect: %s", name, tc.Expect.Name)
			}
		})
	}
}
