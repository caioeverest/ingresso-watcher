package main

import (
	"github.com/caioeverest/ingressoWatcher/service"
	"log"

	"github.com/caioeverest/ingressoWatcher/client"
	"github.com/caioeverest/ingressoWatcher/config"
	"github.com/caioeverest/ingressoWatcher/handler"
	"github.com/caioeverest/ingressoWatcher/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	conf := config.InitConfig()
	contactDb := repository.NewMemory()
	eventDb := repository.NewMemory()
	whatsapp, err := client.InitWhatsAppConnection()
	if err != nil {
		log.Panicf("Not able to connect with whatsapp %s", err)
	}

	go service.Monitor(contactDb, eventDb, whatsapp, conf)

	h := handler.SetHandlers(conf, whatsapp, contactDb, eventDb)
	app := gin.Default()

	api := app.Group("/api")
	{
		contact := api.Group("/contact")
		{
			contact.POST("/", h.CreateContact)
			contact.GET("/", h.GetAllContacts)
			contact.GET("/:phone", h.GetContactByPhone)
			contact.PATCH("/:phone", h.UpdateContact)
			contact.DELETE("/:phone", h.DeleteContact)
		}

		event := api.Group("/event")
		{
			event.POST("/", h.CreateEvent)
			event.GET("/", h.GetAllEvents)
			event.GET("/:id", h.GetEventById)
			event.PATCH("/:id", h.UpdateEvent)
			event.DELETE("/:id", h.DeleteEvent)
		}

		api.GET("/test-event", h.TestEvent)
		api.POST("/notify", h.NotifyTest)
	}

	if err := app.Run(); err != nil {
		log.Panicf("something goes wrong: %s", err)
	}
}
