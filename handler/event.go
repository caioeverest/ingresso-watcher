package handler

import (
	"fmt"
	"net/http"

	"github.com/caioeverest/ingresso-watcher/service"
	"github.com/caioeverest/ingresso-watcher/service/errors"
	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) CreateEvent(c *gin.Context) {
	var body service.EventBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Corpo da requisição inválido",
		})
		return
	}

	service.SaveNewEvent(h.EventDB, body)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Evento criado!",
	})
}

func (h *BaseHandler) GetAllEvents(c *gin.Context) {
	events := service.GetAllRegistredEvents(h.EventDB)
	c.JSON(http.StatusOK, events)
}

func (h *BaseHandler) GetEventById(c *gin.Context) {
	id := c.Param("id")
	event, err := service.GetEventById(h.EventDB, id)
	if err != nil {
		switch err {
		case errors.EventNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": event,
	})
}

func (h *BaseHandler) UpdateEvent(c *gin.Context) {
	newName := c.Query("new_name")
	id := c.Param("id")
	if newName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Campo \"new_name\" veio vazio",
		})
		return
	}

	if err := service.UpdateEvent(h.EventDB, id, newName); err != nil {
		switch err {
		case errors.EventNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("Nome do evento alterado para: %s", newName),
	})
}

func (h *BaseHandler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteContact(h.EventDB, id)
	if err != nil {
		switch err {
		case errors.EventNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": fmt.Sprintf("Evento id %s deletado com sucesso!", id),
	})
}
