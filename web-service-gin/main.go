package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)


// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}


// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


// adds an album from JSON present in request body
func postAlbums(context *gin.Context) {
	var record album

	// bind received JSON to newAlbum
	if err := context.BindJSON(&record); err != nil {
		log.Println(err)
		return
	}

	// add the new album to the slice
	albums = append(albums, record)
	context.IndentedJSON(http.StatusCreated, record)
}


// responds with the list of all albums as JSON.
func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}


// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(context *gin.Context) {
    id := context.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, record := range albums {
        if record.ID == id {
            context.IndentedJSON(http.StatusOK, record)
            return
        }
    }

    context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func main() {
	log.SetPrefix("api: ")
	log.SetFlags(0) // suppresses the time, source file, and line number

	router := gin.Default()

	router.POST("/albums", postAlbums)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:7000")
}
