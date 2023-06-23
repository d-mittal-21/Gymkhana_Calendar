package handlers

import (
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/d-mittal-21/GymkhanaCalendar/types"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func ListEventHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var result []types.Event

		iter := client.Collection(types.EVENT_COLLECTION).Documents(c)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, "")
				return
			}
			result = append(result,
				types.Event{
					ID:          doc.Ref.ID,
					Title:       doc.Data()["title"].(string),
					Description: doc.Data()["desc"].(string),
					Start:       doc.Data()["start"].(time.Time),
					End:         doc.Data()["end"].(time.Time),
					Link:        doc.Data()["link"].(string),
				})
		}

		c.JSON(http.StatusOK, result)

	}
}
