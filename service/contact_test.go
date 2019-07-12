package service

import (
	"github.com/caioeverest/ingresso-watcher/client"
	"testing"

	"github.com/caioeverest/ingresso-watcher/repository"
	"github.com/caioeverest/ingresso-watcher/service/errors"
)

func TestContactService(t *testing.T) {
	mockDB := repository.NewMemory()
	mockWpp := client.InitMock()

	for _, tc := range []struct {
		Scenario string
		Create   ContactBody
		Patch    ContactBody
		Delete   bool
		Expect   ContactBody
	}{
		{
			Scenario: "Shoud create and recover a simple contact",
			Create:   ContactBody{"123456789", "Testing"},
			Expect:   ContactBody{"123456789", "Testing"},
		},
		{
			Scenario: "Shoud create and delete a simple contact",
			Create:   ContactBody{"123456789", "Testing"},
			Delete:   true,
			Expect:   ContactBody{},
		},
		{
			Scenario: "Shoud create, patch and recover a simple contact",
			Create:   ContactBody{"123456789", "Testing"},
			Patch:    ContactBody{"123456789", "Fulano"},
			Expect:   ContactBody{"123456789", "Fulano"},
		},
		{
			Scenario: "Shoud create, patch, delete a simple contact",
			Create:   ContactBody{"123456789", "Testing"},
			Patch:    ContactBody{"123456789", "Fulano"},
			Delete:   true,
			Expect:   ContactBody{},
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			AddNewContact(mockDB, tc.Create, mockWpp)
			if tc.Patch != (ContactBody{}) {
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
			if err != nil && (tc.Expect == (ContactBody{}) && err != errors.ContactNotFound) {
				t.Errorf(err.Error())
			}
			if name != tc.Expect.Name {
				t.Errorf("Got: %s\nExpect: %s", name, tc.Expect.Name)
			}
		})
	}
}
