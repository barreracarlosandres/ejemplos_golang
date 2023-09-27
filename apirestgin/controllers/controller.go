package controllers

import (
	db "cards/apirestgin/database"
	"cards/apirestgin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO change slice to map [Person]uuid

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.People)
}

func postAlbums(c *gin.Context) {
	var newAlbum domain.Person

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	db.People = append(db.People, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range db.People {
		if a.UUID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albun no encontrado"})
}
