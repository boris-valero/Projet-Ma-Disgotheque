package main

// Importation de packages externes nécessaires pour le fonctionnement de l'application.
import (
	"net/http" // Package permettant de manipuler des requêtes et réponses HTTP.
	"github.com/gin-gonic/gin"
)

// Définition d'une structure `album` car en Go, une structure est un type de données composite qui permet de regrouper des variables sous un même nom
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   float32 `json:"year"`
}

func main() {
// Création d'un nouveau routeur Gin avec les paramètres par défaut.
	router := gin.Default()
// Associe la route GET "/albums" à la fonction `getAlbums` pour récupérer la liste des albums.
	router.GET("/albums", getAlbums)
// Associe la route POST "/albums" à la fonction `postAlbums` pour ajouter un nouvel album.
	router.POST("/albums", postAlbums)
// Associe la route GET "/albums/:id" à la fonction `getAlbumByID` pour récupérer un album spécifique par son ID.
	router.GET("/albums/:id", getAlbumByID)
// Démarre le serveur HTTP sur le port 8080 et écoute les requêtes entrantes.
	router.Run("localhost:8080")
}

// Déclaration d'une variable globale `albums` qui est une slice de `album`. Cette slice contient une liste prédéfinie d'albums musicaux avec leurs détails.
var albums = []album{
// Chaque élément de la slice est une instance de la structure `album`, initialisée avec des valeurs spécifiques.
	{ID: "1", Title: "Thriller", Artist: "Michael Jackson", Year: 1982},
	{ID: "2", Title: "Songs in the Key of Life", Artist: "Stevie Wonder", Year: 1976},
	{ID: "3", Title: "A Love Supreme", Artist: "John Coltrane", Year: 1965},
	{ID: "4", Title: "What's Going On?", Artist: "Marvin Gaye", Year: 1971},
}

// La fonction getAlbums crée un JSON depuis la slice de la variable globale `albums`, écrivant le JSON dans la réponse.
func getAlbums(c *gin.Context) {
// Utilise la méthode IndentedJSON de l'objet context `c` pour envoyer une réponse. Le premier paramètre, `http.StatusOK`, indique que la requête a réussi et que le serveur renvoie un code de statut HTTP 200, qui signifie "OK". Le deuxième paramètre est la donnée à envoyer en réponse, ici la liste des albums. La méthode IndentedJSON assure que la réponse JSON est bien formatée avec des indentations, rendant le contenu plus lisible pour les humains.
	c.IndentedJSON(http.StatusOK, albums)
}

// Définition de la fonction `postAlbums` pour gérer les requêtes POST sur "/albums"
func postAlbums(c *gin.Context) {
// Déclare une nouvelle variable `newAlbum` de type `album`
	var newAlbum album
// Appelle BindJSON pour tenter de lier le JSON reçu dans la requête à `newAlbum`. Si une erreur survient, termine la fonction
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
// Ajoute `newAlbum` à la liste des albums
	albums = append(albums, newAlbum)
// Envoie le nouvel album en JSON avec une indentation, avec le code de statut HTTP 201 Créé
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Définition de la fonction `getAlbumByID` pour gérer les requêtes GET sur "/albums/:id"
func getAlbumByID(c *gin.Context) {
// Récupère l'ID de l'album depuis l'URL
	id := c.Param("id")
// Parcourt la liste des albums au sein d'une boucle
	for _, a := range albums {
// Si l'ID de l'album courant correspond à l'ID recherché
		if a.ID == id {
// Envoie l'album en JSON avec une indentation, avec le code de statut HTTP 200
			c.IndentedJSON(http.StatusOK, a)
// Termine la fonction
			return
		}
	}
// Si aucun album correspondant n'a été trouvé, envoie un message d'erreur en JSON avec le code de statut HTTP 404 Non trouvé
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
