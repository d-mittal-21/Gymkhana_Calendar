package handlers

import (
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/d-mittal-21/GymkhanaCalendar/types"
	"github.com/gin-gonic/gin"
)

func CreateEventHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var event types.Event
		if err := c.BindJSON(&event); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ref := client.Collection(types.EVENT_COLLECTION).NewDoc()

		_, err := ref.Set(c, map[string]interface{}{
			"title": event.Title,
			"desc":  event.Description,
			"venue": event.Venue,
			"start": event.Start,
			"end":   event.End,
			"link":  event.Link,
		})

		if err != nil {
			log.Printf("An error has occured: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusCreated, event)
	}

}
