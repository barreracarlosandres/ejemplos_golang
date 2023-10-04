package apirestgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO change slice to map [Person]uuid

func handleGetPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, StoragePeople)
}

func handleAddPerson(c *gin.Context) {
	var newAlbum DomainPerson

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	StoragePeople = append(StoragePeople, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func handleDeletePerson(c *gin.Context) {
	id := c.Param("id")

	posToDelete := 0

	for index, a := range StoragePeople {
		if a.UUID == id {
			posToDelete = index
			return
		}
	}

	StoragePeople = deleteElement(StoragePeople, posToDelete)
	c.IndentedJSON(http.StatusCreated, StoragePeople)
}

func deleteElement(slice []DomainPerson, index int) []DomainPerson {
	return append(slice[:index], slice[index+1:]...)
}

func handleGetPersonByID(c *gin.Context) {
	uuid := c.Param("uuid")

	for _, a := range StoragePeople {
		if a.UUID == uuid {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albun no encontrado"})
}
