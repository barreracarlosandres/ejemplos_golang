package apirestgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Domain Person
var people = []domainPerson{
	{UUID: "61be616a-176e-4f41-8730-53a8366981d5", FistName: "Chelcey", Email: "mimi_solarjub@cultural.jnt", Ege: 16},
	{UUID: "918e414a-b031-4706-bf74-bfaea011aaba", FistName: "Alta", Email: "zedrick_pickleshe@titanium.gek", Ege: 30},
	{UUID: "e9844c1f-5669-4929-95c8-8410ac52e53d", FistName: "Shawanda", Email: "shareka_pitts8@floral.md", Ege: 40},
	{UUID: "2f1d8415-e141-4569-b34b-0aba2d968178", FistName: "Ryna", Email: "noura_vaillancourtwl@wind.jux", Ege: 20},
	{UUID: "bab4ef60-2044-4192-aaa7-49a50ae77317", FistName: "Paolo", Email: "aubrie_burchfieldd@marking.fv", Ege: 23},
}

//TODO change slice to map [Person]uuid

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func postAlbums(c *gin.Context) {
	var newAlbum domainPerson

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	people = append(people, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range people {
		if a.UUID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Albun no encontrado"})
}
