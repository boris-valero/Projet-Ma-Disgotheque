//A standalone program (as opposed to a library) is always in package main.
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
    Year  float32 `json:"year"`
}

func main() {
	router := gin.Default()
// Initialize a Gin router using Default.
	router.GET("/albums", getAlbums)
// Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	router.Run("localhost:8080")
//Use the Run function to attach the router to an http.Server and start the server.
}

// albums slice to seed record album data. This slice of album structs containing data you’ll use to start.
var albums = []album{
    {ID: "1", Title: "Mauvais oeil", Artist: "Lunatic", Year: 2000},
    {ID: "2", Title: "L'école du micro d'argent", Artist: "IAM", Year: 1997},
    {ID: "3", Title: "Le combat continue", Artist: "Ideal J", Year: 1998},
    {ID: "4", Title: "Temps mort", Artist: "Booba", Year: 2002},
    {ID: "5", Title: "Du rire aux larmes", Artist: "Sniper", Year: 2001},
    {ID: "6", Title: "Jeunes coupables et libres", Artist: "X-Men", Year: 2001},
    {ID: "7", Title: "Où je vis", Artist: "Shurik'n", Year: 1998},
    {ID: "8", Title: "Suprême NTM", Artist: "Suprême NTM", Year: 1998},
    {ID: "9", Title: "Conçu pour durer", Artist: "La Cliqua", Year: 1995},
    {ID: "10", Title: "KLR", Artist: "Saïan Supa Crew", Year: 1999},
    {ID: "11", Title: "Opéra Puccino", Artist: "Oxmo Puccino", Year: 1998},
    {ID: "12", Title: "Détournement de son", Artist: "Fabe", Year: 1998},
    {ID: "13", Title: "Si Dieu veut", Artist: "Fonky Family", Year: 1997},
    {ID: "14", Title: "Scred Selexion 99/2000", Artist: "Scred Connexion ", Year: 2000},
}

// This getAlbums function creates JSON from the slice of album structs, writing the JSON into the response..
func getAlbums(c *gin.Context) {
// getAlbums function takes a gin.Context parameter. Note that you could have given this function any name – neither Gin nor Go require a particular function name format. gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)
	c.IndentedJSON(http.StatusOK, albums)
//The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK. Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.
}
