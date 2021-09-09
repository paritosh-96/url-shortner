package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"urlShortner/common"
)

func GetActualURL(c *gin.Context) {
	fmt.Println("Inside func")
	shortUrl := c.Request.URL.Query().Get("url")
	key := c.Request.URL.Query().Get("key")

	if shortUrl == "" || key == "" {
		http.Error(c.Writer, "URL/ key can not be blank", 400)
		return
	}

	actualURL, err := common.GetActualURL(key, shortUrl)
	if err != nil {
		log.Println(err.Error())
		http.Error(c.Writer, "Error creating actual URL", 500)
	} else {
		response := common.GetResponse(actualURL)
		_ = json.NewEncoder(c.Writer).Encode(response)
	}
}

func ShortenURL(c * gin.Context) {
	actualUrl := c.Request.URL.Query().Get("url")
	key := c.Request.URL.Query().Get("key")

	if actualUrl == "" || key == "" {
		http.Error(c.Writer, "URL/ key can not be blank", 400)
		return
	}

	shortURL, err := common.GetShortURL(key, actualUrl)
	if err != nil {
		log.Println(err.Error())
		http.Error(c.Writer, "Error creating short URL", 500)
	} else {
		response := common.GetResponse(shortURL)
		_ = json.NewEncoder(c.Writer).Encode(response)
	}
}