package controllers

import (
	db "cards/apirestgin/database"
	"cards/apirestgin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO change slice to map [Person]uuid

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.People)
}

func PostAlbums(c *gin.Context) {
	var newAlbum domain.Person

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	db.People = append(db.People, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbums(c *gin.Context) {
	id := c.Param("id")

	posToDelete := 0

	for index, a := range db.People {
		if a.UUID == id {
			posToDelete = index
			return
		}
	}	

	db.People = deleteElement(db.People, posToDelete)
	c.IndentedJSON(http.StatusCreated, db.People )
}

func deleteElement(slice []domain.Person, index int) []domain.Person {
	return append(slice[:index], slice[index+1:]...)
 }

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range db.People {
		if a.UUID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albun no encontrado"})
}
