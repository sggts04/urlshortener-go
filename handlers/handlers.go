package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sggts04/urlshortener-go/data"
)

func ServeFrontend(c *gin.Context) {
	// TODO
}

func RegisterLongURL(c *gin.Context) {
	longURL := c.PostForm("longURL")
	if longURL == "" {
		// Long URL not specified
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "long url not specified"})
		return
	}

	customId := c.PostForm("customId")
	id, err := data.StoreLongURL(longURL, customId)

	if err != nil {
		if err.Error() == "custom id already exists" {
			// Custom ID collision
			c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
		} else {
			// ID couldn't be generated.
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func RedirectToLongURL(c *gin.Context) {
	id := c.Param("id")

	longURL, err := data.GetLongURL(id)
	if err != nil {
		// ID doesn't exist
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
