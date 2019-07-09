package handler

import (
	"fmt"
	"net/http"

	"github.com/caioeverest/ingressoWatcher/service"
	"github.com/caioeverest/ingressoWatcher/service/errors"
	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) CreateEvent(c *gin.Context) {
	var body service.PostEventBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Body does not match",
		})
		return
	}

	service.SaveNewEvent(h.EventDB, body)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Event created",
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
	newName := c.GetString("new_name")
	id := c.Param("id")
	if newName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "empty field \"new_name\"",
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
		"message": fmt.Sprintf("Event name changed to %s", newName),
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
		"message": fmt.Sprintf("event with id %s deleted!", id),
	})
}

func (h *BaseHandler) TestEvent(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "empty field \"id\"",
		})
		return
	}

	items, err := service.TestEvent(h.Config, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"code":    http.StatusFailedDependency,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": items,
	})
}