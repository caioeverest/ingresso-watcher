package service

import (
	"testing"

	"github.com/caioeverest/ingresso-watcher/repository"
	"github.com/caioeverest/ingresso-watcher/service/errors"
)

func TestEventService(t *testing.T) {
	mockDB := repository.NewMemory()

	for _, tc := range []struct {
		Scenario string
		Create   EventBody
		Patch    EventBody
		Delete   bool
		Expect   EventBody
	}{
		{
			Scenario: "Shoud create and recover an event",
			Create:   EventBody{"4321", "Show of Teste"},
			Expect:   EventBody{"4321", "Show of Teste"},
		},
		{
			Scenario: "Shoud create and delete an event",
			Create:   EventBody{"4321", "Show of Teste"},
			Delete:   true,
			Expect:   EventBody{},
		},
		{
			Scenario: "Shoud create, patch and recover an event",
			Create:   EventBody{"4321", "Show of Teste"},
			Patch:    EventBody{"4321", "Test's show"},
			Expect:   EventBody{"4321", "Test's show"},
		},
		{
			Scenario: "Shoud create, patch, delete an event",
			Create:   EventBody{"4321", "Show of Teste"},
			Patch:    EventBody{"4321", "Test's show"},
			Delete:   true,
			Expect:   EventBody{},
		},
	} {
		t.Run(tc.Scenario, func(t *testing.T) {
			SaveNewEvent(mockDB, tc.Create)
			if tc.Patch != (EventBody{}) {
				if err := UpdateEvent(mockDB, tc.Patch.Id, tc.Patch.Name); err != nil {
					t.Errorf(err.Error())
				}
			}
			if tc.Delete {
				err := StopWatchEvent(mockDB, tc.Create.Id)
				if err != nil {
					t.Errorf(err.Error())
				}
			}
			name, err := GetEventById(mockDB, tc.Create.Id)
			if err != nil && (tc.Expect == (EventBody{}) && err != errors.EventNotFound) {
				t.Errorf(err.Error())
			}
			if name != tc.Expect.Name {
				t.Errorf("Got: %s\nExpect: %s", name, tc.Expect.Name)
			}
		})
	}
}
