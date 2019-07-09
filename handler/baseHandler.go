package handler

import (
	"github.com/caioeverest/ingressoWatcher/client"
	"github.com/caioeverest/ingressoWatcher/config"
	"github.com/caioeverest/ingressoWatcher/repository"
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