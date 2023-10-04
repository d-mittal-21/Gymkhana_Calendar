package handlers

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/d-mittal-21/GymkhanaCalendar/Backend_Go/types"
	"github.com/gin-gonic/gin"
)

func DeleteEventHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {

		title := c.Param("title")

		_, err := client.Collection(types.EVENT_COLLECTION).Doc(title).Delete(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
