package main

import (
	"log"

	"github.com/caioeverest/ingresso-watcher/service"

	"github.com/caioeverest/ingresso-watcher/client"
	"github.com/caioeverest/ingresso-watcher/config"
	"github.com/caioeverest/ingresso-watcher/handler"
	"github.com/caioeverest/ingresso-watcher/repository"
	"github.com/gin-gonic/contrib/static"
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

	app.Use(static.Serve("/", static.LocalFile("./ui/build", true)))

	app.Use(static.Serve("/contatos", static.LocalFile("./ui/build", true)))

	app.NoRoute(static.Serve("/", static.LocalFile("./ui/build", true)))

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

		api.POST("/notify", h.NotifyTest)
	}

	if err := app.Run(":" + conf.HttpPort); err != nil {
		log.Panicf("something goes wrong: %s", err)
	}
}
