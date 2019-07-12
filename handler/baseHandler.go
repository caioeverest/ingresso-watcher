package handler

import (
	"github.com/caioeverest/ingresso-watcher/client"
	"github.com/caioeverest/ingresso-watcher/config"
	"github.com/caioeverest/ingresso-watcher/repository"
)

type BaseHandler struct {
	Config 		*config.Config
	WppClient 	*client.WppConnection
	ContactDB 	repository.Interface
	EventDB 	repository.Interface
}

func SetHandlers(c *config.Config, wpp *client.WppConnection, cdb, edb repository.Interface) *BaseHandler {
	return &BaseHandler{
		c,
		wpp,
		cdb,
		edb,
	}
}