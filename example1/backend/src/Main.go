package main

import ( 
	"net/http" 
	"github.com/gin-gonic/gin" 
	// "github.com/sashabaranov/fastapi"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var albumsData = []album{
	{
		ID:     "1",
		Title:  "The Secret",
		Artist: "John",
		Price:  56.89,
	},
	{
		ID:     "2",
		Title:  "The Lock",
		Artist: "Brown",
		Price:  23.89,
	}, {
		ID:     "3",
		Title:  "The Fire Wings",
		Artist: "Kalam",
		Price:  16.00,
	}, {
		ID:     "4",
		Title:  "The Kinder Garden",
		Artist: "Shoban",
		Price:  89.89,
	},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOk, albumsData)
}

func main() {
	router := gin.Default()

	// fastAPIRouter := fastapi.NewRouter()
	// fastAPIRouter.AddCall("/albums", getAlbums)
	// router.GET("/albums", fastAPIRouter.GinHandler)
	
	router.GET("/albums", getAlbums)
	router.Run("localhost:6060")
}
