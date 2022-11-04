package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct { // 주어진 형태(틀)
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{ // 초기값 저장
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1") // v1으로 묶어서 라우팅 가능
	{
		v1.GET("/albums", getAlbums) // 각각 해당하는 함수 지정
		v1.GET("/albums/:id", getAlbumByID)
		v1.POST("/albums", postAlbums)
	}
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return // 에러 발생시 출력
	}

	albums = append(albums, newAlbum) // 새로운 앨범을 추가 (append)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
