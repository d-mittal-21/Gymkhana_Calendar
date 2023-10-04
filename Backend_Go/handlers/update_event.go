package handlers

import (
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/d-mittal-21/GymkhanaCalendar/Backend_Go/types"
	"github.com/gin-gonic/gin"
)

func UpdateEventHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {

		title := c.Param("title")
		var event types.Event
		if err := c.BindJSON(&event); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		event.Title = title
		_, err := client.Collection(types.EVENT_COLLECTION).Doc(event.Title).Set(c, map[string]interface{}{
			"title":   event.Title,
			"desc":    event.Description,
			"venue":   event.Venue,
			"start":   event.Start,
			"end":     event.End,
			"reglink": event.RegLink,
		}, firestore.MergeAll)

		if err != nil {
			log.Printf("An error has occured while updating: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusOK, event)
	}
}
