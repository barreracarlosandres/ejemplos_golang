package apirestgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleGetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, *getAll())
}

func handleAddPerson(c *gin.Context) {
	var newAlbum DomainPerson

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, Message{Msg: "Don't add person"})
		return
	}

	add(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func handleDeletePerson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, delete(&StoragePeople, c.Param("uuid")))
}

func handleGetPersonByID(c *gin.Context) {
	uuid := c.Param("uuid")

	var person = getPersonByUUID(&StoragePeople, uuid)
	if person.UUID != "" {
		c.IndentedJSON(http.StatusOK, person)
		return
	}

	c.IndentedJSON(http.StatusNotFound, Message{Msg: "not found UUID " + uuid})
}
