package handler

import (
	"fmt"
	"net/http"

	"github.com/caioeverest/ingressoWatcher/service"
	"github.com/caioeverest/ingressoWatcher/service/errors"
	"github.com/gin-gonic/gin"
)

func (h *BaseHandler) CreateContact(c *gin.Context) {
	var body service.PostContactBody
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Body does not match",
		})
		return
	}

	service.AddNewContact(h.ContactDB, body)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Contact created with success",
	})
}

func (h *BaseHandler) GetAllContacts(c *gin.Context) {
	contactList := service.GetContactList(h.ContactDB)
	c.JSON(http.StatusOK, contactList)
}

func (h *BaseHandler) GetContactByPhone(c *gin.Context) {
	phone := c.Param("phone")
	contact, err := service.GetContactByNumber(h.ContactDB, phone)
	if err != nil {
		switch err {
		case errors.ContactNotFound:
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
		"message": contact,
	})
}

func (h *BaseHandler) UpdateContact(c *gin.Context) {
	newName := c.GetString("new_name")
	phone := c.Param("phone")
	if newName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "empty field \"new_name\"",
		})
		return
	}

	if err := service.ChangeContactName(h.ContactDB, phone, newName); err != nil {
		switch err {
		case errors.ContactNotFound:
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
		"message": fmt.Sprintf("Contact changed to name: %s", newName),
	})
}

func (h *BaseHandler) DeleteContact(c *gin.Context) {
	phone := c.Param("phone")
	err := service.DeleteContact(h.ContactDB, phone)
	if err != nil {
		switch err {
		case errors.ContactNotFound:
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
		"message": fmt.Sprintf("deleted contact number %s with success", phone),
	})
}
