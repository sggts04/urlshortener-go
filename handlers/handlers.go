package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sggts04/urlshortener-go/data"
)

type URLShorteningHandler struct {
	r *data.Repository
}

func NewURLShorteningHandler(r *data.Repository) *URLShorteningHandler {
	return &URLShorteningHandler{r: r}
}

func (h *URLShorteningHandler) ServeFrontend(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *URLShorteningHandler) RegisterLongURL(c *gin.Context) {
	longURL := c.PostForm("longURL")
	if longURL == "" {
		// Long URL not specified
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "long url not specified"})
		return
	}

	customId := c.PostForm("customId")
	id, err := h.r.StoreLongURL(longURL, customId)

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

func (h *URLShorteningHandler) RedirectToLongURL(c *gin.Context) {
	id := c.Param("id")

	longURL, err := h.r.GetLongURL(id)
	if err != nil {
		if err.Error() == "shorturl not found" {
			// ID doesn't exist
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			// ID couldn't be generated.
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	} else {
		c.Redirect(http.StatusMovedPermanently, longURL)
	}
}
