// A standalone program (as opposed to a library) is always in package main.
package main

//import the packages you’ll need to support the code you’ve just written.
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album. Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   float32 `json:"year"`
}

func main() {
	router := gin.Default()
	// Initialize a Gin router using Default.
	router.GET("/albums", getAlbums)
	// Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
	// Use the Run function to attach the router to an http.Server and start the server.
}

// albums slice to seed record album data. This slice of album structs containing data you’ll use to start.
var albums = []album{
	{ID: "1", Title: "Thriller", Artist: "Michael Jackson", Year: 1982},
	{ID: "2", Title: "Songs in the Key of Life", Artist: "Stevie Wonder", Year: 1976},
	{ID: "3", Title: "A Love Supreme", Artist: "John Coltrane", Year: 1965},
	{ID: "4", Title: "What's Going On?", Artist: "Marvin Gaye", Year: 1971},
}

// This getAlbums function creates JSON from the slice of album structs, writing the JSON into the response..
func getAlbums(c *gin.Context) {
	// getAlbums function takes a gin.Context parameter. Note that you could have given this function any name – neither Gin nor Go require a particular function name format. gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)
	c.IndentedJSON(http.StatusOK, albums)
	// The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK. Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	// Use Context.BindJSON to bind the request body to newAlbum && Append the album struct initialized from the JSON to the albums slice.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice. Add a 201 status code to the response, along with JSON representing the album you added.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
