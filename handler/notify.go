package handler

import (
	"log"
	"net/http"

	"github.com/caioeverest/ingressoWatcher/service"
	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) NotifyTest(c *gin.Context) {
	phone, ok := c.GetQuery("phone")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "phone nedded",
		})
		return
	}

	if err := service.SendTestMessage(h.WppClient, phone); err != nil {
		log.Print(err)
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"code":    http.StatusFailedDependency,
			"message": err.Error(),
		})
		return
	}

	c.AbortWithStatus(http.StatusAccepted)
}
